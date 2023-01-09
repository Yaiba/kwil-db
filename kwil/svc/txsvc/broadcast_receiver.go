package txsvc

import (
	"context"
	"fmt"
	"kwil/x/proto/txpb"
	"kwil/x/transactions"
	"kwil/x/transactions/utils"
)

// Broadcast handles broadcasted transactions
func (s *Service) Broadcast(ctx context.Context, req *txpb.BroadcastRequest) (*txpb.BroadcastResponse, error) {
	// convert the transaction
	tx := utils.TxFromMsg(req.GetTx())
	err := tx.Verify() // verify verifies the hash and signature
	if err != nil {
		return nil, fmt.Errorf("failed to verify transaction: %w", err)
	}

	// handle the transaction according to its type
	switch tx.PayloadType {
	case transactions.DEPLOY_DATABASE:
		return s.handleDeployDatabase(ctx, tx)
	case transactions.MODIFY_DATABASE:
		return nil, fmt.Errorf("not implemented")
	case transactions.DROP_DATABASE:
		return s.handleDropDatabase(ctx, tx)
	case transactions.EXECUTE_QUERY:
		return s.handleExecution(ctx, tx)
	case transactions.WITHDRAW:
		return nil, fmt.Errorf("not implemented")
	default:
		return nil, fmt.Errorf("invalid payload type")
	}
}