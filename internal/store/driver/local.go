package driver

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"math"
	"os"
	"path/filepath"
)

type LocalDriver struct {
	basePath string
}

func NewLocalDriver(basePath string) Driver {
	return &LocalDriver{basePath: basePath}
}

func (l *LocalDriver) Name() string {
	return Local
}

func (l *LocalDriver) SupportsSeek() bool {
	return true
}

func (l *LocalDriver) Path() string {
	return l.basePath
}

func (l *LocalDriver) Reserve(_ context.Context, name string, size uint64) error {
	if size > math.MaxInt64 {
		return fmt.Errorf("size must be less than or equal to %d", math.MaxInt64)
	}

	filePath := filepath.Join(l.basePath, name)

	if err := os.MkdirAll(filepath.Dir(filePath), 0755); err != nil {
		return err
	}

	f, err := os.Create(filePath)
	if err != nil {
		if errors.Is(err, os.ErrExist) {
			return ErrorFileExists
		}
		return err
	}
	defer func(f *os.File) {
		if err := f.Close(); err != nil {
			slog.Error("failed to close file", "file", name, "err", err)
		}
	}(f)

	if err = f.Truncate(int64(size)); err != nil {
		if err := os.Remove(filePath); err != nil {
			slog.Error("failed to remove file", "file", name, "err", err)
		}
		return err
	}

	return nil
}

func (l *LocalDriver) Release(_ context.Context, name string) {
	if err := os.Remove(filepath.Join(l.basePath, name)); err != nil {
		slog.Error("failed to remove file", "file", name, "err", err)
	}
}

func (l *LocalDriver) WriteData(name string, data []byte, offset int64) error {
	f, err := os.OpenFile(filepath.Join(l.basePath, name), os.O_WRONLY, 0755)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return ErrorFileExists
		}
		return err
	}

	defer func(f *os.File) {
		if err := f.Close(); err != nil {
			slog.Error("failed to close file", "file", name, "err", err)
		}
	}(f)

	stat, err := f.Stat()
	if err != nil {
		return err
	}

	if int64(len(data))+offset > stat.Size() {
		return ErrorInvalidOffset
	}
	if _, err := f.Seek(offset, io.SeekStart); err != nil {
		return err
	}

	if _, err := f.Write(data); err != nil {
		return err
	}
	return nil
}
