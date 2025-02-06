package erc20reward

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"math"
	"math/big"
	"strings"

	ethAccounts "github.com/ethereum/go-ethereum/accounts"
	ethCommon "github.com/ethereum/go-ethereum/common"

	kcommon "github.com/kwilteam/kwil-db/common"
	"github.com/kwilteam/kwil-db/core/types"
	pc "github.com/kwilteam/kwil-db/extensions/precompiles"
	"github.com/kwilteam/kwil-db/node/exts/erc20reward/meta"
	"github.com/kwilteam/kwil-db/node/exts/erc20reward/reward"
	"github.com/kwilteam/kwil-db/node/types/sql"
)

// This Kwil extension makes the following assumptions:
// - All TXs to GnosisSafe wallet are through Kwil, so that this ext can easily
//   increment the nonce(GnosisSafe and Reward contract) instead of using Oracle
//   to sync from outside. TODO: maybe we can add extra methods to update nonce.

// This is how this extension works:
// 1. Extension will be used ba a Kwil App(A Kuneiform app). App calls 'issue_rewards' to issue certain number of Rewards to a user.
// 2. An Epoch will be proposed so all issued rewards will be aggregated, and a merkle tree will be generated.
// 3. The SignerService run by different operators will vote(sign) an Epoch whenever it sees one. If one vote reaches the quota,
//    a FinalizedReward will be created.
// 4. The PosterService will post FinalizedReward.

const (
	uint256Precision = 78

	txIssueRewardCounterKey = "issue_reward_counter"
)

type chainInfo struct {
	Name      string
	Etherscan string
}

func (c chainInfo) GetEtherscanAddr(contract string) string {
	return c.Etherscan + contract + "#writeContract"
}

var chainConvMap = map[string]chainInfo{
	"1": {
		Name:      "Ethereum",
		Etherscan: "https://etherscan.io/address/",
	},
	"11155111": {
		Name:      "Sepolia",
		Etherscan: "https://sepolia.etherscan.io/address/",
	},
}

