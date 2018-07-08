package ibbq

import (
	"github.com/go-ble/ble"
	"github.com/sworisbreathing/go-ibbq/ibbq/os"
)

// NewDevice creates a new device
func NewDevice(impl string) (d ble.Device, err error) {
	return os.DefaultDevice()
}
