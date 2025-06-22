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

func TestSelectMetricShowsFields(t *testing.T) {
	env := setupTestEnv(t, "http://localhost:8080/metrics")
	defer env.teardown()

	metricsPage := NewMetricsPageObject(env.page)
	err := metricsPage.SelectMetric("Weight")
	assert.NoError(t, err, "failed to select metric 'Weight'")

	err = pwa.Locator(env.page.GetByLabel("Title")).ToHaveValue("Weight")
	assert.NoError(t, err, "expected title field to have value 'Weight'")

	err = pwa.Locator(env.page.GetByLabel("Unit")).ToHaveValue("kg")
	assert.NoError(t, err, "expected unit field to have value 'kg'")

	err = pwa.Locator(env.page.GetByLabel("Description")).ToHaveValue("Body weight in kilograms")
	assert.NoError(t, err, "expected description field to have correct value")
}
