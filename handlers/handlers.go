package handlers

import (
	"net/http"

	"github.com/belarte/metrix2/web/templates"
)

func Home(w http.ResponseWriter, r *http.Request) {
	templates.HomePage().Render(r.Context(), w)
}

func Metrics(w http.ResponseWriter, r *http.Request) {
	templates.MetricsPage().Render(r.Context(), w)
}
