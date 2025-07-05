package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMetricByID(t *testing.T) {
	Metrics = []Metric{
		{ID: 1, Title: "Weight", Unit: "kg", Description: "Body weight in kilograms"},
		{ID: 2, Title: "Steps", Unit: "steps", Description: "Daily step count"},
	}
	tests := []struct {
		name   string
		id     int64
		exists bool
		title  string
	}{
		{"existing metric 1", 1, true, "Weight"},
		{"existing metric 2", 2, true, "Steps"},
		{"non-existing metric", 999, false, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := FindMetricByID(tt.id)
			if tt.exists {
				assert.NotNil(t, m, "expected to find metric with id %d", tt.id)
				if m != nil {
					assert.Equal(t, tt.title, m.Title, "expected title for id %d", tt.id)
				}
			} else {
				assert.Nil(t, m, "expected not to find metric with id %d", tt.id)
			}
		})
	}
}

func TestNextMetricID(t *testing.T) {
	cases := []struct {
		name     string
		metrics  []Metric
		expected int64
	}{
		{"empty list", []Metric{}, 1},
		{"one metric", []Metric{{ID: 5}}, 6},
		{"multiple metrics", []Metric{{ID: 1}, {ID: 2}, {ID: 10}}, 11},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			Metrics = c.metrics
			id := NextMetricID()
			assert.Equal(t, c.expected, id, "expected next ID for case %q", c.name)
		})
	}
}
