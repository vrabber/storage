package models

import "time"

type FileInfo struct {
	ID      int
	Name    string
	Size    int64
	Hash    string
	Owner   string
	Created time.Time
	Updated time.Time
}
