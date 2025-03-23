package service

import (
	"context"
	"fmt"
	"math"
	"time"

	"github.com/google/uuid"
	pb "github.com/vrabber/storage/gen/storage"
	"github.com/vrabber/storage/internal/models"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) Upload(ctx context.Context, req *pb.UploadRequest, timeout time.Duration) error {
	startAt := time.Now()

	u, err := i.validateUploadRequest(req)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	return i.repo.Tx(ctx, func(ctx context.Context) error {
		upload, err := i.repo.GetFileUpload(ctx, req.GetUploadId())
		if err != nil {
			return err
		}

		if int64(req.GetOffset())+int64(len(req.GetContent())) > upload.FileInfo.Size {
			return ErrorUploadInvalidOffset
		}

		part := &models.FileUploadPart{
			FileUpload: &models.FileUpload{
				UploadID: u,
			},
			Size:      int64(len(req.Content)),
			Offset:    int64(req.Offset),
			StartedAt: startAt,
		}

		if err = i.repo.InsertFileUploadPart(ctx, part); err != nil {
			return err
		}

		tmp, err := i.store.Temporary()
		if err != nil {
			return err
		}

		if err = tmp.WriteData(upload.Name, req.GetContent(), int64(req.GetOffset())); err != nil {
			return err
		}

		now := time.Now()
		part.FinishedAt = &now
		return i.repo.UpdateFileUploadPartFinished(ctx, part)
	})
}

func (i *Implementation) OnDownloadFinished(ctx context.Context, uploadID string) error {
	//todo
	return nil
}

func (i *Implementation) validateUploadRequest(req *pb.UploadRequest) (u uuid.UUID, err error) {
	if req == nil {
		return uuid.UUID{}, status.Errorf(codes.InvalidArgument, "upload request requred")
	}
	if u, err = uuid.Parse(req.UploadId); err != nil {
		return u, ErrorUploadInvalidUploadID
	}
	if len(req.Content) == 0 || len(req.Content) > math.MaxInt64 {
		return u, fmt.Errorf("content length must be between 1 and %d", math.MaxInt64)
	}
	return u, nil
}
