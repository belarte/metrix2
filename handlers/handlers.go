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
	if name == "__new__" {
		templates.MetricFields(nil, true).Render(r.Context(), w)
		return
	}
	var metric *model.Metric
	for i, m := range model.Metrics {
		if m.Title == name {
			metric = &model.Metrics[i]
			break
		}
	}
	templates.MetricFields(metric, false).Render(r.Context(), w)
}

func CreateMetric(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Invalid form", http.StatusBadRequest)
		return
	}
	title := r.FormValue("title")
	unit := r.FormValue("unit")
	description := r.FormValue("description")
	if title == "" || unit == "" {
		http.Error(w, "Title and Unit are required", http.StatusBadRequest)
		return
	}
	newMetric := model.Metric{Title: title, Unit: unit, Description: description}
	model.Metrics = append(model.Metrics, newMetric)
	w.Header().Set("Content-Type", "text/html")
	templates.MetricsForm(&newMetric).Render(r.Context(), w)
}
