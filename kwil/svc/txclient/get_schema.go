package txclient

import (
	"context"
	"fmt"
	"kwil/x/proto/commonpb"
	"kwil/x/proto/txpb"
	"kwil/x/types/databases"
	"kwil/x/utils/serialize"
)

func (c *client) GetSchema(ctx context.Context, db *databases.DatabaseIdentifier) (*databases.Database, error) {
	res, err := c.txs.GetSchema(ctx, &txpb.GetSchemaRequest{
		Owner:    db.Owner,
		Database: db.Name,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get schema: %w", err)
	}

	return convertDatabase(res.Database)
}

func (c *client) GetSchemaById(ctx context.Context, id string) (*databases.Database, error) {
	res, err := c.txs.GetSchemaById(ctx, &txpb.GetSchemaByIdRequest{
		Id: id,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get schema: %w", err)
	}

	return convertDatabase(res.Database)
}

func convertDatabase(db *commonpb.Database) (*databases.Database, error) {
	// convert tables
	// convert response to database
	dbRes, err := serialize.Convert[commonpb.Database, databases.Database](db)
	if err != nil {
		return nil, fmt.Errorf("failed to convert response: %w", err)
	}

	return dbRes, nil
}