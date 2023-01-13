package txclient

import (
	"context"
	"fmt"
	"kwil/x/proto/commonpb"
	"kwil/x/proto/txpb"
	"kwil/x/types/transactions"
	"kwil/x/utils/serialize"
)

func (c *client) Broadcast(ctx context.Context, tx *transactions.Transaction) (*transactions.Response, error) {
	// convert transaction to proto
	pbTx, err := serialize.Convert[transactions.Transaction, commonpb.Tx](tx)
	if err != nil {
		return nil, fmt.Errorf("failed to convert transaction: %w", err)
	}

	res, err := c.txs.Broadcast(ctx, &txpb.BroadcastRequest{Tx: pbTx})
	if err != nil {
		return nil, fmt.Errorf("failed to broadcast transaction: %w", err)
	}

	// convert response to transaction
	txRes, err := serialize.Convert[txpb.BroadcastResponse, transactions.Response](res)
	if err != nil {
		return nil, fmt.Errorf("failed to convert response: %w", err)
	}

	return txRes, nil
}