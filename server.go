package adb

// StartServer ensures that there is a server running.
func StartServer() error {
	_, err := Exec(10, "start-server")
	return err
}

// StopServer stops the server if it is running.
func StopServer() error {
	_, err := Exec(5, "kill-server")
	return err
}

// RestartServer restarts the server.
func RestartServer() error {
	err := StopServer()
	if err != nil {
		return err
	}
	return StartServer()
}
