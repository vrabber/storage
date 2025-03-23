package server

import (
	"context"
	"errors"

	pb "github.com/vrabber/storage/gen/storage"
	"github.com/vrabber/storage/internal/service"
)

func (s *Server) InitUpload(ctx context.Context, req *pb.InitUploadRequest) (*pb.InitUploadResponse, error) {
	result := &pb.InitUploadResponse{
		Status: 0,
		Result: nil,
	}

	uploadID, err := s.srv.InitUpload(ctx, req.Info)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrorInitUploadFileExists):
			result.Status = pb.InitUploadStatus_INIT_UPLOAD_STATUS_ERROR_FILE_EXISTS
		case errors.Is(err, service.ErrorInitUploadInternal):
			result.Status = pb.InitUploadStatus_INIT_UPLOAD_STATUS_ANOTHER_ERROR
		default:
			result.Status = pb.InitUploadStatus_INIT_UPLOAD_STATUS_ANOTHER_ERROR
		}

		result.Result = &pb.InitUploadResponse_Error{Error: err.Error()}
		return result, nil
	}

	result.Status = pb.InitUploadStatus_INIT_UPLOAD_STATUS_READY
	result.Result = &pb.InitUploadResponse_UploadId{UploadId: uploadID}
	return result, nil
}
