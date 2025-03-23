package models

import (
	"time"

	"github.com/google/uuid"
)

type FileUpload struct {
	ID        int
	UploadID  uuid.UUID
	FileInfo  *FileInfo
	Name      string
	Path      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