func init() {
	err := pc.RegisterInitializer("erc20_rewards",
		func(ctx context.Context, service *kcommon.Service, db sql.DB, alias string, metadata map[string]any) (p pc.Precompile, err error) {
			chainID, contractAddr, contractNonce, signers, threshold, safeAddr, safeNonce, decimals, err := getMetadata(metadata)
			if err != nil {
				return p, fmt.Errorf("parse ext configuration: %w", err)
			}

			ext := &Erc20RewardExt{
				contractID:    meta.GenRewardContractID(chainID, contractAddr),
				alias:         alias,
				decimals:      decimals,
				ContractAddr:  contractAddr,
				SafeAddr:      safeAddr,
				ChainID:       chainID,
				Signers:       signers,
				Threshold:     threshold,
				ContractNonce: contractNonce,
				SafeNonce:     safeNonce,
			}

			rewardAmtDecimal, err := types.NewNumericType(uint256Precision-decimals, decimals)
			if err != nil {
				return p, fmt.Errorf("failed to create decimal type: %w", err)
			}

			return pc.Precompile{
				Methods: []pc.Method{ // NOTE: engine ensures 'resultFn' in Handler is always non-nil
					{
						// Supposed to be called by App
						Name:            "issue_reward",
						AccessModifiers: []pc.Modifier{pc.PUBLIC},
						Handler: func(ctx *kcommon.EngineContext, app *kcommon.App, inputs []any, resultFn func([]any) error) error {
							return ext.issueReward(ctx, app, inputs, resultFn)
						},
						Parameters: []pc.PrecompileValue{
							{Type: types.TextType, Nullable: false},
							{Type: rewardAmtDecimal, Nullable: false},
						},
					},
					{
						// Supposed to be called by Signer service
						// Returns epoch rewards after(non-include) after_height, in ASC order.
						Name:            "list_epochs",
						AccessModifiers: []pc.Modifier{pc.PUBLIC, pc.VIEW},
						Handler: func(ctx *kcommon.EngineContext, app *kcommon.App, inputs []any, resultFn func([]any) error) error {
							return ext.listEpochs(ctx, app, inputs, resultFn)
						},
						Parameters: []pc.PrecompileValue{
							{Type: types.IntType, Nullable: false}, // after height
							{Type: types.IntType, Nullable: false}, // limit
						},
						Returns: &pc.MethodReturn{
							IsTable:    true,
							Fields:     (&Epoch{}).UnpackTypes(rewardAmtDecimal),
							FieldNames: (&Epoch{}).UnpackColumns(),
						},
					},
					{
						// Supposed to be called by the SignerService, to verify the reward root.
						// Could be merged into 'list_epochs'
						// Returns pending rewards from(include) start_height to(include) end_height, in ASC order.
						// NOTE: Rewards of same address will be aggregated.
						Name:            "search_rewards", // maybe not useful for Signer Service.
						AccessModifiers: []pc.Modifier{pc.PUBLIC, pc.VIEW},
						Handler: func(ctx *kcommon.EngineContext, app *kcommon.App, inputs []any, resultFn func([]any) error) error {
							return ext.searchRewards(ctx, app, inputs, resultFn)
						},
						Parameters: []pc.PrecompileValue{
							{Type: types.IntType, Nullable: false}, // start height
							{Type: types.IntType, Nullable: false}, // end height
						},
						Returns: &pc.MethodReturn{
							IsTable:    true,
							Fields:     (&Reward{}).UnpackTypes(rewardAmtDecimal),
							FieldNames: (&Reward{}).UnpackColumns(),
						},
					},
					{
						// Supposed to be called by Kwil network.
						Name:            "propose_epoch",
						AccessModifiers: []pc.Modifier{pc.PUBLIC}, // TODO: make this SYSTEM or Private
						Handler: func(ctx *kcommon.EngineContext, app *kcommon.App, inputs []any, resultFn func([]any) error) error {
							return ext.proposeEpoch(ctx, app, inputs, resultFn)
						},
					},
					{
						// Supposed to be called by SignerService.
						Name:            "vote_epoch",
						AccessModifiers: []pc.Modifier{pc.PUBLIC},
						Handler: func(ctx *kcommon.EngineContext, app *kcommon.App, inputs []any, resultFn func([]any) error) error {
							return ext.voteEpochReward(ctx, app, inputs, resultFn)
						},
						Parameters: []pc.PrecompileValue{
							// TODO: change to uuid??
							{Type: types.ByteaType, Nullable: false}, // sign hash
							{Type: types.ByteaType, Nullable: false}, // signature
						},
					},
					{
						// Supposed to be called by PosterService.
						// Returns finalized rewards after(non-include) start_height, in ASC order.
						Name:            "list_finalized",
						AccessModifiers: []pc.Modifier{pc.PUBLIC, pc.VIEW},
						Handler: func(ctx *kcommon.EngineContext, app *kcommon.App, inputs []any, resultFn func([]any) error) error {
							return ext.listFinalized(ctx, app, inputs, resultFn)
						},
						Parameters: []pc.PrecompileValue{
							{Type: types.IntType, Nullable: false}, // after height
							{Type: types.IntType, Nullable: false}, // limit
						},
						Returns: &pc.MethodReturn{
							IsTable:    true,
							Fields:     (&FinalizedReward{}).UnpackTypes(rewardAmtDecimal),
							FieldNames: (&FinalizedReward{}).UnpackColumns(),
						},
					},
					{
						// Supposed to be called by PosterService ?? seems this is not PosterService wants
						// Returns finalized rewards from(include) latest, in DESC order.
						Name:            "latest_finalized",
						AccessModifiers: []pc.Modifier{pc.PUBLIC, pc.VIEW},
						Handler: func(ctx *kcommon.EngineContext, app *kcommon.App, inputs []any, resultFn func([]any) error) error {
							return ext.latestFinalized(ctx, app, inputs, resultFn)
						},
						Parameters: []pc.PrecompileValue{
							{Type: types.IntType, Nullable: true}, // limit, default to 1
						},
						Returns: &pc.MethodReturn{
							IsTable:    true,
							Fields:     (&FinalizedReward{}).UnpackTypes(rewardAmtDecimal),
							FieldNames: (&FinalizedReward{}).UnpackColumns(),
						},
					},
					//{
					//	// TODO
					//	// Supposed to be called by PosterService,
					//	// Returns finalized rewards whose safeNonce are newer than afterSafeNonce, in DESC order.
					//	Name:            "newer_finalized",
					//	AccessModifiers: []pc.Modifier{pc.PUBLIC, pc.VIEW},
					//	Handler: func(ctx *kcommon.EngineContext, app *kcommon.App, inputs []any, resultFn func([]any) error) error {
					//		return ext.latestFinalized(ctx, app, inputs, resultFn)
					//	},
					//	Parameters: []pc.PrecompileValue{
					//		{Type: types.IntType, Nullable: true}, // afterSafeNonce
					//		{Type: types.IntType, Nullable: true}, // limit, default to 1
					//	},
					//	Returns: &pc.MethodReturn{
					//		IsTable:    true,
					//		Fields:     (&FinalizedReward{}).UnpackTypes(rewardAmtDecimal),
					//		FieldNames: (&FinalizedReward{}).UnpackColumns(),
					//	},
					//},
					{
						// Supposed to be called by App/User
						Name:            "list_wallet_rewards",
						AccessModifiers: []pc.Modifier{pc.PUBLIC, pc.VIEW},
						Parameters: []pc.PrecompileValue{
							{Type: types.TextType, Nullable: false}, // wallet address
						},
						Returns: &pc.MethodReturn{
							IsTable:    true,
							Fields:     (&WalletReward{}).UnpackTypes(),
							FieldNames: (&WalletReward{}).UnpackColumns(),
						},
						Handler: func(ctx *kcommon.EngineContext, app *kcommon.App, inputs []any, resultFn func([]any) error) error {
							return ext.listWalletRewards(ctx, app, inputs, resultFn)
						},
					},
					{
						// Supposed to be called by App/User
						Name:            "claim_param",
						AccessModifiers: []pc.Modifier{pc.PUBLIC, pc.VIEW},
						Handler: func(ctx *kcommon.EngineContext, app *kcommon.App, inputs []any, resultFn func([]any) error) error {
							return ext.getClaimParam(ctx, app, inputs, resultFn)
						},
						Parameters: []pc.PrecompileValue{
							// TODO: change to uuid??
							{Type: types.ByteaType, Nullable: false}, // sign hash
							{Type: types.TextType, Nullable: false},  // wallet address
						},
						Returns: &pc.MethodReturn{
							IsTable: true,
							Fields: []pc.PrecompileValue{
								{Type: types.TextType, Nullable: false},
								{Type: types.TextType, Nullable: false},
								{Type: types.TextType, Nullable: false},
								{Type: types.TextType, Nullable: false},
								{Type: types.TextArrayType, Nullable: true},
							},
							FieldNames: []string{"recipient", "amount", "block_hash", "root", "proofs"},
						},
					},
					// TODO: modify posterFee, modify signers
					{
						Name:            "propose_settings",
						AccessModifiers: []pc.Modifier{pc.PUBLIC},
						Parameters: []pc.PrecompileValue{
							{Type: types.TextArrayType, Nullable: true}, // signers
							{Type: types.IntType, Nullable: true},       // threshold
							{Type: rewardAmtDecimal, Nullable: true},    // posterFee
						},
						Returns: &pc.MethodReturn{},
						Handler: func(ctx *kcommon.EngineContext, app *kcommon.App, inputs []any, resultFn func([]any) error) error {
							return ext.proposeSettings(ctx, app, inputs, resultFn)
						},
					},
					{
						Name:            "vote_settings",
						AccessModifiers: []pc.Modifier{pc.PUBLIC},
						Parameters: []pc.PrecompileValue{
							{Type: types.UUIDType, Nullable: false},  // proposal id
							{Type: types.ByteaType, Nullable: false}, // vote
							{Type: types.TextType, Nullable: false},  // wallet address,
						},
						Returns: &pc.MethodReturn{},
						Handler: func(ctx *kcommon.EngineContext, app *kcommon.App, inputs []any, resultFn func([]any) error) error {
							return ext.proposeSettings(ctx, app, inputs, resultFn)
						},
					},
				},
				OnStart: func(ctx context.Context, app *kcommon.App) error {
					tx, err := app.DB.BeginTx(ctx)
					if err != nil {
						return err
					}
					defer tx.Rollback(ctx)

					// OnStart is not called at a certain block height, nor in a transaction
					emptyEngineCtx := &kcommon.EngineContext{
						TxContext: &kcommon.TxContext{
							Ctx: ctx,
						}}

					// check if the reward contract exists
					contract, err := meta.GetRewardContract(emptyEngineCtx, app.Engine, app.DB, ext.ChainID, ext.ContractAddr)
					if err != nil {
						if !errors.Is(err, sql.ErrNoRows) {
							return err
						}
						return nil
					}

					// exist, use values from DB
					ext.contractID = meta.GenRewardContractID(ext.ChainID, ext.ContractAddr)
					ext.Signers = contract.Signers
					ext.Threshold = contract.Threshold
					ext.ContractNonce = contract.Nonce
					ext.SafeNonce = contract.SafeNonce
					return nil
				},
				OnUse: func(ctx *kcommon.EngineContext, app *kcommon.App) error {
					app.Service.Logger.Info("Register a new erc20_rewards contract",
						"chainID", ext.ChainID, "contractAddr", ext.ContractAddr, "alias", ext.alias)
					initRewardMeta := "USE IF NOT EXISTS erc20_rewards_meta as " + meta.ExtAlias + ";"
					err := app.Engine.Execute(ctx, app.DB, initRewardMeta, nil, nil)
					if err != nil {
						return err
					}

					_, err = app.Engine.Call(ctx, app.DB, meta.ExtAlias, "register",
						[]any{
							ext.ChainID,
							ext.ContractAddr,
							ext.ContractNonce,
							strings.Join(ext.Signers, ","),
							ext.Threshold,
							ext.SafeAddr,
							ext.SafeNonce,
						},
						nil)
					if err != nil {
						return err
					}

					return initTables(ctx, app, ext.alias)
				},
				OnUnuse: func(ctx *kcommon.EngineContext, app *kcommon.App) error {
					return nil
				},
			}, nil
		})
	if err != nil {
		panic(fmt.Errorf("failed to register erc20_rewards initializer: %w", err))
	}
}

