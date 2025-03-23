package store

import (
	"errors"

	"github.com/vrabber/storage/internal/store/driver"
)

var (
	errDriverExists = errors.New("driver already exists")
)

func (i *Implementation) RegisterDriver(driver driver.Driver) error {
	// Optimistic lock
	if i.driverExists(driver, true) {
		return errDriverExists
	}

	return i.registerDriver(driver)
}

func (i *Implementation) registerDriver(driver driver.Driver) error {
	i.lock.Lock()
	defer i.lock.Unlock()

	if i.driverExists(driver, false) {
		return errDriverExists
	}
	i.drivers[driver.Name()] = driver
	return nil
}

func (i *Implementation) driverExists(driver driver.Driver, withLock bool) bool {
	if withLock {
		i.lock.RLock()
		defer i.lock.RUnlock()
	}

	_, ok := i.drivers[driver.Name()]
	return ok
}
