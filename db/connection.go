package db

import (
	"context"
	"database/sql"

	"github.com/mfulki/ewallet-go-fiber/constant"
)

type DBConnection interface {
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
	PrepareContext(ctx context.Context, query string, args ...any) (*sql.Stmt, error)
}

type dbConnection struct {
	conn *sql.DB
}

func NewDBConnection(db *sql.DB) *dbConnection {
	return &dbConnection{
		conn: db,
	}
}

func (d *dbConnection) ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error) {
	tx, ok := ctx.Value(constant.TxContext).(*sql.Tx)
	if !ok {
		return d.conn.ExecContext(ctx, query, args...)
	}
	return tx.ExecContext(ctx, query, args...)
}

func (d *dbConnection) QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error) {
	tx, ok := ctx.Value(constant.TxContext).(*sql.Tx)
	if !ok {
		return d.conn.QueryContext(ctx, query, args...)
	}

	return tx.QueryContext(ctx, query, args...)
}

func (d *dbConnection) QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row {
	tx, ok := ctx.Value(constant.TxContext).(*sql.Tx)
	if !ok {
		return d.conn.QueryRowContext(ctx, query, args...)
	}

	return tx.QueryRowContext(ctx, query, args...)
}
func (d *dbConnection) PrepareContext(ctx context.Context, query string, args ...any) (*sql.Stmt, error) {
	tx, ok := ctx.Value(constant.TxContext).(*sql.Tx)
	if !ok {
		return d.conn.PrepareContext(ctx, query)
	}
	return tx.PrepareContext(ctx, query)

}
