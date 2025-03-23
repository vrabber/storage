package store

import (
	"github.com/vrabber/storage/internal/store/driver"
)

type Store interface {
	RegisterDriver(driver driver.Driver) error
	SetTemporary(tmp Temporary)
	Temporary() (Temporary, error)
}
