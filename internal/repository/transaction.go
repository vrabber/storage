package repository

import (
	"context"
	"log/slog"

	"github.com/jackc/pgx/v5"
)

type transactionFunc func(ctx context.Context, tx pgx.Tx) error

func (i *Implementation) tx(ctx context.Context, fn transactionFunc) error {
	tx, err := i.pool.Begin(ctx)
	if err != nil {
		return err
	}

	if err = fn(ctx, tx); err != nil {
		if err := tx.Rollback(ctx); err != nil {
			slog.Error("failed to roll back transaction ", "err", err)
		}
		return err
	}
	return tx.Commit(ctx)
}
