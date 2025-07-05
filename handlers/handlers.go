package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

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
	if id, err := strconv.ParseInt(idStr, 10, 64); err == nil {
		metric = model.FindMetricByID(id)
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

	nextID := model.NextMetricID()
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
	if metricID := r.URL.Query().Get("metric"); metricID != "" {
		for i, m := range model.Metrics {
			if fmt.Sprintf("%d", m.ID) == metricID {
				selected = &model.Metrics[i]
				break
			}
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

func AddEntry(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Invalid form", http.StatusBadRequest)
		return
	}
	metricID := r.FormValue("metric")
	valueStr := r.FormValue("value")
	var selected *model.Metric
	for i, m := range model.Metrics {
		if fmt.Sprintf("%d", m.ID) == metricID {
			selected = &model.Metrics[i]
			break
		}
	}
	feedback := ""
	feedbackClass := ""
	if selected == nil {
		feedback = "Metric not found."
		feedbackClass = "error"
	} else {
		value, err := strconv.ParseFloat(valueStr, 64)
		if err != nil {
			feedback = "Invalid value."
			feedbackClass = "error"
		} else {
			var newID int64 = 1
			if len(model.MetricValues) > 0 {
				newID = model.MetricValues[len(model.MetricValues)-1].ID + 1
			}
			mv := model.MetricValue{
				ID:        newID,
				MetricID:  selected.ID,
				Value:     value,
				Timestamp: time.Now().Unix(),
			}
			model.MetricValues = append(model.MetricValues, mv)
			feedback = "Success! Value added."
			feedbackClass = "success"
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
	templates.AddValueFormAndTable(selected, values, feedback, feedbackClass).Render(r.Context(), w)
}
