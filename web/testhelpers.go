package web

import (
	"fmt"
	"io"
	"net/http"
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

	// Wait for server to be ready
	for i := 0; i < 20; i++ { // up to 2 seconds
		resp, err := http.Get("http://localhost:8080")
		if err == nil && resp.StatusCode < 500 {
			resp.Body.Close()
			fmt.Println("Server is ready (status:", resp.Status, ")")
			return cmd, nil
		}
		if resp != nil {
			body, _ := io.ReadAll(resp.Body)
			fmt.Printf("Server not ready (status: %v): %s\n", resp.Status, string(body))
			resp.Body.Close()
		}
		time.Sleep(100 * time.Millisecond)
	}
	_ = cmd.Process.Kill()
	return nil, errServerNotReady
}

var errServerNotReady = &ServerNotReadyError{}

type ServerNotReadyError struct{}

func (e *ServerNotReadyError) Error() string { return "server did not become ready in time" }

func stopServer(cmd *exec.Cmd) {
	if cmd != nil {
		_ = cmd.Process.Kill()
	}
}
