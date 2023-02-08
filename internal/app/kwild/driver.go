package kwild

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"kwil/pkg/databases"
	"kwil/pkg/kclient"
	"kwil/pkg/types/data_types/any_type"
	"strings"
	"sync"
)

// Driver is a driver for the grpc client for integration tests
type Driver struct {
	cfg *kclient.Config

	connOnce sync.Once
	conn     *grpc.ClientConn
	client   *kclient.Client
}

func NewDriver(cfg *kclient.Config) *Driver {
	return &Driver{
		cfg: cfg,
	}
}

func (d *Driver) DeployDatabase(ctx context.Context, db *databases.Database[[]byte]) error {
	client, err := d.getClient(ctx)
	if err != nil {
		return fmt.Errorf("failed to create client: %w", err)
	}

	_, err = client.DeployDatabase(ctx, db)
	return err
}

func (d *Driver) DatabaseShouldExists(ctx context.Context, dbName string) error {
	client, err := d.getClient(ctx)
	if err != nil {
		return fmt.Errorf("failed to create client: %w", err)
	}

	schema, err := client.GetDatabaseSchema(ctx, dbName)
	if err != nil {
		return fmt.Errorf("failed to get database schema: %w", err)
	}

	if strings.ToLower(schema.Owner) == strings.ToLower(d.cfg.Fund.GetAccountAddress()) && schema.Name == dbName {
		return nil
	} else {
		return fmt.Errorf("database does not exist")
	}
}

func (d *Driver) ExecuteQuery(ctx context.Context, dbName string, queryName string, queryInputs []any) error {
	client, err := d.getClient(ctx)
	if err != nil {
		return fmt.Errorf("failed to create client: %w", err)
	}

	ins := make([]anytype.KwilAny, len(queryInputs))
	for i := 0; i < len(queryInputs); i++ {
		ins[i], err = anytype.New(queryInputs[i])
		if err != nil {
			return fmt.Errorf("failed to create query input: %w", err)
		}
	}

	_, err = client.ExecuteDatabase(ctx, dbName, queryName, ins)
	return err
}

func (d *Driver) DropDatabase(ctx context.Context, dbName string) error {
	client, err := d.getClient(ctx)
	if err != nil {
		return fmt.Errorf("failed to create client: %w", err)
	}

	_, err = client.DropDatabase(ctx, dbName)
	return err
}

func (d *Driver) Close() error {
	if d.conn != nil {
		return d.conn.Close()
	}
	return nil
}

func (d *Driver) getClient(ctx context.Context) (*kclient.Client, error) {
	var err error
	d.connOnce.Do(func() {
		d.client, err = kclient.New(ctx, d.cfg)
		if err != nil {
			return
		}
	})
	return d.client, err
}
