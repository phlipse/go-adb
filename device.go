package adb

import (
	"fmt"
	"strings"
)

// GetDevices is a helper function that returns a list of attached devices.
func GetDevices() ([]*Device, error) {
	s, err := Exec(5, "devices")
	if err != nil {
		return nil, err
	}

	return ParseDevices(s)
}

// GetDeviceState returns state of specific device.
func GetDeviceState(serial string) (DeviceState, error) {
	if !RegexpSerial.MatchString(serial) {
		// boundary only needed to format error
		boundary := len(serial)
		if boundary > 10 {
			boundary = 10
		}
		return DeviceUnknown, fmt.Errorf("serial has wrong format: %s", serial[:boundary])
	}

	devices, err := GetDevices()
	if err != nil {
		return DeviceUnknown, err
	}

	for _, d := range devices {
		if d.Serial == serial {
			return d.State, nil
		}
	}

	// we found no device with given serial
	return DeviceUnknown, nil
}

// ParseDevices parses the output from ADB devices command to Device struct.
func ParseDevices(s string) ([]*Device, error) {
	a := strings.SplitAfter(s, "List of devices attached")
	if len(a) != 2 {
		return nil, ErrNoDevicesFound
	}

	lines := strings.Split(a[1], "\n")
	devices := make([]*Device, 0, len(lines))
	for _, line := range lines {
		fields := strings.Fields(line)
		switch len(fields) {
		case 0:
			continue
		case 2:
			if !RegexpSerial.MatchString(fields[0]) {
				continue
			}

			var state DeviceState
			switch strings.ToLower(strings.TrimSpace(fields[1])) {
			case "offline":
				state = DeviceOffline
			case "device":
				state = DeviceOnline
			case "unauthorized":
				state = DeviceUnauthorized
			case "unknown":
				state = DeviceUnknown
			default:
				return nil, ErrParseDeviceState
			}

			device := &Device{
				Serial: fields[0],
				State:  state,
			}
			devices = append(devices, device)
		default:
			return nil, ErrParseDeviceList
		}
	}

	return devices, nil
}
