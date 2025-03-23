package store

import (
	"context"
)

type Temporary interface {
	Reserve(ctx context.Context, name string, size uint64) error
	Release(ctx context.Context, name string)
	WriteData(name string, data []byte, offset int64) error
	Path() string
}
