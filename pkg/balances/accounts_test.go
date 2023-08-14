package balances_test

import (
	"context"
	"errors"
	"math/big"
	"testing"

	"github.com/kwilteam/kwil-db/pkg/balances"
	accountTesting "github.com/kwilteam/kwil-db/pkg/balances/testing"
	"github.com/stretchr/testify/assert"
)

const (
	account1 = "account1"
	account2 = "account2"
)

func Test_Accounts(t *testing.T) {
	type testCase struct {
		name          string
		spends        []*balances.Spend
		gasOn         bool
		noncesOn      bool
		finalBalances map[string]*balances.Account
		// the error must be triggered once
		err error
	}

	// once we have a way to increase balances in accounts, we will have to add tests
	// for spending a valid amount
	testCases := []testCase{
		{
			name: "gas off, nonces on",
			spends: []*balances.Spend{
				newSpend(account1, 100, 1),
				newSpend(account1, 100, 2),
				newSpend(account2, -100, 1),
			},
			gasOn:    false,
			noncesOn: true,
			finalBalances: map[string]*balances.Account{
				account1: newAccount(account1, 0, 2),
				account2: newAccount(account2, 0, 1),
			},
			err: nil,
		},
		{
			name: "gas and nonces off",
			spends: []*balances.Spend{
				newSpend(account1, 100, 1),
				newSpend(account1, 100, 2),
				newSpend(account2, -100, 1),
			},
			gasOn:    false,
			noncesOn: false,
			finalBalances: map[string]*balances.Account{
				account1: newAccount(account1, 0, 0),
				account2: newAccount(account2, 0, 0),
			},
			err: nil,
		},
		{
			name: "gas and nonces on",
			spends: []*balances.Spend{
				newSpend(account1, 100, 1),
			},
			gasOn:         true,
			noncesOn:      true,
			finalBalances: map[string]*balances.Account{},
			err:           balances.ErrInsufficientFunds,
		},
		{
			name:     "no account",
			spends:   []*balances.Spend{},
			gasOn:    false,
			noncesOn: false,
			finalBalances: map[string]*balances.Account{
				account1: newAccount(account1, 0, 0),
			},
			err: nil,
		},
		{
			name: "invalid nonce",
			spends: []*balances.Spend{
				newSpend(account1, 100, 1),
				newSpend(account1, 100, 3),
				newSpend(account2, -100, 1),
			},
			gasOn:    false,
			noncesOn: true,
			finalBalances: map[string]*balances.Account{
				account1: newAccount(account1, 0, 1),
				account2: newAccount(account2, 0, 1),
			},
			err: balances.ErrInvalidNonce,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()

			ar, td, err := accountTesting.NewTestAccountStore(ctx, balances.WithGasCosts(tc.gasOn), balances.WithNonces(tc.noncesOn))
			if err != nil {
				t.Fatalf("unexpected error: %s", err)
			}
			defer td()

			errs := []error{}
			for _, spend := range tc.spends {
				err := ar.Spend(ctx, spend)
				if err != nil {
					errs = append(errs, err)
				}
			}
			assertErr(t, errs, tc.err)

			for address, expectedBalance := range tc.finalBalances {
				account, err := ar.GetAccount(ctx, address)
				if err != nil {
					t.Fatalf("unexpected error: %s", err)
				}

				assert.Equal(t, expectedBalance.Balance, account.Balance, "expected balance %s, got %s", expectedBalance, account.Balance)
				assert.Equal(t, expectedBalance.Nonce, account.Nonce, "expected nonce %d, got %d", expectedBalance.Nonce, account.Nonce)
			}
		})
	}
}

func newSpend(address string, amount int64, nonce int64) *balances.Spend {
	return &balances.Spend{
		AccountAddress: address,
		Amount:         big.NewInt(amount),
		Nonce:          nonce,
	}
}

func newAccount(address string, balance int64, nonce int64) *balances.Account {
	return &balances.Account{
		Address: address,
		Balance: big.NewInt(balance),
		Nonce:   nonce,
	}
}

func assertErr(t *testing.T, errs []error, target error) {
	if target == nil {
		if len(errs) > 0 {
			t.Fatalf("expected no error, got %s", errs)
		}
		return
	}

	contains := false
	for _, err := range errs {
		if errors.Is(err, target) {
			contains = true
		}
	}

	if !contains {
		t.Fatalf("expected error %s, got %s", target, errs)
	}
}