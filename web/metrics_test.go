package web

import (
	"testing"

	"github.com/belarte/metrix2/model"
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

func (m *MetricsPageObject) SelectNewMetric() *MetricsPageObject {
	labels := []string{"+ New Metric"}
	_, err := m.page.GetByLabel("Select metric:").SelectOption(playwright.SelectOptionValues{Labels: &labels})
	require.NoError(m.t, err, "failed to select '+ New Metric'")
	return m
}

func (m *MetricsPageObject) FillMetricForm(title, unit, description string) *MetricsPageObject {
	require.NoError(m.t, m.page.GetByLabel("Title").Fill(title))
	require.NoError(m.t, m.page.GetByLabel("Unit").Fill(unit))
	require.NoError(m.t, m.page.GetByLabel("Description").Fill(description))
	return m
}

func (m *MetricsPageObject) ClickCreate() *MetricsPageObject {
	btn := m.page.GetByRole("button", playwright.PageGetByRoleOptions{Name: "Create"})
	require.NoError(m.t, btn.Click())
	return m
}

type AddValuesPageObject struct {
	page playwright.Page
	t    *testing.T
}

func (p *AddValuesPageObject) SelectMetric(name string) *AddValuesPageObject {
	labels := []string{name}
	_, err := p.page.GetByRole("combobox").SelectOption(playwright.SelectOptionValues{Labels: &labels})
	require.NoError(p.t, err, "failed to select metric '%s'", name)
	return p
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

	env := setupTestEnv(t, "/metrics")
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

func TestCreateNewMetric(t *testing.T) {
	setupSampleMetricsAndValues()
	env := setupTestEnv(t, "/metrics")
	defer env.teardown()
	metricsPage := NewMetricsPageObject(env.page, t)

	newTitle := "Blood Pressure"
	newUnit := "mmHg"
	newDesc := "Systolic/Diastolic blood pressure"

	metricsPage.SelectNewMetric().
		FillMetricForm(newTitle, newUnit, newDesc).
		ClickCreate()

	_, err := env.page.Reload()
	require.NoError(t, err)

	metricsPage.SelectMetric(newTitle)

	err = pwa.Locator(env.page.GetByLabel("Title")).ToHaveValue(newTitle)
	assert.NoError(t, err, "expected title field to have value '%s'", newTitle)
	err = pwa.Locator(env.page.GetByLabel("Unit")).ToHaveValue(newUnit)
	assert.NoError(t, err, "expected unit field to have value '%s'", newUnit)
	err = pwa.Locator(env.page.GetByLabel("Description")).ToHaveValue(newDesc)
	assert.NoError(t, err, "expected description field to have value '%s'", newDesc)
}

func setupSampleMetricsAndValues() {
	model.Metrics = []model.Metric{
		{ID: 1, Title: "Weight", Unit: "kg", Description: "Body weight in kilograms"},
		{ID: 2, Title: "Steps", Unit: "steps", Description: "Daily step count"},
		{ID: 3, Title: "Calories", Unit: "kcal", Description: "Calories burned"},
	}
	model.MetricValues = []model.MetricValue{
		{ID: 1, MetricID: 1, Value: 70.5, Timestamp: 1719500000},
		{ID: 2, MetricID: 1, Value: 71.0, Timestamp: 1719586400},
		{ID: 3, MetricID: 2, Value: 10000, Timestamp: 1719500000},
	}
}

func TestAddValuePageRendersDropdownAndTable(t *testing.T) {
	setupSampleMetricsAndValues()
	metrics := []struct {
		name   string
		values []string
	}{
		{"Weight", []string{"70.5", "71"}},
		{"Steps", []string{"10000"}},
		{"Calories", []string{}},
	}

	for _, metric := range metrics {
		t.Run(metric.name, func(t *testing.T) {
			env := setupTestEnv(t, "/entries")
			defer env.teardown()

			err := env.page.GetByText("Add Value to Metric").WaitFor()
			require.NoError(t, err, "expected 'Add Value to Metric' text to be visible")

			addValuesPage := &AddValuesPageObject{page: env.page, t: t}
			addValuesPage.SelectMetric(metric.name)

			for _, v := range metric.values {
				row := env.page.Locator("table tr").Filter(playwright.LocatorFilterOptions{HasText: v})
				visible, err := row.First().IsVisible()
				assert.NoError(t, err, "expected to find value '%s' in table", v)
				assert.True(t, visible, "expected value '%s' to be visible in table", v)
			}
		})
	}
}

func TestAddValueForm_AddsValueAndShowsFeedback(t *testing.T) {
	setupSampleMetricsAndValues()
	env := setupTestEnv(t, "/entries")
	defer env.teardown()

	err := env.page.GetByText("Add Value to Metric").WaitFor()
	require.NoError(t, err, "expected 'Add Value to Metric' text to be visible")

	addValuesPage := &AddValuesPageObject{page: env.page, t: t}
	addValuesPage.SelectMetric("Weight")

	valueInput := env.page.GetByLabel("Value")
	require.NoError(t, valueInput.Fill("72.3"), "failed to fill value input")
	addBtn := env.page.GetByRole("button", playwright.PageGetByRoleOptions{Name: "Add Value"})
	require.NoError(t, addBtn.Click(), "failed to click Add Value button")

	err = env.page.GetByText("Success! Value added.").WaitFor()
	assert.NoError(t, err, "expected success message after adding value")

	row := env.page.Locator("table tr").Filter(playwright.LocatorFilterOptions{HasText: "72.3"})
	visible, err := row.First().IsVisible()
	assert.NoError(t, err, "expected to find new value in table")
	assert.True(t, visible, "expected new value to be visible in table")

	val, err := valueInput.InputValue()
	assert.NoError(t, err)
	assert.Equal(t, "", val, "expected value input to be cleared after submit")
}
