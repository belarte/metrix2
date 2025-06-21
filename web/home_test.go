package web

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHomePage(t *testing.T) {
	env := setupTestEnv(t, "http://localhost:8080")
	defer env.teardown()

	title, err := env.page.Title()
	require.NoError(t, err, "could not get title")
	require.Equal(t, "Metrix 2025", title, "expected title 'Metrix 2025'")

	content, err := env.page.Content()
	require.NoError(t, err, "could not get content")
	require.Contains(t, content, "Welcome to Metrix2! Track your metrics with ease.", "expected welcome message in content")
}
