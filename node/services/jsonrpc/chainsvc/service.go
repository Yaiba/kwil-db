package chainsvc

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/kwilteam/kwil-db/config"
	"github.com/kwilteam/kwil-db/core/log"
	jsonrpc "github.com/kwilteam/kwil-db/core/rpc/json"
	chainjson "github.com/kwilteam/kwil-db/core/rpc/json/chain"
	userjson "github.com/kwilteam/kwil-db/core/rpc/json/user"
	ktypes "github.com/kwilteam/kwil-db/core/types"
	chaintypes "github.com/kwilteam/kwil-db/core/types/chain"
	rpcserver "github.com/kwilteam/kwil-db/node/services/jsonrpc"
	nodetypes "github.com/kwilteam/kwil-db/node/types"
	"github.com/kwilteam/kwil-db/version"
)

const (
	apiVerMajor = 0
	apiVerMinor = 1 // should it be 2 ?
	apiVerPatch = 0

	serviceName = "chain"
)

// API version log
//
// apiVerMinor = 1 initial api

var (
	apiSemver = fmt.Sprintf("%d.%d.%d", apiVerMajor, apiVerMinor, apiVerPatch)
)

// Node specifies the methods required for chain service to interact with the blockchain.
type Node interface {
	// BlockByHeight returns block info at height. If height=0, the latest block will be returned.
	BlockByHeight(height int64) (ktypes.Hash, *ktypes.Block, ktypes.Hash, error)
	BlockByHash(hash ktypes.Hash) (*ktypes.Block, ktypes.Hash, error)
	BlockResultByHash(hash ktypes.Hash) ([]ktypes.TxResult, error)
	ChainTx(hash ktypes.Hash) (*chaintypes.ChainTx, error)
	BlockHeight() int64
	ChainUnconfirmedTx(limit int) (int, []nodetypes.NamedTx)
	ConsensusParams() *ktypes.ConsensusParams
}

type Validators interface {
	GetValidators() []*ktypes.Validator
}

type Service struct {
	log        log.Logger
	genesisCfg *config.GenesisConfig
	voting     Validators
	blockchain Node // node is the local node that can accept transactions.
}

func NewService(log log.Logger, blockchain Node, voting Validators, genesisCfg *config.GenesisConfig) *Service {
	return &Service{
		log:        log,
		genesisCfg: genesisCfg,
		voting:     voting,
		blockchain: blockchain,
	}
}

func (svc *Service) Name() string {
	return serviceName
}

func (svc *Service) Health(ctx context.Context) (detail json.RawMessage, happy bool) {
	healthResp, jsonErr := svc.HealthMethod(ctx, &chainjson.HealthRequest{})
	if jsonErr != nil { // unable to even perform the health check
		// This is not for a JSON-RPC client.
		svc.log.Error("health check failure", "error", jsonErr)
		resp, _ := json.Marshal(struct {
			Healthy bool `json:"healthy"`
		}{}) // omit everything else since
		return resp, false
	}

	resp, _ := json.Marshal(healthResp)

	return resp, healthResp.Healthy
}

func (svc *Service) Methods() map[jsonrpc.Method]rpcserver.MethodDef {
	return map[jsonrpc.Method]rpcserver.MethodDef{
		chainjson.MethodVersion: rpcserver.MakeMethodDef(verHandler,
			"retrieve the API version of the chain service",
			"service info including semver and kwild version"),
		chainjson.MethodHealth: rpcserver.MakeMethodDef(svc.HealthMethod,
			"retrieve the health status of the chain service",
			"health status of the service"),
		chainjson.MethodBlock: rpcserver.MakeMethodDef(svc.Block,
			"retrieve certain block info",
			"block information at a certain height"),
		chainjson.MethodBlockResult: rpcserver.MakeMethodDef(svc.BlockResult,
			"retrieve certain block result info",
			"block result information at a certain height"),
		chainjson.MethodTx: rpcserver.MakeMethodDef(svc.Tx,
			"retrieve certain transaction info",
			"transaction information at a certain hash"),
		chainjson.MethodGenesis: rpcserver.MakeMethodDef(svc.Genesis,
			"retrieve the genesis info",
			"genesis information"),
		chainjson.MethodConsensusParams: rpcserver.MakeMethodDef(svc.ConsensusParams,
			"retrieve the consensus parameers",
			"consensus parameters"),
		chainjson.MethodValidators: rpcserver.MakeMethodDef(svc.Validators,
			"retrieve validator info at certain height",
			"validator information at certain height"),
		chainjson.MethodUnconfirmedTxs: rpcserver.MakeMethodDef(svc.UnconfirmedTxs,
			"retrieve unconfirmed txs",
			"unconfirmed txs"),
	}
}

