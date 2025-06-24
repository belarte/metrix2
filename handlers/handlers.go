package handlers

import (
	"net/http"

	"github.com/belarte/metrix2/model"
	"github.com/belarte/metrix2/web/templates"
)

func Home(w http.ResponseWriter, r *http.Request) {
	templates.HomePage().Render(r.Context(), w)
}

func Metrics(w http.ResponseWriter, r *http.Request) {
	templates.MetricsPage().Render(r.Context(), w)
}

func MetricFormFields(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("metric")
	var metric model.Metric
	for _, m := range model.Metrics {
		if m.Title == name {
			metric = m
			break
		}
	}
	templates.MetricFormContent(metric).Render(r.Context(), w)
}
