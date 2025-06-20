package web

import (
	"os/exec"
	"testing"
	"time"

	"github.com/playwright-community/playwright-go"
)

func startServer() (*exec.Cmd, error) {
	cmd := exec.Command("go", "run", "main.go", "serve")
	cmd.Stdout = nil
	cmd.Stderr = nil
	cmd.Dir = ".." // Run from project root so main.go is found
	if err := cmd.Start(); err != nil {
		return nil, err
	}
	// Wait for the server to start
	time.Sleep(2 * time.Second)
	return cmd, nil
}

func stopServer(cmd *exec.Cmd) {
	if cmd != nil {
		_ = cmd.Process.Kill()
	}
}

func TestHomePage(t *testing.T) {
	cmd, err := startServer()
	if err != nil {
		t.Fatalf("failed to start server: %v", err)
	}
	defer stopServer(cmd)

	pw, err := playwright.Run()
	if err != nil {
		t.Fatalf("could not launch playwright: %v", err)
	}
	defer pw.Stop()

	browser, err := pw.Chromium.Launch()
	if err != nil {
		t.Fatalf("could not launch browser: %v", err)
	}
	page, err := browser.NewPage()
	if err != nil {
		t.Fatalf("could not create page: %v", err)
	}

	_, err = page.Goto("http://localhost:8080")
	if err != nil {
		t.Fatalf("could not goto: %v", err)
	}

	title, err := page.Title()
	if err != nil {
		t.Fatalf("could not get title: %v", err)
	}
	if title != "Metrix2" {
		t.Errorf("expected title 'Metrix2', got '%s'", title)
	}

	content, err := page.Content()
	if err != nil {
		t.Fatalf("could not get content: %v", err)
	}
	if !contains(content, "Welcome to Metrix2! Track your metrics with ease.") {
		t.Errorf("expected welcome message in content")
	}
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(s) > len(substr) && (contains(s[1:], substr) || contains(s[:len(s)-1], substr)))
}
