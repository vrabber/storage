package service

import (
	"context"
	"errors"
	"time"

	pb "github.com/vrabber/storage/gen/storage"
	"github.com/vrabber/storage/internal/repository"
	"github.com/vrabber/storage/internal/store"
)

var (
	ErrorInitUploadInternal   = errors.New("internal error")
	ErrorInitUploadFileExists = errors.New("file exists")

	ErrorUploadInvalidUploadID = errors.New("invalid upload id")
	ErrorUploadTimeout         = errors.New("upload timeout")
	ErrorUploadInvalidOffset   = errors.New("invalid offset")
)

type Service interface {
	InitUpload(ctx context.Context, req *pb.FileInfo) (string, error)
	Upload(ctx context.Context, req *pb.UploadRequest, timeout time.Duration) error
	OnDownloadFinished(ctx context.Context, uploadID string) error
}

type Implementation struct {
	repo  repository.Repository
	store store.Store
}

func NewService(repo repository.Repository, store store.Store) Service {
	return &Implementation{
		repo:  repo,
		store: store,
	}
}
