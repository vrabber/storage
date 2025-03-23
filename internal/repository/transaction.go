package repository

import (
	"context"
)

func (i *Implementation) Tx(ctx context.Context, fn transactionFunc) error {
	return i.db.WithCtx(ctx, fn)
}