func initTables(ctx *kcommon.EngineContext, app *kcommon.App, ns string) error {
	ctx.OverrideAuthz = true
	defer func() { ctx.OverrideAuthz = false }()

	err := app.Engine.Execute(ctx, app.DB, fmt.Sprintf(sqlInitTableErc20rwRewards, ns, meta.ExtAlias), nil, nil)
	if err != nil {
		return err
	}

	err = app.Engine.Execute(ctx, app.DB, fmt.Sprintf(sqlInitTableErc20rwEpochs, ns, meta.ExtAlias), nil, nil)
	if err != nil {
		return err
	}

	err = app.Engine.Execute(ctx, app.DB, fmt.Sprintf(sqlInitTableErc20rwEpochVotes, ns, ns, meta.ExtAlias), nil, nil)
	if err != nil {
		return err
	}

	err = app.Engine.Execute(ctx, app.DB, fmt.Sprintf(sqlInitTableErc20rwFinalizedRewards, ns, ns), nil, nil)
	if err != nil {
		return err
	}

	err = app.Engine.Execute(ctx, app.DB, fmt.Sprintf(sqlInitTableRecipientReward, ns, ns), nil, nil)
	if err != nil {
		return err
	}

	return nil
}

// Erc20RewardExt a struct that implements the precompiles.Instance interface
type Erc20RewardExt struct {
	// contractID identifies a reward extension Erc20RewardExt, it's the contractID in table
	// erc20rw_meta_contracts
	contractID *types.UUID
	// alias is the namespace/schema current extension resides
	alias string
	// BuiltinMode indicates whether this ext is initialized in builtin mode.
	//BuiltinMode bool
	// ContractAddr is the reward escrow EVM contract address.
	ContractAddr string
	// SafeAddr is the GnosisSafe wallet address that has permission to update
	// the reward escrow contract.
	SafeAddr string
	ChainID  int64
	decimals uint16 // the denotation of the reward token, most of the ERC20 are 18

	// Those are the parameters that could be updated over time.
	// Upon reload(node restart), should populate those values from DB.
	// Upon updates, those should also get updated.
	// NOTE: We assume they can only be modified through ext.
	Signers       []string
	Threshold     int64
	ContractNonce int64
	SafeNonce     int64
}

