package web

import (
	"testing"

	"github.com/playwright-community/playwright-go"
	"github.com/stretchr/testify/assert"
)

var pwa = playwright.NewPlaywrightAssertions()

type MetricsPageObject struct {
	page playwright.Page
}

func NewMetricsPageObject(page playwright.Page) *MetricsPageObject {
	return &MetricsPageObject{page: page}
}

func (m *MetricsPageObject) SelectMetric(name string) error {
	labels := []string{name}
	_, err := m.page.GetByLabel("Select metric:").SelectOption(playwright.SelectOptionValues{Labels: &labels})
	return err
}

type metricTestCase struct {
	name        string
	title       string
	unit        string
	description string
}

func TestSelectMetricShowsFields(t *testing.T) {
	testCases := []metricTestCase{
		{"Weight", "Weight", "kg", "Body weight in kilograms"},
		{"Steps", "Steps", "steps", "Daily step count"},
		{"Calories", "Calories", "kcal", "Calories burned"},
	}

	env := setupTestEnv(t, "http://localhost:8080/metrics")
	defer env.teardown()
	metricsPage := NewMetricsPageObject(env.page)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := metricsPage.SelectMetric(tc.name)
			assert.NoError(t, err, "failed to select metric '%s'", tc.name)

			err = pwa.Locator(env.page.GetByLabel("Title")).ToHaveValue(tc.title)
			assert.NoError(t, err, "expected title field to have value '%s'", tc.title)

			err = pwa.Locator(env.page.GetByLabel("Unit")).ToHaveValue(tc.unit)
			assert.NoError(t, err, "expected unit field to have value '%s'", tc.unit)

			err = pwa.Locator(env.page.GetByLabel("Description")).ToHaveValue(tc.description)
			assert.NoError(t, err, "expected description field to have correct value '%s'", tc.description)
		})
	}
}
