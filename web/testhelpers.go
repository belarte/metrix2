package web

import (
	"context"
	"net"
	"net/http"
	"testing"
	"time"

	"github.com/belarte/metrix2/router"
	"github.com/playwright-community/playwright-go"
	"github.com/stretchr/testify/require"
)

type TestEnv struct {
	server   *http.Server
	listener net.Listener
	pw       *playwright.Playwright
	browser  playwright.Browser
	page     playwright.Page
	url      string
}

func setupTestEnv(t *testing.T, endpoint string) *TestEnv {
	ln, err := net.Listen("tcp", ":0")
	require.NoError(t, err)

	srv := &http.Server{Handler: router.New()}
	go srv.Serve(ln)

	addr := ln.Addr().String()
	url := "http://" + addr + endpoint
	for i := 0; i < 20; i++ {
		resp, err := http.Get(url)
		if err == nil && resp.StatusCode < 500 {
			resp.Body.Close()
			break
		}
		if resp != nil {
			resp.Body.Close()
		}
		time.Sleep(100 * time.Millisecond)
	}

	pw, err := playwright.Run()
	require.NoError(t, err)

	browser, err := pw.Chromium.Launch()
	require.NoError(t, err)

	page, err := browser.NewPage()
	require.NoError(t, err)

	page.SetDefaultTimeout(3000)
	_, err = page.Goto(url)
	require.NoError(t, err)

	env := &TestEnv{srv, ln, pw, browser, page, url}
	t.Cleanup(func() { env.teardown() })
	return env
}

func (env *TestEnv) teardown() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if env.page != nil {
		env.page.Close()
	}
	if env.browser != nil {
		env.browser.Close()
	}
	if env.pw != nil {
		env.pw.Stop()
	}
	if env.server != nil {
		env.server.Shutdown(ctx)
	}
	if env.listener != nil {
		env.listener.Close()
	}
}
