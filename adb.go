package adb

import (
	"errors"
	"os/exec"
	"path/filepath"
	"regexp"
)

// ErrADBNotFound is returned when the ADB executable is not found.
var ErrADBNotFound = errors.New("ADB command not found")

// ErrNoDevicesFound is returned when the ADB executable returns an empty device list.
var ErrNoDevicesFound = errors.New("no devices found")

// ErrParseDeviceList is returned when the ADB executable returns an invalid device list.
var ErrParseDeviceList = errors.New("could not parse device list")

// ErrParseDeviceState is returned when the ADB executable returns an invalid device state.
var ErrParseDeviceState = errors.New("could not parse device state")

// ADBPath contains the absolute path to the adb executable, empty string if the adb executable was not found.
// In last case it should be set manually.
var ADBPath string

// DeviceState is the state of an device.
type DeviceState int

const (
	//DeviceOffline referes to the state offline.
	DeviceOffline DeviceState = iota
	//DeviceOnline referes to the state online.
	DeviceOnline
	//DeviceUnauthorized referes to the state unauthorized.
	DeviceUnauthorized
	//DeviceUnknown refers to the state unknown.
	DeviceUnknown
)

// Device contains details about an attached device.
type Device struct {
	Serial string
	State  DeviceState
}

// RegexpSerial contains regular expression to match serial numbers.
var RegexpSerial = regexp.MustCompile(`^[0-9a-f]{8}$`)

func init() {
	// search in PATH variable
	if path, err := exec.LookPath("adb"); err == nil {
		if path, err = filepath.Abs(path); err == nil {
			ADBPath = path
		}
	}
}
