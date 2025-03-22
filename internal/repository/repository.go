package repository

import "github.com/jackc/pgx/v5/pgxpool"

type Repository interface {
}

type Implementation struct {
	pool *pgxpool.Pool
}

func NewRepositoryImplementation(pool *pgxpool.Pool) Repository {
	return &Implementation{pool: pool}
}