func getMetadata(metadata map[string]any) (chainID int64, contractAddr string,
	contractNonce int64, signers []string, threshold int64, safeAddr string,
	safeNonce int64, decimals uint16, err error) {
	var ok bool

	allKeys := []string{"chain_id", "contract_address", "contract_nonce",
		"threshold", "signers", "safe_address", "safe_nonce", "decimals"}
	for _, key := range allKeys {
		_, ok := metadata[key]
		if !ok {
			err = fmt.Errorf("missing %s", key)
			return
		}
	}

	chainID, ok = metadata["chain_id"].(int64)
	if !ok {
		err = fmt.Errorf("invalid chain_id")
		return
	}

	contractAddr, ok = metadata["contract_address"].(string)
	if !ok {
		err = fmt.Errorf("invalid contract_address")
		return
	}

	contractNonce, ok = metadata["contract_nonce"].(int64)
	if !ok {
		err = fmt.Errorf("invalid contract_nonce")
		return
	}

	if !ethCommon.IsHexAddress(contractAddr) {
		err = fmt.Errorf("invalid contract_address")
		return
	}

	signersStr, ok := metadata["signers"].(string)
	if !ok {
		err = fmt.Errorf("invalid signers")
		return
	}

	if len(signersStr) == 0 {
		err = fmt.Errorf("signers is empty")
		return
	}

	signers = strings.Split(signersStr, ",")
	for _, signer := range signers {
		if !ethCommon.IsHexAddress(signer) {
			err = fmt.Errorf("invalid signer")
			return
		}
	}

	threshold, ok = metadata["threshold"].(int64)
	if !ok {
		err = fmt.Errorf("invalid threshold")
		return
	}

	if threshold == 0 {
		err = fmt.Errorf("threshold is 0")
		return
	}

	if threshold > int64(len(signers)) {
		err = fmt.Errorf("threshold is larger than the number of signers")
		return
	}

	safeAddr, ok = metadata["safe_address"].(string)
	if !ok {
		err = fmt.Errorf("invalid safe_address")
		return
	}

	if !ethCommon.IsHexAddress(safeAddr) {
		err = fmt.Errorf("invalid safe_address")
	}

	safeNonce, ok = metadata["safe_nonce"].(int64)
	if !ok {
		err = fmt.Errorf("invalid safe_nonce")
		return
	}

	decimals64, ok := metadata["decimals"].(int64)
	if !ok {
		err = fmt.Errorf("invalid decimals")
		return
	}

	if decimals64 <= 0 {
		err = fmt.Errorf("decimals should be positive")
		return
	}

	if decimals64 > math.MaxUint16 {
		err = fmt.Errorf("decimals too large")
		return
	}

	decimals = uint16(decimals64)

	return
}

// scaleUpUint256 turns a decimal into uint256, i.e. (11.22, 4) -> 112200
func scaleUpUint256(amount *types.Decimal, decimals uint16) (*types.Decimal, error) {
	unit, err := types.ParseDecimal("1" + strings.Repeat("0", int(decimals)))
	if err != nil {
		return nil, fmt.Errorf("create decimal unit failed: %w", err)
	}

	n, err := types.DecimalMul(amount, unit)
	if err != nil {
		return nil, fmt.Errorf("expand amount decimal failed: %w", err)
	}

	err = n.SetPrecisionAndScale(uint256Precision, 0)
	if err != nil {
		return nil, fmt.Errorf("expand amount decimal failed: %w", err)
	}

	return n, nil
}

