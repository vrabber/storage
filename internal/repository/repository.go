package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/vrabber/storage/internal/models"
)

type Repository interface {
	InsertFileInfo(ctx context.Context, fileInfo *models.FileInfo) error
	InsertFileUpload(ctx context.Context, fileUpload *models.FileUpload) error
}

type Implementation struct {
	pool *pgxpool.Pool
}

func NewRepositoryImplementation(pool *pgxpool.Pool) Repository {
	return &Implementation{pool: pool}
}
