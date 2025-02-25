package pg

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"sync"

	"github.com/kwilteam/kwil-db/common/sql"
	"github.com/kwilteam/kwil-db/core/utils/random"

	"github.com/jackc/pgx/v5"
)

// DB is a session-aware wrapper that creates and stores a write Tx on request,
// and provides top level Exec/Set methods that error if no Tx exists. This
// design prevents any out-of-session write statements from executing, and makes
// uncommitted reads explicit (and impossible in the absence of an active
// transaction).
//
// This type is tailored to use in kwild in the following ways:
//
//  1. Controlled transactional interaction that requires beginning a
//     transaction before using the Exec method, unless put in "autocommit" mode
//     using the AutoCommit method. Use of the write connection when not
//     executing a block's transactions is prevented.
//
//  2. Using an underlying connection pool, with multiple readers and a single
//     write connection to ensure all uses of Execute operate on the active
//     transaction.
//
//  3. Emulating SQLite changesets by collecting WAL data for updates from a
//     dedicated logical replication connection and slot. The Precommit method
//     is used to retrieve the commit ID prior to Commit.
type DB struct {
	pool *Pool    // raw connection pool
	repl *replMon // logical replication monitor for collecting commit IDs

	// Guarantee that we are in-session by tracking and using a Tx for the write methods.
	mtx        sync.Mutex
	autoCommit bool   // skip the explicit transaction (begin/commit automatically)
	tx         pgx.Tx // interface
	txid       string // uid of the prepared transaction
}

// DBConfig is the configuration for the Kwil DB backend, which includes the
// connection parameters and a schema filter used to selectively include WAL
// data for certain PostgreSQL schemas in commit ID calculation.
type DBConfig struct {
	PoolConfig

	// SchemaFilter is used to include WAL data for certain *postgres* schema
	// (not Kwil schema). If nil, the default is to include updates to tables in
	// any schema prefixed by "ds_".
	SchemaFilter func(string) bool
}

const DefaultSchemaFilterPrefix = "ds_"

var defaultSchemaFilter = func(schema string) bool {
	return strings.HasPrefix(schema, DefaultSchemaFilterPrefix)
}

// [dev note] Transaction sequencing flow:
// - when ready to commit a tx, increment (UPDATE) the seq int8 in kwild_internal.sentry table
// - request from the repl monitor a promise for the commit ID for that seq
// - commit the tx
// - repl captures the ordered updates for the transaction
// - in repl receiver, decode and record the seq row update from WAL data (the final update before the commit message)
// - send complete commit digest back to the consumer via the promise channel for that seq value
// - ensure it matches the seq in the exec just prior
//
// To prepare for the above, initialize as follows:
// - create kwild_internal.sentry table if not exists
// - insert row with seq=0, if no rows