func (svc *Service) HealthMethod(_ context.Context, _ *chainjson.HealthRequest) (*chainjson.HealthResponse, *jsonrpc.Error) {
	return &chainjson.HealthResponse{
		ChainID: svc.genesisCfg.ChainID,
		Height:  svc.blockchain.BlockHeight(),
		Healthy: true,
	}, nil
}

func verHandler(context.Context, *userjson.VersionRequest) (*userjson.VersionResponse, *jsonrpc.Error) {
	return &userjson.VersionResponse{
		Service:     serviceName,
		Version:     apiSemver,
		Major:       apiVerMajor,
		Minor:       apiVerMinor,
		Patch:       apiVerPatch,
		KwilVersion: version.KwilVersion,
	}, nil
}

// Block returns block information either by block height or block hash.
// If both provided, block hash will be used.
func (svc *Service) Block(_ context.Context, req *chainjson.BlockRequest) (*chainjson.BlockResponse, *jsonrpc.Error) {
	if req.Height < 0 {
		return nil, jsonrpc.NewError(jsonrpc.ErrorInvalidParams, "height cannot be negative", nil)
	}

	// prioritize req.Hash over req.Height
	if !req.Hash.IsZero() {
		block, appHash, err := svc.blockchain.BlockByHash(req.Hash)
		if err != nil {
			svc.log.Error("block by hash", "hash", req.Hash, "error", err)
			return nil, jsonrpc.NewError(jsonrpc.ErrorNodeInternal, "failed to get block", nil)
		}

		return &chainjson.BlockResponse{
			Header:    (*chainjson.BlockHeader)(block.Header),
			Txns:      block.Txns,
			Signature: block.Signature,
			Hash:      req.Hash,
			AppHash:   appHash,
		}, nil
	}

	blockHash, block, appHash, err := svc.blockchain.BlockByHeight(req.Height)
	svc.log.Error("block by height", "height", req.Height, "hash", req.Hash, "error", err)
	if err != nil {
		return nil, jsonrpc.NewError(jsonrpc.ErrorNodeInternal, "failed to get block", nil)
	}

	return &chainjson.BlockResponse{
		Header:    (*chainjson.BlockHeader)(block.Header),
		Txns:      block.Txns,
		Signature: block.Signature,
		Hash:      blockHash,
		AppHash:   appHash,
	}, nil
}

// BlockResult returns block result either by block height or bloch hash.
// If both provided, block hash will be used.
func (svc *Service) BlockResult(_ context.Context, req *chainjson.BlockResultRequest) (*chainjson.BlockResultResponse, *jsonrpc.Error) {
	if req.Height < 0 {
		return nil, jsonrpc.NewError(jsonrpc.ErrorInvalidParams, "height cannot be negative", nil)
	}

	if !req.Hash.IsZero() {
		block, _, err := svc.blockchain.BlockByHash(req.Hash)
		if err != nil {
			svc.log.Error("block by hash", "hash", req.Hash, "error", err)
			return nil, jsonrpc.NewError(jsonrpc.ErrorNodeInternal, "failed to get block: "+err.Error(), nil)
		}

		txResults, err := svc.blockchain.BlockResultByHash(req.Hash)
		if err != nil {
			svc.log.Error("block result by hash", "hash", req.Hash, "error", err)
			return nil, jsonrpc.NewError(jsonrpc.ErrorNodeInternal, "failed to get block result: "+err.Error(), nil)
		}

		return &chainjson.BlockResultResponse{
			Height:    block.Header.Height,
			TxResults: txResults,
		}, nil
	}

	blockHash, block, _, err := svc.blockchain.BlockByHeight(req.Height)
	svc.log.Error("block by height", "height", req.Height, "hash", req.Hash, "error", err)
	if err != nil {
		return nil, jsonrpc.NewError(jsonrpc.ErrorNodeInternal, "failed to get block", nil)
	}

	txResults, err := svc.blockchain.BlockResultByHash(blockHash)
	if err != nil {
		svc.log.Error("block result by hash", "hash", req.Hash, "error", err)
		return nil, jsonrpc.NewError(jsonrpc.ErrorNodeInternal, "failed to get block result: "+err.Error(), nil)
	}

	return &chainjson.BlockResultResponse{
		Height:    block.Header.Height,
		TxResults: txResults,
	}, nil
}

