package store

import (
	"fmt"
	"sync"

	"github.com/vrabber/storage/internal/store/driver"
)

type Implementation struct {
	drivers   map[string]driver.Driver
	lock      sync.RWMutex
	temporary Temporary
}

func NewImplementation() *Implementation {
	return &Implementation{
		drivers: make(map[string]driver.Driver),
		lock:    sync.RWMutex{},
	}
}

func (i *Implementation) SetTemporary(tmp Temporary) {
	i.temporary = tmp
}

func (i *Implementation) Temporary() (Temporary, error) {
	if i.temporary == nil {
		return nil, fmt.Errorf("temporary driver not initialized")
	}
	return i.temporary, nil
}
