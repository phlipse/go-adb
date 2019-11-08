package adb_test

import (
	"strings"
	"testing"

	"github.com/phlipse/go-adb"
)

func TestExec(t *testing.T) {
	s, err := adb.Exec(5, "version")
	if err != nil {
		t.Errorf("error executing command: %v", err)
	}

	if !strings.HasPrefix(s, "Android Debug Bridge version") {
		t.Errorf("error executing command: stdout does not match expected output")
	}
}
