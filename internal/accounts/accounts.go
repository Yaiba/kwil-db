package accounts

import (
	"context"
	"encoding/binary"
	"errors"
	"fmt"
	"math/big"
	"sync"

	"github.com/kwilteam/kwil-db/core/log"
)

// CommitRegister is an interface for registering a commit.
type CommitRegister interface {
	// Skip returns true if the commit should be skipped.
	// This signals that the account store should not be updated,
	// and simply return nil.
	Skip() bool

	// Register registers a commit.
	// This should be called when data is written to the database.
	Register(value []byte) error
}

type Datastore interface {
	Execute(ctx context.Context, stmt string, args map[string]any) ([]map[string]any, error)
	Query(ctx context.Context, query string, args map[string]any) ([]map[string]any, error)
}

type AccountStore struct {
	db            Datastore
	log           log.Logger
	rw            sync.RWMutex
	gasEnabled    bool
	noncesEnabled bool

	committable CommitRegister
}

func NewAccountStore(ctx context.Context, datastore Datastore, committable CommitRegister, opts ...AccountStoreOpts) (*AccountStore, error) {
	ar := &AccountStore{
		log:         log.NewNoOp(),
		db:          datastore,
		committable: committable,
	}

	for _, opt := range opts {
		opt(ar)
	}

	err := ar.initTables(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize tables: %w", err)
	}

	return ar, nil
}

func (a *AccountStore) Close() error {
	return nil
}

func (a *AccountStore) GetAccount(ctx context.Context, ident []byte) (*Account, error) {
	a.rw.RLock()
	defer a.rw.RUnlock()

	return a.getAccountReadOnly(ctx, ident)
}

// Transfer sends an amount from the sender's balance to another account. The
// amount sent is given by the amount. This does not affect the sending
// account's nonce; a Spend should precede this to pay for required transaction
// gas and validate/advance the nonce.
func (a *AccountStore) Transfer(ctx context.Context, to, from []byte, amt *big.Int) error {
	// Ensure that the from account balance is sufficient.
	account, err := a.getAccountSynchronous(ctx, from)
	if err != nil {
		return err
	}
	newBal, err := account.validateSpend(amt)
	if err != nil {
		return err
	}
	// Update or create the to account with the transferred amount.
	toAcct, err := a.getOrCreateAccount(ctx, to)
	if err != nil {
		return err
	}
	// Decrement the from account balance first.
	err = a.updateAccount(ctx, from, newBal, account.Nonce)
	if err != nil {
		return err
	}
	toBal := big.NewInt(0).Add(toAcct.Balance, amt)
	err = a.updateAccount(ctx, to, toBal, toAcct.Nonce)
	if err != nil {
		return err
	}
	return a.committable.Register(transferBytes(to, from, amt))
}

func transferBytes(to, from []byte, amt *big.Int) []byte {
	var b []byte
	b = append(b, to...)
	b = append(b, from...)
	return append(b, amt.Bytes()...)
}

// Spend specifies a the fee and nonce of a transaction for an account. The
// amount has historically been associated with the transaction's fee (to pay
// for gas) i.e. the price of a certain transaction type.
type Spend struct {
	AccountID []byte
	Amount    *big.Int
	Nonce     int64
}

func (s *Spend) bytes() []byte {
	bts := s.AccountID
	bts = append(bts, s.Amount.Bytes()...)

	binary.LittleEndian.AppendUint64(bts, uint64(s.Nonce))

	return bts
}

// Send might be used to have the value transfer be atomic with the sender's
// transaction related updates (pay fee and update nonce). But I think these
// operations are distinct since the transaction is in a block if we're doing
// this operation, so they pay gas and update their nonce.
/*zzz
type Send struct {
	From  Spend
	To    []byte
	Value *big.Int
}
*/

// Spend spends an amount from an account. It blocks until the spend is written to the database.
func (a *AccountStore) Spend(ctx context.Context, spend *Spend) error {
	a.rw.Lock()
	defer a.rw.Unlock()

	if a.committable.Skip() {
		return nil
	}

	balance, nonce, err := a.checkSpend(ctx, spend)
	if err != nil {
		return fmt.Errorf("failed to check spend: %w", err)
	}

	err = a.updateAccount(ctx, spend.AccountID, balance, nonce)
	if err != nil {
		return fmt.Errorf("failed to update account: %w", err)
	}

	return a.committable.Register(spend.bytes())
}

// Credit credits an account. If the account does not exist, it will be created.
func (a *AccountStore) Credit(ctx context.Context, acctID []byte, amt *big.Int) error {
	a.rw.Lock()
	defer a.rw.Unlock()

	if a.committable.Skip() {
		return nil
	}

	// If exists, add to balance; if not, insert this balance and zero nonce.
	account, err := a.getAccountSynchronous(ctx, acctID)
	if err != nil {
		if !errors.Is(err, ErrAccountNotFound) {
			return err
		}
		return a.createAccountWithBalance(ctx, acctID, amt)
	}

	bal := new(big.Int).Add(account.Balance, amt)
	err = a.updateAccount(ctx, account.Identifier, bal, account.Nonce) // same nonce
	if err != nil {
		return fmt.Errorf("failed to update account: %w", err)
	}

	b := append(acctID, amt.Bytes()...)
	return a.committable.Register(b)
}

// checkSpend checks that a spend is valid.  If gas costs are enabled, it checks that the account has enough gas to pay for the spend.
// If nonces are enabled, it checks that the nonce is correct.  It returns the new balance and nonce if the spend is valid. It returns an
// error if the spend is invalid.
func (a *AccountStore) checkSpend(ctx context.Context, spend *Spend) (*big.Int, int64, error) {
	var account *Account
	var err error
	if a.gasEnabled && spend.Amount.Cmp(big.NewInt(0)) > 0 { // don't automatically create accounts when gas is required
		account, err = a.getAccountSynchronous(ctx, spend.AccountID)
	} else { // with no gas or a free transaction, we'll create the account if it doesn't exist
		account, err = a.getOrCreateAccount(ctx, spend.AccountID)
	}
	if err != nil {
		return nil, 0, fmt.Errorf("failed to get account: %w", err)
	}

	nonce := account.Nonce
	if a.noncesEnabled {
		err = account.validateNonce(spend.Nonce)
		if err != nil {
			return nil, 0, fmt.Errorf("failed to validate nonce: %w", err)
		}

		nonce = spend.Nonce
	}

	balance := account.Balance
	if a.gasEnabled {
		balance, err = account.validateSpend(spend.Amount)
		if err != nil {
			return nil, 0, fmt.Errorf("failed to subtract gas: %w", err)
		}
	}

	return balance, nonce, nil
}