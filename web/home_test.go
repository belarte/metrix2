package web

import (
	"testing"
	"strings"
)

func TestHomePage(t *testing.T) {
	env := setupTestEnv(t, "http://localhost:8080")
	defer env.teardown()

	title, err := env.page.Title()
	if err != nil {
		t.Fatalf("could not get title: %v", err)
	}
	if title != "Metrix 2025" {
		t.Errorf("expected title 'Metrix 2025', got '%s'", title)
	}

	content, err := env.page.Content()
	if err != nil {
		t.Fatalf("could not get content: %v", err)
	}
	if !strings.Contains(content, "Welcome to Metrix2! Track your metrics with ease.") {
		t.Errorf("expected welcome message in content")
	}
}
