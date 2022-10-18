package apisvc

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"kwil/pkg/types/chain/pricing"
	"kwil/x/deposits"
	"kwil/x/logx"
	"kwil/x/proto/apipb"
)

type Service struct {
	apipb.UnimplementedKwilServiceServer

	ds      deposits.Deposits
	log     logx.Logger
	pricing pricing.PriceBuilder
	cc      ContractClient
}

type ContractClient interface {
	ReturnFunds(ctx context.Context, recip common.Address, amt *big.Int, fee *big.Int) (*types.Transaction, error)
}

func NewService(ds deposits.Deposits, p pricing.PriceBuilder, cc ContractClient) *Service {
	return &Service{
		ds:      ds,
		pricing: p,
		log:     logx.New(),
		cc:      cc,
	}
}

// validateBalances checks to ensure that the sender has enough funds to cover the fee.
// It also checks to ensure that the fee is not too low.
// Finally, it returns what the new balance should be if the operation is to be executed.
// It also returns an error if the amount is not enough
func (s *Service) validateBalances(from *string, op *int32, cr *int32, fe *big.Int) bool {

	// get the cost of the operation
	c := s.pricing.Operation(byte(*op)).Crud(byte(*cr)).Build()

	// convert cost from int64 to big.Int
	cost := big.NewInt(c)

	// compare the cost to what is sent
	if cost.Cmp(fe) > 0 {
		s.log.Debug("fee is too low for the requested operation")
		return false
	}

	return true
}
