package models

import "time"

type FileUploadPart struct {
	ID         int
	FileUpload *FileUpload
	Size       int64
	Offset     int64
	StartedAt  time.Time
	FinishedAt *time.Time
}
