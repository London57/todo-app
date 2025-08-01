package postgres

import (
	"context"

	mock_postgres "github.com/London57/todo-app/pkg/postgres/mocks"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

//go:generate mockgen -source=pgx_adapter.go -destination=mocks/pool_mock.go
type Pool interface {
	Exec(ctx context.Context, sql string, arguments ...any) (commandTag pgconn.CommandTag, err error)
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	Begin(ctx context.Context) (pgx.Tx, error)
	BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error)
	Close()
}


type PgxTx interface {
	Begin(ctx context.Context) (pgx.Tx, error)
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
	Prepare(ctx context.Context, name, sql string) (*pgconn.StatementDescription, error)
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	Conn() *pgx.Conn
	CopyFrom(ctx context.Context, tableName pgx.Identifier, columnNames []string, rowSrc pgx.CopyFromSource) (int64, error)
	SendBatch(ctx context.Context, b *pgx.Batch) pgx.BatchResults
	LargeObjects() pgx.LargeObjects
}

type PgxRow interface {
	Scan(dest ...any) error
}

type PgxTxAdapter struct {
	tx *mock_postgres.MockPgxTx
}

func (a *PgxTxAdapter) GetTx() *mock_postgres.MockPgxTx {
	return a.tx
}

func NewPgxTxAdapter(tx *mock_postgres.MockPgxTx) *PgxTxAdapter {
	return &PgxTxAdapter{tx}
}
func (a *PgxTxAdapter) CopyFrom(
	ctx context.Context,
	tableName pgx.Identifier,
	columnNames []string,
	rowSrc pgx.CopyFromSource,
) (int64, error) {
	return a.tx.CopyFrom(ctx, tableName, columnNames, rowSrc)
}

// SendBatch делегирует вызов к оригинальному *pgx.Tx
func (a *PgxTxAdapter) SendBatch(ctx context.Context, b *pgx.Batch) pgx.BatchResults {
	return a.tx.SendBatch(ctx, b)
}

// LargeObjects делегирует вызов к оригинальному *pgx.Tx
func (a *PgxTxAdapter) LargeObjects() pgx.LargeObjects {
	return a.tx.LargeObjects()
}

func (a *PgxTxAdapter) Begin(ctx context.Context) (pgx.Tx, error) {
	return a.tx.Begin(ctx)
}

func (a *PgxTxAdapter) Commit(ctx context.Context) error {
	return a.tx.Commit(ctx)
}

func (a *PgxTxAdapter) Rollback(ctx context.Context) error {
	return a.tx.Rollback(ctx)
}

func (a *PgxTxAdapter) Prepare(ctx context.Context, name, sql string) (*pgconn.StatementDescription, error) {
	return a.tx.Prepare(ctx, name, sql)
}

func (a *PgxTxAdapter) Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error) {
	return a.tx.Exec(ctx, sql, arguments...)
}

func (a *PgxTxAdapter) Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error) {
	return a.tx.Query(ctx, sql, args...)
}

func (a *PgxTxAdapter) QueryRow(ctx context.Context, sql string, args ...any) pgx.Row {
	return a.tx.QueryRow(ctx, sql, args...)
}

func (a *PgxTxAdapter) Conn() *pgx.Conn {
	return a.tx.Conn()
}


type PgxPoolAdapter struct {
	pool *pgxpool.Pool
}
func (a *PgxPoolAdapter) Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error) {
	return a.pool.Exec(ctx, sql, arguments...)
}

func (a *PgxPoolAdapter) Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error) {
	return a.pool.Query(ctx, sql, args...)
}

func (a *PgxPoolAdapter) QueryRow(ctx context.Context, sql string, args ...any) pgx.Row {
	return a.pool.QueryRow(ctx, sql, args...)
}

func (a *PgxPoolAdapter) Begin(ctx context.Context) (pgx.Tx, error) {
	return a.pool.Begin(ctx)
}

// func (a *PgxPoolAdapter) BeginTx(ctx context.Context, txOptions pgx.TxOptions) (PgxTx, error) {
// 	tx, err := a.pool.BeginTx(ctx, txOptions)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &PgxTxAdapter{tx: tx}, nil
// }