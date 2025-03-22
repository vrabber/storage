package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/vrabber/storage/internal/config"
)

func CreatePool(ctx context.Context, cnf config.DatabaseConfig) (*pgxpool.Pool, error) {
	conf, err := pgxpool.ParseConfig(createConnString(cnf))
	if err != nil {
		return nil, fmt.Errorf("error parsing db config: %w", err)
	}

	configure(conf)

	pool, err := pgxpool.NewWithConfig(ctx, conf)
	if err != nil {
		return nil, fmt.Errorf("error creating pool: %w", err)
	}

	return pool, ping(ctx, pool)
}

func createConnString(cnf config.DatabaseConfig) string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		cnf.User,
		cnf.Password,
		cnf.Host,
		cnf.Port,
		cnf.Name,
	)
}