// scaleDownUint256 turns an uint256 to a decimal, i.e. (112200, 4) -> 11.22
func scaleDownUint256(amount *types.Decimal, decimals uint16) (*types.Decimal, error) {
	unit, err := types.ParseDecimal("1" + strings.Repeat("0", int(decimals)))
	if err != nil {
		return nil, fmt.Errorf("create decimal unit failed: %w", err)
	}

	n, err := types.DecimalDiv(amount, unit)
	if err != nil {
		return nil, fmt.Errorf("expand amount decimal failed: %w", err)
	}

	scale := n.Scale()
	if scale > decimals {
		scale = decimals
	}

	err = n.SetPrecisionAndScale(uint256Precision-decimals, scale)
	if err != nil {
		return nil, fmt.Errorf("expand amount decimal failed: %w", err)
	}

	return n, nil
}

func (h *Erc20RewardExt) issueReward(ctx *kcommon.EngineContext, app *kcommon.App, inputs []any, resultFn func([]any) error) error {
	wallet, ok := inputs[0].(string)
	if !ok {
		return fmt.Errorf("invalid wallet address")
	}

	if !ethCommon.IsHexAddress(wallet) {
		return fmt.Errorf("invalid wallet address")
	}

	amount, ok := inputs[1].(*types.Decimal)
	if !ok {
		return fmt.Errorf("invalid amount")
	}

	// require amount is positive
	zero, _ := types.ParseDecimal("0.0")
	r, err := types.DecimalCmp(zero, amount)
	if err != nil {
		return fmt.Errorf("invalid amount")
	}
	if r != -1 {
		return fmt.Errorf("invalid amount")
	}

	uint256Amount, err := scaleUpUint256(amount, h.decimals)
	if err != nil {
		return err
	}

	counter := 0
	c, exist := ctx.TxContext.Value(txIssueRewardCounterKey)
	if exist {
		counter = c.(int)
	}

	// not matter if db operations success, increase the counter
	defer func() {
		ctx.TxContext.SetValue(txIssueRewardCounterKey, counter+1)
	}()

	if err := IssueReward(ctx, app.Engine, app.DB, h.alias, &Reward{
		ID:         GenRewardID(wallet, uint256Amount.String(), ctx.TxContext.TxID, counter),
		Recipient:  wallet,
		Amount:     uint256Amount,
		CreatedAt:  ctx.TxContext.BlockContext.Height,
		ContractID: h.contractID,
	}); err != nil {
		return err
	}

	return nil
}

// searchRewards returns rewards between a starting height and ending height.
func (h *Erc20RewardExt) searchRewards(ctx *kcommon.EngineContext, app *kcommon.App, inputs []any, resultFn func([]any) error) error {
	startHeight, ok := inputs[0].(int64)
	if !ok {
		return fmt.Errorf("invalid start height")
	}

	if startHeight < 0 {
		return fmt.Errorf("invalid start height")
	}

	endHeight, ok := inputs[1].(int64)
	if !ok {
		return fmt.Errorf("invalid end height")
	}

	if endHeight < 0 {
		return fmt.Errorf("invalid end height")
	}

	if startHeight > endHeight {
		return fmt.Errorf("invalid start height")
	}

	if endHeight-startHeight > 10000 {
		return fmt.Errorf("search range too large")
	}

	rewards, err := SearchRewards(ctx, app.Engine, app.DB, h.alias, h.decimals, startHeight, endHeight)
	if err != nil {
		return err
	}

	for _, r := range rewards {
		err := resultFn(r.UnpackValues())
		if err != nil {
			return err
		}
	}

	return nil
}