// NewDB creates a new Kwil DB instance. On creation, it will connect to the
// configured postgres process, creating as many connections as specified by the
// PoolConfig plus a special connection for a logical replication slot receiver.
// The database user (postgresql "role") must be a super user for several
// reasons: creating triggers, collations, and the replication publication.
func NewDB(ctx context.Context, cfg *DBConfig) (*DB, error) {
	// Create the unrestricted connection pool.
	pool, err := NewPool(ctx, &cfg.PoolConfig)
	if err != nil {
		return nil, err
	}

	// Ensure that the postgres host is running with an acceptable version.
	pgVer, pgVerNum, err := pgVersion(ctx, pool.writer)
	if err != nil {
		return nil, err
	}
	logger.Infof("Connected to %v", pgVer) // Connected to PostgreSQL 16.1 (Ubuntu 16.1-1.pgdg22.04+1) on ...

	major, minor, okVer := validateVersion(pgVerNum, verMajorRequired, verMinorRequired)
	if !okVer {
		return nil, fmt.Errorf("required PostgreSQL version not satisfied. Required %d.%d but connected to %d.%d",
			verMajorRequired, verMinorRequired, major, minor)
	}

	// Now check system settings, including logical replication and prepared transactions.
	if err = verifySettings(ctx, pool.writer); err != nil {
		return nil, err
	}

	// Verify that the db user/role is superuser with replication privileges.
	if err = checkSuperuser(ctx, pool.writer); err != nil {
		return nil, err
	}

	if err = setTimezoneUTC(ctx, pool.writer); err != nil {
		return nil, err
	}

	// Clean up orphaned prepared transaction that may have been left over from
	// an unclean shutdown. If we don't, postgres will hang on query.
	if _, err = rollbackPreparedTxns(ctx, pool.writer); err != nil {
		return nil, fmt.Errorf("failed to create publication: %w", err)
	}

	// Create the NOCASE collation to emulate SQLite's collation.
	if err = ensureCollation(ctx, pool.writer); err != nil {
		return nil, fmt.Errorf("failed to create custom collations: %w", err)
	}

	// Ensure all tables that are created with no primary key or unique index
	// are altered to have "full replication identity" for UPDATE and DELETES.
	if err = ensureTriggerReplIdentity(ctx, pool.writer); err != nil {
		return nil, fmt.Errorf("failed to create replication identity trigger: %w", err)
	}

	// Create the publication that is required for logical replication.
	if err = ensurePublication(ctx, pool.writer); err != nil {
		return nil, fmt.Errorf("failed to create publication: %w", err)
	}

	okSchema := cfg.SchemaFilter
	if okSchema == nil {
		okSchema = defaultSchemaFilter
	}

	repl, err := newReplMon(ctx, cfg.Host, cfg.Port, cfg.User, cfg.Pass, cfg.DBName, okSchema)
	if err != nil {
		return nil, err
	}

	// Create the tx sequence table with single row if it doesn't exists.
	if err = ensureSentryTable(ctx, pool.writer); err != nil {
		return nil, fmt.Errorf("failed to create transaction sequencing table: %w", err)
	}

	// Register the error function so a statement like `SELECT error('boom');`
	// will raise an exception and cause the query to error.
	if err = ensureErrorPLFunc(ctx, pool.writer); err != nil {
		return nil, fmt.Errorf("failed to create ERROR function: %w", err)
	}

	return &DB{
		pool: pool,
		repl: repl,
	}, nil
}

// Close shuts down the Kwil DB. This stops all connections and the WAL data
// receiver.
func (db *DB) Close() error {
	db.repl.stop()
	return db.pool.Close()
}

// AutoCommit toggles auto-commit mode, in which the Execute method may be used
// without having to begin/commit. This is to support startup and initialization
// tasks that occur prior to the start of the atomic commit process used while
// executing blocks.
func (db *DB) AutoCommit(auto bool) {
	db.mtx.Lock()
	defer db.mtx.Unlock()
	if db.tx != nil {
		panic("already in a tx")
	}
	db.autoCommit = auto
}

// For {accounts,validators}.Datasets / registry.DB
var _ sql.Executor = (*DB)(nil)

var _ sql.OuterTxMaker = (*DB)(nil) // for dataset Registry

// BeginTx makes the DB's singular transaction, which is used automatically by
// consumers of the Query and Execute methods. This is the mode of operation
// used by Kwil to have one system coordinating transaction lifetime, with one
// or more other systems implicitly using the transaction for their queries.
//
// The returned transaction is also capable of creating nested transactions.
// This functionality is used to prevent user dataset query errors from rolling
// back the outermost transaction.
func (db *DB) BeginTx(ctx context.Context) (sql.OuterTx, error) {
	tx, err := db.beginWriterTx(ctx)
	if err != nil {
		return nil, err
	}

	ntx := &nestedTx{
		Tx:         tx,
		accessMode: sql.ReadWrite,
	}
	return &dbTx{
		nestedTx:   ntx,
		db:         db,
		accessMode: sql.ReadWrite,
	}, nil
}

