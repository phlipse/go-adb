package adb

import (
	"bytes"
	"fmt"
	"os/exec"
	"syscall"
	"time"
)

// Exec builds up an ADB command and executes it.
func Exec(t time.Duration, arg ...string) (string, error) {
	// check if path to ADB is set
	if ADBPath == "" {
		return "", ErrADBNotFound
	}

	// build up command
	cmd := exec.Command(ADBPath, arg...)
	// hide command windows when executing ADB commands
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}

	// redirect stdout to buffer
	var stdout bytes.Buffer
	cmd.Stdout = &stdout

	err := cmd.Start()
	if err != nil {
		return "", err
	}

	// wait until command is finished or timeout occurs
	wait := make(chan error, 1)
	go func() {
		err := cmd.Wait()
		wait <- err
	}()

	select {
	case err = <-wait:
		if err != nil {
			return "", err
		}
	case <-time.After(t * time.Second):
		return "", fmt.Errorf("timeout during ADB command execution")
	}

	return stdout.String(), err
}