// proposeEpoch proposes a new epoch of rewards that are pending from last batch
// to current height, it requires a correct Gnosis Safe wallet nonce.
// inputs[0] is the Gnosis Safe wallet nonce, which will be used to generate Gnosis Safe TX hash.
//
// This is supposed to be called by Network owner or Signer service??
// In either case, there should be just one proposer at a time.
// Seems reasonable to be called by KwilNetwork, as this action requires nothing
// but safeNonce(which could also be inferred inside the Extension, no need to be provided by caller)
// For simplicity, we just use safeNonce tracked by extension.
// NOTE: well, do we need to check permission? e.g., check the caller??
func (h *Erc20RewardExt) proposeEpoch(ctx *kcommon.EngineContext, app *kcommon.App, inputs []any, resultFn func([]any) error) error {
	endHeight := ctx.TxContext.BlockContext.Height
	blockHash := ctx.TxContext.BlockContext.Hash

	// get last finalized batch
	var lastEpochEndHeight int64
	finalizedReward, err := GetLatestFinalizedReward(ctx, app.Engine, app.DB, h.alias, h.decimals)
	if err != nil {
		return err
	}
	if finalizedReward != nil {
		lastEpochEndHeight = finalizedReward.EndHeight
	}

	epochStartHeight := lastEpochEndHeight + 1
	rewards, err := SearchRewards(ctx, app.Engine, app.DB, h.alias, h.decimals, epochStartHeight, endHeight)
	if err != nil {
		return err
	}

	if len(rewards) == 0 {
		return fmt.Errorf("no rewards")
	}

	recipients := make([]string, len(rewards))
	bigIntAmounts := make([]*big.Int, len(rewards))
	var totalAmount *types.Decimal // nil
	for i, r := range rewards {
		recipients[i] = r.Recipient

		if totalAmount == nil {
			totalAmount = r.Amount
		} else {
			totalAmount, err = types.DecimalAdd(totalAmount, r.Amount)
			if err != nil {
				return err
			}
		}
		bigIntAmounts[i] = r.Amount.BigInt()
	}

	// NOTE: since we don't have a limit on how many leafs(recipients) a tree can
	// have, this could be big
	jsonMtree, rootHash, err := reward.GenRewardMerkleTree(recipients, bigIntAmounts, h.ContractAddr, blockHash)
	if err != nil {
		return err
	}

	safeTxData, err := reward.GenPostRewardTxData(rootHash, totalAmount.BigInt())
	if err != nil {
		return err
	}

	// safeTxHash is the data that all signers will be signing(using personal_sign)
	_, safeTxHash, err := reward.GenGnosisSafeTx(h.ContractAddr, h.SafeAddr,
		0, safeTxData, h.ChainID, h.SafeNonce)
	if err != nil {
		return err
	}

	// NOTE: we save the digest of the msg, so it's fix length
	// well, safeTxHash should also be a fix length, we save the digest anyway.
	signHash := ethAccounts.TextHash(safeTxHash)

	ctx.OverrideAuthz = true
	defer func() { ctx.OverrideAuthz = false }()

	uint256TotalAmount, err := scaleUpUint256(totalAmount, h.decimals)
	if err != nil {
		return err
	}

	return app.Engine.Execute(ctx, app.DB, fmt.Sprintf(sqlNewEpoch, h.alias), map[string]any{
		"$id":            GenBatchRewardID(endHeight, signHash),
		"$start_height":  epochStartHeight,
		"$end_height":    endHeight,
		"$total_rewards": uint256TotalAmount,
		"$mtree_json":    []byte(jsonMtree),
		"$reward_root":   rootHash,
		"$safe_nonce":    h.SafeNonce,
		"$sign_hash":     signHash,
		"$contract_id":   h.contractID,
		"$block_hash":    blockHash[:],
		"$created_at":    ctx.TxContext.BlockContext.Height, // TODO: seems we can remove this field, it's always the same as end_height
	}, nil)
}

// listEpochs returns reward epochs starting from a given height.
// inputs[0] is the starting height, inputs[1] is the return size and default to 10.
func (h *Erc20RewardExt) listEpochs(ctx *kcommon.EngineContext, app *kcommon.App, inputs []any, resultFn func([]any) error) error {
	afterHeight, ok := inputs[0].(int64)
	if !ok {
		return fmt.Errorf("invalid after height")
	}

	if afterHeight < 0 { // or set to default
		return fmt.Errorf("invalid after height")
	}

	limit, ok := inputs[1].(int64)
	if !ok {
		return fmt.Errorf("invalid limit type")
	}

	if limit < 0 { // or set to default
		return fmt.Errorf("invalid limit")
	}

	if limit == 0 {
		limit = 1 // default to 1
	}

	if limit > 10 {
		limit = 10 // max to 10
	}
	rewards, err := ListEpochs(ctx, app.Engine, app.DB, h.alias, h.decimals, afterHeight, limit)
	if err != nil {
		return err
	}

	for _, r := range rewards {
		err := resultFn(r.UnpackValues())
		if err != nil {
			return err
		}
	}

	return nil
}

