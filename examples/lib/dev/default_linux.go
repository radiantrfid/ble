package dev

import (
	"github.com/radiantrfid/ble"
	"github.com/radiantrfid/ble/linux"
)

// DefaultDevice ...
func DefaultDevice(opts ...ble.Option) (d ble.Device, err error) {
	return linux.NewDevice(opts...)
}
