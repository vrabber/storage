package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Transactor interface {
	WithCtx(ctx context.Context, fn func(ctx context.Context) error) error
	Exec(ctx context.Context, sql string, args ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
}

type TransactionKey struct{}

type TransactorImpl struct {
	pool *pgxpool.Pool
}

func NewTransactorImpl(pool *pgxpool.Pool) Transactor {
	return &TransactorImpl{pool: pool}
}

func extractTransaction(ctx context.Context) (pgx.Tx, bool) {
	tx, ok := ctx.Value(TransactionKey{}).(pgx.Tx)
	return tx, ok
}

func (d *TransactorImpl) WithCtx(ctx context.Context, fn func(ctx context.Context) error) error {
	tx, err := d.pool.Begin(ctx)
	if err != nil {
		return err
	}

	txCtx := context.WithValue(ctx, TransactionKey{}, tx)

	if err = fn(txCtx); err != nil {
		if rbErr := tx.Rollback(ctx); rbErr != nil {
			return fmt.Errorf("rollback failed: %v, original err: %v", rbErr, err)
		}
		return err
	}

	return tx.Commit(ctx)
}

func (d *TransactorImpl) Exec(ctx context.Context, sql string, args ...interface{}) (pgconn.CommandTag, error) {
	tx, ok := extractTransaction(ctx)
	if ok {
		return tx.Exec(ctx, sql, args...)
	}
	return d.pool.Exec(ctx, sql, args...)
}

func (d *TransactorImpl) Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error) {
	tx, ok := extractTransaction(ctx)
	if ok {
		return tx.Query(ctx, sql, args...)
	}
	return d.pool.Query(ctx, sql, args...)
}

func (d *TransactorImpl) QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row {
	tx, ok := extractTransaction(ctx)
	if ok {
		return tx.QueryRow(ctx, sql, args...)
	}
	return d.pool.QueryRow(ctx, sql, args...)
}
