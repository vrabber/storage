package models

import (
	"fmt"
	"math"
	"time"

	pb "github.com/vrabber/storage/gen/storage"
)

type FileInfo struct {
	ID        int
	Name      string
	Size      int64
	Hash      string
	Owner     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (f *FileInfo) FromRequest(fileInfo *pb.FileInfo) error {
	if fileInfo.Size > math.MaxInt64 {
		return fmt.Errorf("file size must smaller than %d", math.MaxInt64)
	}

	f.Name = fileInfo.Filename
	f.Size = int64(fileInfo.Size)
	f.Hash = fileInfo.Hash

	if fileInfo.Metadata != nil {
		f.Owner = fileInfo.Metadata.Owner
		if fileInfo.Metadata.CreatedAt != nil {
			f.CreatedAt = fileInfo.Metadata.CreatedAt.AsTime()
		}
		if fileInfo.Metadata.UpdatedAt != nil {
			f.UpdatedAt = fileInfo.Metadata.UpdatedAt.AsTime()
		}
	}
	return nil
}
