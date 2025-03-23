package store

import (
	"fmt"
	"sync"

	"github.com/vrabber/storage/internal/store/driver"
)

type Implementation struct {
	drivers   map[string]driver.Driver
	lock      sync.RWMutex
	temporary driver.Driver
}

func NewImplementation() *Implementation {
	return &Implementation{
		drivers: make(map[string]driver.Driver),
		lock:    sync.RWMutex{},
	}
}

func (i *Implementation) SetTemporary(driver driver.Driver) error {
	if !driver.SupportsSeek() {
		return fmt.Errorf("driver %s does not support seeking", driver.Name())
	}
	i.temporary = driver
	return nil
}
