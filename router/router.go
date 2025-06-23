package router

import (
	"net/http"

	"github.com/belarte/metrix2/handlers"
)

func New() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.Home)
	mux.HandleFunc("/metrics", handlers.Metrics)
	return mux
}
