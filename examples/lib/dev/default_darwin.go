package dev

import (
	"github.com/radiantrfid/ble"
	"github.com/radiantrfid/ble/darwin"
)

// DefaultDevice ...
func DefaultDevice(opts ...ble.Option) (d ble.Device, err error) {
	return darwin.NewDevice(opts...)
}
