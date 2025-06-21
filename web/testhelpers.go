package web

import (
	"os/exec"
	"time"
)

func startServer() (*exec.Cmd, error) {
	cmd := exec.Command("go", "run", "main.go", "serve")
	cmd.Stdout = nil
	cmd.Stderr = nil
	cmd.Dir = ".."
	if err := cmd.Start(); err != nil {
		return nil, err
	}
	time.Sleep(2 * time.Second)
	return cmd, nil
}

func stopServer(cmd *exec.Cmd) {
	if cmd != nil {
		_ = cmd.Process.Kill()
	}
}
