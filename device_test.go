package adb_test

import (
	"testing"

	"github.com/phlipse/go-adb"
)

func TestParseDevices(t *testing.T) {
	o := `foobar\n
	42


List of devices attached
	36b8b42e        offline
14a7f42b        device
	11d5a42a        unauthorized
	
	`

	d, err := adb.ParseDevices(o)
	if err != nil {
		t.Errorf("error parsing devices: %v", err)
	}

	if len(d) != 3 {
		t.Errorf("error parsing devices: number of parsed devices does not match expected number")
	}

	if d[1].Serial != "14a7f42b" {
		t.Errorf("error parsing devices: serial of second device does not match expected serial")
	}

	if d[2].State != adb.DeviceUnauthorized {
		t.Errorf("error parsing devices: state of third device does not match expected state")
	}
}
