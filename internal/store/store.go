package store

import (
	"sync"

	"github.com/vrabber/storage/internal/store/driver"
)

type Store interface {
	RegisterDriver(driver driver.Driver) error
}

type Implementation struct {
	drivers map[string]driver.Driver
	lock    sync.RWMutex
}

func NewImplementation() *Implementation {
	return &Implementation{
		drivers: make(map[string]driver.Driver),
		lock:    sync.RWMutex{},
	}
}
