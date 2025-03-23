package models

import "time"

type File struct {
	ID       int
	Name     string
	Path     string
	FileInfo *FileInfo
	Created  time.Time
	Updated  time.Time
}
