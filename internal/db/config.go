package db

import (
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	MaxConns        int32 = 10
	MinConns        int32 = 2
	MaxConnIdleTime       = time.Second * 60
	MaxConnLifetime       = time.Hour
)

func configure(config *pgxpool.Config) {
	config.MaxConns = MaxConns
	config.MinConns = MinConns
	config.MaxConnIdleTime = MaxConnIdleTime
	config.MaxConnLifetime = MaxConnLifetime
}
