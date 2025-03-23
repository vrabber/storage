package service

import (
	"context"
	"errors"

	pb "github.com/vrabber/storage/gen/storage"
	"github.com/vrabber/storage/internal/repository"
	"github.com/vrabber/storage/internal/store"
)

var (
	ErrorInitUploadInternal   = errors.New("internal error")
	ErrorInitUploadFileExists = errors.New("file exists")
)

type Service interface {
	InitUpload(ctx context.Context, req *pb.FileInfo) (string, error)
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