// voteEpochReward votes one epoch of rewards by providing correspond signature.
// inputs[0] is the data digest, inputs[1] is the signature.
// This is supposed to be called by Signer service.
func (h *Erc20RewardExt) voteEpochReward(ctx *kcommon.EngineContext, app *kcommon.App, inputs []any, resultFn func([]any) error) error {
	// verify the caller is signer
	_voter, err := meta.GetSigner(ctx, app.Engine, app.DB, ctx.TxContext.Caller, h.contractID)
	if err != nil {
		return fmt.Errorf("check signer: %w", err)
	}
	if _voter == nil {
		return fmt.Errorf("voter not allowed")
	}

	digest, ok := inputs[0].([]byte)
	if !ok {
		return fmt.Errorf("invalid safe tx hash")
	}

	if len(digest) == 0 {
		return fmt.Errorf("invalid safe tx hash")
	}

	signature, ok := inputs[1].([]byte)
	if !ok {
		return fmt.Errorf("invalid signature")
	}

	caller := ethCommon.HexToAddress(ctx.TxContext.Caller)
	err = reward.EthGnosisVerifyDigest(signature, digest, caller.Bytes())
	if err != nil {
		return err
	}

	// if already voted, skip
	var voted bool
	err = app.Engine.Execute(ctx, app.DB, fmt.Sprintf(sqlGetVoteBySignHash, h.alias, h.alias, meta.ExtAlias),
		map[string]any{
			"$sign_hash":      digest,
			"$signer_address": ctx.TxContext.Caller,
			"$contract_id":    h.contractID,
		},
		func(row *kcommon.Row) error {
			voted = true
			return nil
		})
	if err != nil {
		return err
	}
	if voted {
		return nil
	}

	ctx.OverrideAuthz = true
	defer func() { ctx.OverrideAuthz = false }()

	err = app.Engine.Execute(ctx, app.DB, fmt.Sprintf(sqlVoteEpochBySignHash, h.alias, h.alias, meta.ExtAlias),
		map[string]any{
			"$sign_hash":      digest,
			"$signer_address": ctx.TxContext.Caller,
			"$contract_id":    h.contractID,
			"$signature":      signature,
			"$created_at":     ctx.TxContext.BlockContext.Height,
		}, nil)
	if err != nil {
		return err
	}

	finalized, err := TryFinalizeEpochReward(ctx, app.Engine, app.DB, h.alias, h.contractID,
		digest, ctx.TxContext.BlockContext.Height, h.Threshold)
	if err != nil {
		return err
	}

	if finalized {
		h.SafeNonce += 1
	}

	return nil
}

// listFinalized returns finalized rewards.
// inputs[0] is the starting height, inputs[1] is the batch size and default to 10.
// This is supposed to be called by Poster service.
func (h *Erc20RewardExt) listFinalized(ctx *kcommon.EngineContext, app *kcommon.App, inputs []any, resultFn func([]any) error) error {
	afterHeight, ok := inputs[0].(int64)
	if !ok {
		return fmt.Errorf("invalid after height")
	}

	if afterHeight < 0 { // or set to default
		return fmt.Errorf("invalid after height")
	}

	limit, ok := inputs[1].(int64)
	if !ok {
		return fmt.Errorf("invalid limit type")
	}

	if limit < 0 { // or set to default
		return fmt.Errorf("invalid limit")
	}
	if limit == 0 {
		limit = 10 // default to 10
	}

	if limit > 50 {
		limit = 50 // max to 50
	}

	rewards, err := ListFinalizedRewards(ctx, app.Engine, app.DB, h.alias, h.decimals, afterHeight, limit)
	if err != nil {
		return err
	}

	for _, r := range rewards {
		err := resultFn(r.UnpackValues())
		if err != nil {
			return err
		}
	}

	return nil
}

// latestFinalized returns latest finalized rewards.
// inputs[0] is the size and default to 0, i.e. return the newest.
// This is supposed to be called by Poster service.
func (h *Erc20RewardExt) latestFinalized(ctx *kcommon.EngineContext, app *kcommon.App, inputs []any, resultFn func([]any) error) error {
	var limit int64 = 1
	if inputs[0] != nil {
		var ok bool
		limit, ok = inputs[0].(int64)
		if !ok {
			return fmt.Errorf("invalid limit type")
		}
	}

	if limit <= 0 { // or set to default
		limit = 1
	}

	if limit > 20 {
		limit = 20 // max to 20
	}

	rewards, err := ListLatestFinalizedRewards(ctx, app.Engine, app.DB, h.alias, h.decimals, limit)
	if err != nil {
		return err
	}

	for _, r := range rewards {
		err := resultFn(r.UnpackValues())
		if err != nil {
			return err
		}
	}

	return nil
}

func (h *Erc20RewardExt) listWalletRewards(ctx *kcommon.EngineContext, app *kcommon.App, inputs []any, resultFn func([]any) error) error {
	wallet, ok := inputs[0].(string)
	if !ok {
		return fmt.Errorf("invalid wallet address")
	}

	if !ethCommon.IsHexAddress(wallet) {
		return fmt.Errorf("invalid wallet address")
	}

	walletAddr := ethCommon.HexToAddress(wallet)

	partialWrs, err := GetWalletRewards(ctx, app.Engine, app.DB, h.alias, walletAddr.String())
	if err != nil {
		return err
	}

	wrs := make([]*WalletReward, len(partialWrs))

	for i, pwr := range partialWrs {
		treeRoot, proofs, _, bh, uint256AmtStr, err := reward.GetMTreeProof(pwr.mTreeJSON, walletAddr.String())
		if err != nil {
			return err
		}

		info, ok := chainConvMap[pwr.chainID]
		if !ok {
			return fmt.Errorf("internal bug: unknown chain id")
		}

		wrs[i] = &WalletReward{
			Chain:          info.Name,
			ChainID:        pwr.chainID,
			Contract:       pwr.contract,
			EtherScan:      info.GetEtherscanAddr(pwr.contract),
			CreatedAt:      pwr.createdAt,
			ParamRecipient: walletAddr.String(),
			ParamAmount:    uint256AmtStr,
			ParamBlockHash: toBytes32Str(bh),
			ParamRoot:      toBytes32Str(treeRoot),
			ParamProofs:    meta.Map(proofs, toBytes32Str),
		}
	}

	for _, r := range wrs {
		err := resultFn(r.UnpackValues())
		if err != nil {
			return err
		}
	}

	return nil
}