// Tx returns a transaction by hash.
func (svc *Service) Tx(_ context.Context, req *chainjson.TxRequest) (*chaintypes.ChainTx, *jsonrpc.Error) {
	if req.Hash.IsZero() {
		return nil, jsonrpc.NewError(jsonrpc.ErrorInvalidParams, "hash is required", nil)
	}

	tx, err := svc.blockchain.ChainTx(req.Hash)
	if err != nil {
		svc.log.Error("tx by hash", "hash", req.Hash, "error", err)
		return nil, jsonrpc.NewError(jsonrpc.ErrorNodeInternal, "failed to get tx: "+err.Error(), nil)
	}

	return tx, nil
}

func (svc *Service) Genesis(ctx context.Context, _ *chainjson.GenesisRequest) (*chainjson.GenesisResponse, *jsonrpc.Error) {
	return &chainjson.GenesisResponse{
		ChainID:          svc.genesisCfg.ChainID,
		Leader:           svc.genesisCfg.Leader,
		Validators:       svc.genesisCfg.Validators,
		MaxBlockSize:     svc.genesisCfg.MaxBlockSize,
		JoinExpiry:       svc.genesisCfg.JoinExpiry,
		VoteExpiry:       svc.genesisCfg.VoteExpiry,
		DisabledGasCosts: svc.genesisCfg.DisabledGasCosts,
		MaxVotesPerTx:    svc.genesisCfg.MaxVotesPerTx,
	}, nil
}

func (svc *Service) ConsensusParams(_ context.Context, _ *chainjson.ConsensusParamsRequest) (*chainjson.ConsensusParamsResponse, *jsonrpc.Error) {
	return (*chainjson.ConsensusParamsResponse)(svc.blockchain.ConsensusParams()), nil
}

// Validators returns validator set at certain height. Default to latest height.
func (svc *Service) Validators(_ context.Context, _ *chainjson.ValidatorsRequest) (*chainjson.ValidatorsResponse, *jsonrpc.Error) {
	// NOTE: should be able to get validator set at req.Height
	vals := svc.voting.GetValidators()

	pbValidators := make([]*ktypes.Validator, len(vals))
	for i, vi := range vals {
		pbValidators[i] = &ktypes.Validator{
			Role:   vi.Role,
			PubKey: vi.PubKey,
			Power:  vi.Power,
		}
	}

	return &chainjson.ValidatorsResponse{
		Height:     svc.blockchain.BlockHeight(),
		Validators: nil,
	}, nil
}

// UnconfirmedTxs returns the unconfirmed txs. Default return 10 txs, max return 50 txs.
func (svc *Service) UnconfirmedTxs(_ context.Context, req *chainjson.UnconfirmedTxsRequest) (*chainjson.UnconfirmedTxsResponse, *jsonrpc.Error) {
	if req.Limit < 0 {
		return nil, jsonrpc.NewError(jsonrpc.ErrorInvalidParams, "invalid limit", nil)
	}
	if req.Limit > 50 {
		req.Limit = 50
	}
	if req.Limit == 0 {
		req.Limit = 10
	}
	total, txs := svc.blockchain.ChainUnconfirmedTx(req.Limit)
	return &chainjson.UnconfirmedTxsResponse{
		Total: total,
		Txs:   convertNamedTxs(txs),
	}, nil
}

// The admin Service must be usable as a Svc registered with a JSON-RPC Server.
var _ rpcserver.Svc = (*Service)(nil)

func convertNamedTxs(txs []nodetypes.NamedTx) []chainjson.NamedTx {
	res := make([]chainjson.NamedTx, len(txs))
	for i, tx := range txs {
		res[i] = chainjson.NamedTx{
			Hash: tx.Hash,
			Tx:   tx.Tx,
		}
	}
	return res
}