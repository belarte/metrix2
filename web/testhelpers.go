package web

import (
	"fmt"
	"io"
	"net/http"
	"os/exec"
	"testing"
	"time"

	"github.com/playwright-community/playwright-go"
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

type TestEnv struct {
	cmd     *exec.Cmd
	pw      *playwright.Playwright
	browser playwright.Browser
	page    playwright.Page
}

func setupTestEnv(t *testing.T, url string) *TestEnv {
	cmd, err := startServer()
	if err != nil {
		t.Fatalf("failed to start server: %v", err)
	}

	pw, err := playwright.Run()
	if err != nil {
		stopServer(cmd)
		t.Fatalf("could not launch playwright: %v", err)
	}

	browser, err := pw.Chromium.Launch()
	if err != nil {
		pw.Stop()
		stopServer(cmd)
		t.Fatalf("could not launch browser: %v", err)
	}

	page, err := browser.NewPage()
	if err != nil {
		browser.Close()
		pw.Stop()
		stopServer(cmd)
		t.Fatalf("could not create page: %v", err)
	}

	page.SetDefaultTimeout(3000)
	_, err = page.Goto(url)
	if err != nil {
		page.Close()
		browser.Close()
		pw.Stop()
		stopServer(cmd)
		t.Fatalf("could not goto: %v", err)
	}

	return &TestEnv{cmd, pw, browser, page}
}

func (env *TestEnv) teardown() {
	if env.page != nil {
		env.page.Close()
	}
	if env.browser != nil {
		env.browser.Close()
	}
	if env.pw != nil {
		env.pw.Stop()
	}
	stopServer(env.cmd)
}
