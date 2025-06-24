package web

import (
	"testing"

	"github.com/playwright-community/playwright-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var pwa = playwright.NewPlaywrightAssertions()

type MetricsPageObject struct {
	page playwright.Page
	t    *testing.T
}

func NewMetricsPageObject(page playwright.Page, t *testing.T) *MetricsPageObject {
	return &MetricsPageObject{page: page, t: t}
}

func (m *MetricsPageObject) SelectMetric(name string) *MetricsPageObject {
	labels := []string{name}
	_, err := m.page.GetByLabel("Select metric:").SelectOption(playwright.SelectOptionValues{Labels: &labels})
	require.NoError(m.t, err, "failed to select metric '%s'", name)
	return m
}

type metricTestCase struct {
	title       string
	unit        string
	description string
}

func TestSelectMetricShowsFields(t *testing.T) {
	testCases := []metricTestCase{
		{"Weight", "kg", "Body weight in kilograms"},
		{"Steps", "steps", "Daily step count"},
		{"Calories", "kcal", "Calories burned"},
	}

	env := setupTestEnv(t, "http://localhost:8080/metrics")
	defer env.teardown()

	metricsPage := NewMetricsPageObject(env.page, t)

	for _, tc := range testCases {
		t.Run(tc.title, func(t *testing.T) {
			metricsPage.SelectMetric(tc.title)

			err := pwa.Locator(env.page.GetByLabel("Title")).ToHaveValue(tc.title)
			assert.NoError(t, err, "expected title field to have value '%s'", tc.title)

			err = pwa.Locator(env.page.GetByLabel("Unit")).ToHaveValue(tc.unit)
			assert.NoError(t, err, "expected unit field to have value '%s'", tc.unit)

			err = pwa.Locator(env.page.GetByLabel("Description")).ToHaveValue(tc.description)
			assert.NoError(t, err, "expected description field to have correct value '%s'", tc.description)
		})
	}
}
