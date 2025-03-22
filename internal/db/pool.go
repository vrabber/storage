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
		return nil, err
	}

	configure(conf)

	pool, err := pgxpool.NewWithConfig(ctx, conf)
	if err != nil {
		return nil, err
	}

	return pool, ping(ctx, pool)
}

func createConnString(cnf config.DatabaseConfig) string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=true",
		cnf.User,
		cnf.Password,
		cnf.Host,
		cnf.Port,
		cnf.Name,
	)
}
