package models

import (
	"time"

	"github.com/google/uuid"
)

type FileUpload struct {
	ID        int
	UploadID  uuid.UUID
	FileInfo  *FileInfo
	CreatedAt time.Time
	UpdatedAt time.Time
}
