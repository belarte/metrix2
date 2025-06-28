package handlers

import (
	"fmt"
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
	idStr := r.URL.Query().Get("metric")
	if idStr == "__new__" {
		templates.MetricFields(nil, true).Render(r.Context(), w)
		return
	}
	var metric *model.Metric
	for i, m := range model.Metrics {
		if fmt.Sprintf("%d", m.ID) == idStr {
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

	var nextID int64 = 1
	if len(model.Metrics) > 0 {
		nextID = model.Metrics[len(model.Metrics)-1].ID + 1
	}
	newMetric := model.Metric{ID: nextID, Title: title, Unit: unit, Description: description}
	model.Metrics = append(model.Metrics, newMetric)
	w.Header().Set("Content-Type", "text/html")
	templates.MetricsForm(&newMetric).Render(r.Context(), w)
}

func Entries(w http.ResponseWriter, r *http.Request) {
	var selected *model.Metric
	if len(model.Metrics) > 0 {
		selected = &model.Metrics[0]
	}
	var values []model.MetricValue
	if selected != nil {
		for _, v := range model.MetricValues {
			if v.MetricID == selected.ID {
				values = append(values, v)
			}
		}
	}
	templates.AddValuePage(selected, values).Render(r.Context(), w)
}

func EntriesValuesTable(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("metric")
	var selected *model.Metric
	for i, m := range model.Metrics {
		if fmt.Sprintf("%d", m.ID) == idStr {
			selected = &model.Metrics[i]
			break
		}
	}
	var values []model.MetricValue
	if selected != nil {
		for _, v := range model.MetricValues {
			if v.MetricID == selected.ID {
				values = append(values, v)
			}
		}
	}
	templates.MetricValuesTable(values).Render(r.Context(), w)
}
