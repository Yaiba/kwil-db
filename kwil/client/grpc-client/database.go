package grpc_client

import (
	"context"
	"fmt"
	"kwil/x/execution/clean"
	"kwil/x/execution/validation"
	"kwil/x/types/databases"
	"kwil/x/types/transactions"
	"strings"
)

func (c *Client) DeployDatabase(ctx context.Context, db *databases.Database) (*transactions.Response, error) {
	clean.CleanDatabase(db)

	// validate the database
	err := validation.ValidateDatabase(db)
	if err != nil {
		return nil, fmt.Errorf("error on database: %w", err)
	}

	if !strings.EqualFold(db.Owner, c.Config.Address) {
		return nil, fmt.Errorf("database owner must be the same as the current account.  Owner: %s, Account: %s", db.Owner, c.Config.Address)
	}

	// build tx
	tx, err := c.BuildTransaction(ctx, transactions.DEPLOY_DATABASE, db, c.Config.PrivateKey)
	if err != nil {
		return nil, err
	}

	return c.Txs.Broadcast(ctx, tx)
}