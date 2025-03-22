package db

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var PingTimeout = 5 * time.Second

func ping(ctx context.Context, pool *pgxpool.Pool) error {
	ctx, cancel := context.WithTimeout(ctx, PingTimeout)
	defer cancel()

	return pool.Ping(ctx)
}
