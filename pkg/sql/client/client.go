/*
Package Client provides a useful interface for SQLite databases.

It is essentially a convenience wrapper around pkg/engine/sql/sqlite.

In the future, this can likely just be a part of the sqlite package, however
that package is currently very stable and I don't want to break it.  This
package is also pretty experimental, so it's best to keep it separate for now.

The primary purpose of this package is to create a minimal interface for
interacting with SQLite databases.
*/

package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/kwilteam/kwil-db/pkg/log"

	"github.com/kwilteam/kwil-db/pkg/sql/sqlite"
)

type SqliteClient struct {
	// conn is the underlying connection to the SQLite database.
	conn *sqlite.Connection

	// log is self-explanatory.
	log log.Logger
}

func NewSqliteStore(name string, opts ...SqliteOpts) (*SqliteClient, error) {
	optns := &options{
		log:  log.NewNoOp(),
		path: defaultPath,
		name: name,
	}

	for _, opt := range opts {
		opt(optns)
	}

	conn, err := sqlite.OpenConn(optns.name,
		sqlite.WithPath(optns.path),
		sqlite.WithLogger(optns.log),
	)
	if err != nil {
		return nil, err
	}

	return WrapSqliteConn(conn, optns.log)
}

func WrapSqliteConn(conn *sqlite.Connection, logger log.Logger) (*SqliteClient, error) {
	clnt := &SqliteClient{
		conn: conn,
		log:  logger,
	}

	err := clnt.conn.EnableForeignKey()
	if err != nil {
		return nil, fmt.Errorf("failed to enable foreign keys: %w", err)
	}

	return clnt, nil
}

// Execute executes a statement.
func (s *SqliteClient) Execute(ctx context.Context, stmt string, args map[string]any) error {
	return s.conn.Execute(stmt, args)
}

// Query executes a query and returns the result.
func (s *SqliteClient) Query(ctx context.Context, query string, args map[string]any) ([]map[string]any, error) {
	execOpts := []sqlite.ExecOption{}

	if args != nil {
		execOpts = append(execOpts, sqlite.WithNamedArgs(args))
	}

	results, err := s.conn.Query(ctx, query, execOpts...)
	if err != nil {
		return nil, err
	}

	return NewCursor(results).Export()
}

// Prepare prepares a statement for execution, and returns a Statement.
func (s *SqliteClient) Prepare(stmt string) (*Statement, error) {
	sqliteStmt, err := s.conn.Prepare(stmt)
	if err != nil {
		return nil, err
	}

	return &Statement{
		stmt: sqliteStmt,
	}, nil
}

// TableExists checks if a table exists.
func (s *SqliteClient) TableExists(ctx context.Context, table string) (bool, error) {
	return s.conn.TableExists(ctx, table)
}

// Close closes the connection to the database.
func (s *SqliteClient) Close() error {
	return s.conn.Close()
}

func (s *SqliteClient) Delete() error {
	return s.conn.Delete()
}

func (s *SqliteClient) Savepoint() (*Savepoint, error) {
	sp, err := s.conn.Savepoint()
	if err != nil {
		return nil, err
	}

	return &Savepoint{
		sp: sp,
	}, nil
}

// ResultsfromReader reads the results from a reader and returns them as an array of maps.
// WARNING: this will not convert byte arrays to strings, so you will need to do that yourself.
func ResultsfromReader(reader io.Reader) ([]map[string]any, error) {
	// Declare an empty array of map
	var rows []map[string]interface{}

	// Decode the JSON stream
	dec := json.NewDecoder(reader)
	for {
		var row []map[string]interface{}
		if err := dec.Decode(&row); err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		if row == nil {
			continue
		}

		rows = append(rows, row...)
	}

	return rows, nil
}