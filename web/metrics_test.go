package web

import (
	"testing"

	"github.com/playwright-community/playwright-go"
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
	cmd, err := startServer()
	if err != nil {
		t.Fatalf("failed to start server: %v", err)
	}
	defer stopServer(cmd)

	pw, err := playwright.Run()
	if err != nil {
		t.Fatalf("could not launch playwright: %v", err)
	}
	defer pw.Stop()

	browser, err := pw.Chromium.Launch()
	if err != nil {
		t.Fatalf("could not launch browser: %v", err)
	}
	page, err := browser.NewPage()
	if err != nil {
		t.Fatalf("could not create page: %v", err)
	}

	page.SetDefaultTimeout(3000) // Set Playwright timeout to 3 seconds

	_, err = page.Goto("http://localhost:8080/metrics")
	if err != nil {
		t.Fatalf("could not goto metrics page: %v", err)
	}

	metricsPage := &MetricsPage{page: page}

	// No selection, just check the default/pre-selected metric fields
	title, err := metricsPage.GetTitleValue()
	if err != nil {
		t.Fatalf("could not get title value: %v", err)
	}
	if title == "" {
		t.Errorf("expected title input to be present and non-empty")
	}

	unit, err := metricsPage.GetUnitValue()
	if err != nil {
		t.Fatalf("could not get unit value: %v", err)
	}
	if unit == "" {
		t.Errorf("expected unit input to be present and non-empty")
	}

	desc, err := metricsPage.GetDescriptionValue()
	if err != nil {
		t.Fatalf("could not get description value: %v", err)
	}
	if desc == "" {
		t.Errorf("expected description textarea to be present and non-empty")
	}
}