// ReadTx creates a read-only transaction for the database.
// It obtains a read connection from the pool, which will be returned
// to the pool when the transaction is closed.
func (db *DB) BeginReadTx(ctx context.Context) (sql.Tx, error) {
	conn, err := db.pool.pgxp.Acquire(ctx) // ensure we have a connection
	if err != nil {
		return nil, err
	}

	tx, err := conn.BeginTx(ctx, pgx.TxOptions{
		AccessMode: pgx.ReadOnly,
		IsoLevel:   pgx.RepeatableRead,
	})
	if err != nil {
		conn.Release()
		return nil, err
	}

	ntx := &nestedTx{
		Tx:         tx,
		accessMode: sql.ReadOnly,
	}

	return &readTx{
		nestedTx: ntx,
		release:  conn.Release,
	}, nil
}

var _ sql.TxBeginner = (*DB)(nil) // for CommittableStore => MultiCommitter

// Begin is for consumers that require a smaller interface on the return but
// same instance of the concrete type, a case which annoyingly creates
// incompatible interfaces in Go.
func (db *DB) Begin(ctx context.Context) (sql.TxCloser, error) {
	return db.BeginTx(ctx) // just slice down sql.Tx
}

// beginWriterTx is the critical section of BeginTx.
// It creates a new transaction on the write connection, and stores it in the
// DB's tx field. It is not exported, and is only called from BeginTx.
func (db *DB) beginWriterTx(ctx context.Context) (pgx.Tx, error) {
	db.mtx.Lock()
	defer db.mtx.Unlock()

	if db.tx != nil {
		return nil, errors.New("writer tx exists")
	}

	tx, err := db.pool.writer.BeginTx(ctx, pgx.TxOptions{
		AccessMode: pgx.ReadWrite,
		IsoLevel:   pgx.ReadUncommitted,
	})
	if err != nil {
		return nil, err
	}

	// Make the tx available to Execute and QueryPending.
	db.tx = tx

	return tx, nil
}

