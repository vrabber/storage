package store

import (
	"context"
)

type Temporary interface {
	Reserve(ctx context.Context, name string, size uint64) error
	Release(ctx context.Context, name string)
}
