package repository

import (
	"context"

	"github.com/vrabber/storage/internal/db"
	"github.com/vrabber/storage/internal/models"
)

type Repository interface {
	Transactor
	FileInfoRepository
	FileUploadPartRepository
	FileRepository
}

type transactionFunc func(ctx context.Context) error

type Transactor interface {
	Tx(ctx context.Context, fn transactionFunc) error
}

type FileInfoRepository interface {
	InsertFileInfo(ctx context.Context, fileInfo *models.FileInfo) error
	InsertFileUpload(ctx context.Context, fileUpload *models.FileUpload) error
}

type FileUploadPartRepository interface {
	GetFileUpload(ctx context.Context, uploadID string) (*models.FileUpload, error)
	InsertFileUploadPart(ctx context.Context, fileUpload *models.FileUploadPart) error
	UpdateFileUploadPartFinished(ctx context.Context, part *models.FileUploadPart) error
	IsFullLoaded(ctx context.Context, fileUpload *models.FileUpload) (bool, error)
}

type FileRepository interface {
}

type Implementation struct {
	db db.Transactor
}

func NewRepositoryImplementation(db db.Transactor) Repository {
	return &Implementation{db: db}
}
