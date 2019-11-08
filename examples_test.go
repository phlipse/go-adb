package adb_test

import (
	"fmt"
	"os/exec"
	"syscall"

	"github.com/phlipse/go-adb"
)

func ExampleExec() {
	// simple error checking
	s, err := adb.Exec(5, "version")
	if err != nil {
		fmt.Printf("error occurred: %v\n", err)
	} else {
		fmt.Println("no error occurred")
		_ = s
		// fmt.Println(s)
	}

	// extended error checking
	s, err = adb.Exec(5, "foobar")
	if err != nil {
		// check for exit code != 0
		if exiterr, ok := err.(*exec.ExitError); ok {
			// try to get exit code
			if status, ok := exiterr.Sys().(syscall.WaitStatus); ok {
				fmt.Printf("exit code not 0: %d\n", status.ExitStatus())
			}
		} else {
			fmt.Printf("error occurred: %v\n", err)
		}
	} else {
		fmt.Printf("exit code is 0\n")
		_ = s
		//fmt.Println(s)
	}

	// Output:
	// no error occurred
	// exit code not 0: 1
}

func ExampleParseDevices() {
	o := `List of devices attached
	36b8b42e        offline`

	d, err := adb.ParseDevices(o)
	if err != nil {
		fmt.Println("error")
	}

	fmt.Println(d[0].Serial)

	// Output:
	// 36b8b42e
}

func ExampleGetDevices() {
	// we can not simulate ADB device
	//d, err := adb.GetDevices()
	//fmt.Println(d[0].Serial)
}
