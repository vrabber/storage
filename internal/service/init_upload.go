package service

import (
	"context"
	"errors"
	"log/slog"

	"github.com/google/uuid"
	pb "github.com/vrabber/storage/gen/storage"
	"github.com/vrabber/storage/internal/models"
	"github.com/vrabber/storage/internal/store/driver"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) InitUpload(ctx context.Context, req *pb.FileInfo) (string, error) {
	if err := i.validateInitUploadFileInfo(req); err != nil {
		return "", err
	}

	fileName := uuid.NewString()
	path, err := i.reserveFile(ctx, fileName, req.Size)
	if err != nil {
		return "", err
	}

	fileInfo, err := i.storeFileInfo(ctx, req)
	if err != nil {
		i.releaseFile(ctx, fileName)
		return "", err
	}

	uploadID := uuid.New()

	if err = i.storeFileUpload(ctx, fileName, path, fileInfo, uploadID); err != nil {
		return "", err
	}

	return uploadID.String(), nil
}

func (i *Implementation) reserveFile(ctx context.Context, name string, size uint64) (string, error) {
	temp, err := i.store.Temporary()
	if err != nil {
		slog.Error("failed to get temporary storage", "err", err)
		return "", ErrorInitUploadInternal
	}

	if err = temp.Reserve(ctx, name, size); err != nil {
		slog.Error("failed to reserve file", "err", err)

		if errors.Is(err, driver.ErrorFileExists) {
			return "", ErrorInitUploadFileExists
		}
		return "", err
	}
	return temp.Path(), nil
}

func (i *Implementation) releaseFile(ctx context.Context, name string) {
	temp, err := i.store.Temporary()
	if err != nil {
		slog.Error("failed to get temporary storage", "err", err)
		return
	}
	temp.Release(ctx, name)
}

func (i *Implementation) storeFileInfo(ctx context.Context, req *pb.FileInfo) (models.FileInfo, error) {
	result := models.FileInfo{}
	if err := result.FromRequest(req); err != nil {
		slog.Error("failed to parse file info", "err", err)
		return result, ErrorInitUploadInternal
	}

	if err := i.repo.InsertFileInfo(ctx, &result); err != nil {
		slog.Error("failed to store fileInfo", "err", err)
		return result, ErrorInitUploadInternal
	}
	return result, nil
}

func (i *Implementation) storeFileUpload(ctx context.Context, name string, path string, fileInfo models.FileInfo, uploadID uuid.UUID) error {
	upload := models.FileUpload{
		UploadID: uploadID,
		Name:     name,
		Path:     path,
		FileInfo: &fileInfo,
	}

	if err := i.repo.InsertFileUpload(ctx, &upload); err != nil {
		slog.Error("failed to store fileUpload", "err", err)
		return ErrorInitUploadInternal
	}
	return nil
}

func (i *Implementation) validateInitUploadFileInfo(info *pb.FileInfo) error {
	if info == nil {
		return status.Errorf(codes.InvalidArgument, "file info requred")
	}
	if info.Filename == "" {
		return status.Errorf(codes.InvalidArgument, "filename required")
	}
	if info.Size == 0 {
		return status.Errorf(codes.InvalidArgument, "size must be greater than zero")
	}
	if info.Hash == "" {
		return status.Errorf(codes.InvalidArgument, "hash required")
	}
	return nil
}