func toBytes32Str(bs []byte) string {
	return "0x" + hex.EncodeToString(bs)
}

// getClaimParam returns the claim parameters for given signHash and wallet address,
// User can use the parameters directly to call the `claimReward` method on reward contract.
// inputs[0] is the safeTxHash, inputs[1] is the user wallet address.
// This is supposed to be called by User.
func (h *Erc20RewardExt) getClaimParam(ctx *kcommon.EngineContext, app *kcommon.App, inputs []any, resultFn func([]any) error) error {
	signHash, ok := inputs[0].([]byte)
	if !ok {
		return fmt.Errorf("invalid signhash")
	}

	if len(signHash) == 0 {
		return fmt.Errorf("invalid signhash")
	}

	wallet, ok := inputs[1].(string)
	if !ok {
		return fmt.Errorf("invalid wallet address")
	}

	if !ethCommon.IsHexAddress(wallet) {
		return fmt.Errorf("invalid wallet address")
	}

	walletAddr := ethCommon.HexToAddress(wallet)

	mTreeJson, err := GetEpochMTreeBySignhash(ctx, app.Engine, app.DB, h.alias, signHash)
	if err != nil {
		return err
	}

	if mTreeJson == nil {
		return fmt.Errorf("no reward found")
	}

	treeRoot, proofs, _, bh, uint256Amt, err := reward.GetMTreeProof(string(mTreeJson), walletAddr.String())
	if err != nil {
		return err
	}

	return resultFn([]any{
		walletAddr.String(),
		uint256Amt,
		toBytes32Str(bh),
		toBytes32Str(treeRoot),
		meta.Map(proofs, toBytes32Str)})
}

// proposeSettings create a new proposition on the settings.
// NOTE: this should be idempotent; after processing the arguments, we always present
// the propose with FULL settings.
func (h *Erc20RewardExt) proposeSettings(ctx *kcommon.EngineContext, app *kcommon.App, inputs []any, resultFn func([]any) error) error {
	var signers []string
	var threshold int64
	var uint256PosterFee *types.Decimal

	if inputs[0] != nil {
		signers, ok := inputs[0].([]string)
		if !ok {
			return fmt.Errorf("invalid signers")
		}

		for _, signer := range signers {
			if !ethCommon.IsHexAddress(signer) {
				return fmt.Errorf("invalid signer")
			}
		}
	}

	if inputs[1] != nil {
		threshold, ok := inputs[1].(int64)
		if !ok {
			return fmt.Errorf("invalid threshold")
		}
	}

	if inputs[2] != nil {
		posterFee, ok := inputs[2].(*types.Decimal)
		if !ok {
			return fmt.Errorf("invalid poster fee")
		}

		// require amount is positive
		zero, _ := types.ParseDecimal("0.0")
		r, err := types.DecimalCmp(zero, posterFee)
		if err != nil {
			return fmt.Errorf("invalid poster fee")
		}
		if r != -1 {
			return fmt.Errorf("invalid poster fee")
		}

		uint256PosterFee, err := scaleUpUint256(posterFee, h.decimals)
		if err != nil {
			return err
		}
	}

	if len(signers) == 0 && threshold == 0 && uint256PosterFee == nil {
		return fmt.Errorf("no settings provided")
	}

	if len(signers) == 0 {
		// populate it with existing signers
	}

	if threshold == 0 {
		// populate it with existing threshold
	}

	if uint256PosterFee == nil {
		// populate it with existing fee
	}

	// if no setting changes, return error

	err := createSettingProposal()
	if err != nil {
		return err
	}

	return nil
}

func createSettingProposal() error {}

var (
	sqlInitTableSettingProposals = `
-- setting_proposals holds the propositions on the settings of the reward contract.
-- If the system(or global) safe nonce advanced before the proposal is passed, then the proposal is invalidated.
CREATE TABLE IF NOT EXISTS setting_proposals (
    id UUID PRIMARY KEY,
    instance_ID UUID NOT NULL references %s.erc20rw_meta_contracts(id) ON UPDATE CASCADE ON DELETE CASCADE,
    signers TEXT[] NOT NULL,
    threshold INT8 NOT NULL,
    poster_fee NUMERIC(78,0) NOT NULL,
    safe_nonce INT8 NOT NULL, -- the safe nonce associated with this proposal
    created_at INT8 NOT NULL,
    needed_votes INT8 NOT NULL, -- the number of votes needed to pass the proposal
    -- the following fields will be updated upon every vote, once enough votes are received
    voters TEXT[] NOT NULL, -- the voters of the proposal
    signatures BYTEA[] NOT NULL, -- the signatures of the voters
    passed_at INT8 NOT NULL -- the block height when the proposal passed
)`

	sqlCreateSettingProposal = ``
	sqlVoteSettingProposal   = ``
	sqlGetSettingProposal    = ``
)
