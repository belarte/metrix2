package web

import (
	"testing"

	"github.com/playwright-community/playwright-go"
	"github.com/stretchr/testify/require"
)

// MetricsPage models the metric management page for Playwright tests.
type MetricsPage struct {
	page playwright.Page
}

func (mp *MetricsPage) SelectMetric(name string) error {
	_, err := mp.page.SelectOption("#metric-select", playwright.SelectOptionValues{Values: &[]string{name}})
	return err
}

func (mp *MetricsPage) GetTitleValue() (string, error) {
	return mp.page.InputValue("#metric-title")
}

func (mp *MetricsPage) GetUnitValue() (string, error) {
	return mp.page.InputValue("#metric-unit")
}

func (mp *MetricsPage) GetDescriptionValue() (string, error) {
	return mp.page.InputValue("#metric-description")
}

func TestSelectMetricShowsFields(t *testing.T) {
	env := setupTestEnv(t, "http://localhost:8080/metrics")
	defer env.teardown()

	metricsPage := &MetricsPage{page: env.page}

	title, err := metricsPage.GetTitleValue()
	require.NoError(t, err, "could not get title value")
	require.NotEmpty(t, title, "expected title input to be present and non-empty")

	unit, err := metricsPage.GetUnitValue()
	require.NoError(t, err, "could not get unit value")
	require.NotEmpty(t, unit, "expected unit input to be present and non-empty")

	desc, err := metricsPage.GetDescriptionValue()
	require.NoError(t, err, "could not get description value")
	require.NotEmpty(t, desc, "expected description textarea to be present and non-empty")
}
