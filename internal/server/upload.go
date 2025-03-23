package server

import (
	"errors"
	"io"
	"log/slog"

	pb "github.com/vrabber/storage/gen/storage"
	"github.com/vrabber/storage/internal/service"
)

func (s *Server) Upload(stream pb.StorageService_UploadServer) error {
	var uploadID string

	for {
		req, err := stream.Recv()
		if err != nil {
			return s.onUploadError(stream, uploadID, err)
		}

		if req.UploadId == "" || (uploadID != "" && uploadID != req.UploadId) {
			return s.onUploadError(stream, uploadID, service.ErrorUploadInvalidUploadID)
		}
		uploadID = req.UploadId

		if err = s.srv.Upload(s.ctx, req, s.conf.PartUploadTimeout); err != nil {
			return s.onUploadError(stream, uploadID, err)
		}
	}
}

func (s *Server) onUploadError(stream pb.StorageService_UploadServer, uploadID string, err error) error {
	if err == io.EOF {
		if fullUploadErr := s.srv.OnDownloadFinished(s.ctx, uploadID); fullUploadErr != nil {
			return s.sendUploadResponse(stream, fullUploadErr)
		}
		return s.sendUploadResponse(stream, nil)
	}

	slog.Error("upload error", "err", err)
	return s.sendUploadResponse(stream, err)
}

func (s *Server) sendUploadResponse(stream pb.StorageService_UploadServer, err error) error {
	result := &pb.UploadResponse{}

	switch {
	case err == nil:
		result.Status = pb.UploadStatus_UPLOAD_STATUS_SUCCESS
		return stream.SendAndClose(result)
	case errors.Is(err, service.ErrorUploadInvalidUploadID):
		result.Status = pb.UploadStatus_UPLOAD_STATUS_ERROR_INVALID_UPLOAD_ID
	case errors.Is(err, service.ErrorUploadTimeout):
		result.Status = pb.UploadStatus_UPLOAD_STATUS_ERROR_UPLOAD_TIMEOUT
	case errors.Is(err, service.ErrorUploadInvalidOffset):
		result.Status = pb.UploadStatus_UPLOAD_STATUS_ERROR_INVALID_OFFSET
	}

	result.Error = err.Error()
	return stream.SendAndClose(result)
}