// precommit finalizes the transaction with a prepared transaction and returns
// the ID of the commit. The transaction is not yet committed.
func (db *DB) precommit(ctx context.Context) ([]byte, error) {
	db.mtx.Lock()
	defer db.mtx.Unlock()

	if db.tx == nil {
		return nil, errors.New("no tx exists")
	}

	// Do the seq update in sentry table. This ensures a replication message
	// sequence is emitted from this transaction, and that the data returned
	// from it includes the expected seq value.
	seq, err := incrementSeq(ctx, db.tx)
	if err != nil {
		return nil, err
	}
	logger.Debugf("updated seq to %d", seq)

	resChan := db.repl.recvID(seq)

	db.txid = random.String(10)
	sqlPrepareTx := fmt.Sprintf(`PREPARE TRANSACTION '%s'`, db.txid)
	if _, err = db.tx.Exec(ctx, sqlPrepareTx); err != nil {
		return nil, err
	}

	logger.Debugf("prepared transaction %q", db.txid)

	// Wait for the "commit id" from the replication monitor.
	select {
	case commitID := <-resChan:
		logger.Debugf("received commit ID %x", commitID)
		// The transaction is ready to commit, stored in a file with postgres in
		// the pg_twophase folder of the pg cluster data_directory.
		return commitID, nil
	case err = <-db.repl.errChan: // the replMon has died, so probably DB should close too...
		return nil, err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

// commit is called from the Commit method of the sql.Tx (or sql.TxCloser)
// returned from BeginTx (or Begin). See tx.go.
func (db *DB) commit(ctx context.Context) error {
	db.mtx.Lock()
	defer db.mtx.Unlock()

	if db.tx == nil {
		return errors.New("no tx exists")
	}
	if db.txid == "" { // NOTE: we could consider doing a regular commit if not using prepared, but for now we that flow
		return errors.New("transaction not yet prepared")
	}

	defer func() {
		if db.tx == nil {
			return
		}
		sqlRollback := fmt.Sprintf(`ROLLBACK PREPARED '%s'`, db.txid)
		db.txid = ""
		if _, err := db.tx.Exec(ctx, sqlRollback); err != nil {
			logger.Warnf("ROLLBACK PREPARED failed: %v", err)
		}
		db.tx = nil
	}()

	sqlCommit := fmt.Sprintf(`COMMIT PREPARED '%s'`, db.txid)
	if _, err := db.tx.Exec(ctx, sqlCommit); err != nil {
		return fmt.Errorf("COMMIT PREPARED failed: %v", err)
	}

	// Success, the defer should not try to rollback, and we should forget about
	// this prepared transaction's name, otherwise a future tx rollback prior to
	// prepare will try to rollback this old prepared txn.
	db.tx = nil
	db.txid = ""

	return nil
}

// rollback is called from the Rollback method of the sql.Tx (or sql.TxCloser)
// returned from BeginTx (or Begin). See tx.go.
func (db *DB) rollback(ctx context.Context) error {
	db.mtx.Lock()
	defer db.mtx.Unlock()

	if db.tx == nil {
		return errors.New("no tx exists")
	}

	defer func() {
		db.tx = nil
		db.txid = ""
	}()

	// If precommit not yet done, do a regular rollback.
	if db.txid == "" {
		db.tx.Rollback(ctx)
		return nil
	}

	// With precommit already done, rollback the prepared transaction, and do
	// not do the regular rollback, which is a no-op that emits a warning
	// notice: "WARNING:  there is no transaction in progress".
	sqlRollback := fmt.Sprintf(`ROLLBACK PREPARED '%s'`, db.txid)
	if _, err := db.tx.Exec(ctx, sqlRollback); err != nil {
		return fmt.Errorf("ROLLBACK PREPARED failed: %v", err)
	}

	return nil
}

// Query performs a read-only query on a read connection.
func (db *DB) Query(ctx context.Context, stmt string, args ...any) (*sql.ResultSet, error) {
	// Pass through to the read pool, isolated from any active transactions on
	// the write connection.
	return db.pool.Query(ctx, stmt, args...)
}

// discardCommitID is for Execute when in auto-commit mode.
func (db *DB) discardCommitID(ctx context.Context, resChan chan []byte) {
	select {
	case cid := <-resChan:
		logger.Debugf("discarding commit ID %x", cid)
	case <-db.repl.done:
	case <-ctx.Done():
	}
}

// Pool is a trapdoor to get the connection pool. Probably not for normal Kwil
// DB operation, but test setup/teardown.
func (db *DB) Pool() *Pool {
	return db.pool
}

// Execute runs a statement on an existing transaction, or on a short lived
// transaction from the write connection if in auto-commit mode.
func (db *DB) Execute(ctx context.Context, stmt string, args ...any) (*sql.ResultSet, error) {
	db.mtx.Lock()
	defer db.mtx.Unlock()

	if db.tx != nil {
		if db.autoCommit {
			return nil, errors.New("tx already created, cannot use auto commit")
		}
		return query(ctx, db.tx, stmt, args...)
	}
	if !db.autoCommit {
		return nil, sql.ErrNoTransaction
	}

	// We do manual autocommit since postgresql will skip it for some
	// statements, plus we are also injecting the seq update query.
	var resChan chan []byte
	var res *sql.ResultSet
	err := pgx.BeginTxFunc(ctx, db.pool.writer,
		pgx.TxOptions{
			AccessMode: pgx.ReadWrite,
			IsoLevel:   pgx.ReadCommitted,
		},
		func(tx pgx.Tx) error {
			seq, err := incrementSeq(ctx, tx)
			if err != nil {
				return err
			}
			resChan = db.repl.recvID(seq)
			res, err = query(ctx, tx, stmt, args...)
			return err
		},
	)
	if err != nil {
		return nil, err
	}
	db.discardCommitID(ctx, resChan)
	return res, nil
}

// TODO: require rw with target_session_attrs=read-write ?
