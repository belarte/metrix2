package model

type Metric struct {
	ID          int64
	Title       string
	Unit        string
	Description string
}

var Metrics = []Metric{
	{ID: 1, Title: "Weight", Unit: "kg", Description: "Body weight in kilograms"},
	{ID: 2, Title: "Steps", Unit: "steps", Description: "Daily step count"},
	{ID: 3, Title: "Calories", Unit: "kcal", Description: "Calories burned"},
}

type MetricValue struct {
	ID        int64
	MetricID  int64
	Value     float64
	Timestamp int64
}

var MetricValues = []MetricValue{
	{ID: 1, MetricID: 1, Value: 70.5, Timestamp: 1719500000},
	{ID: 2, MetricID: 1, Value: 71.0, Timestamp: 1719586400},
	{ID: 3, MetricID: 2, Value: 10000, Timestamp: 1719500000},
	{ID: 4, MetricID: 2, Value: 12000, Timestamp: 1719586400},
	{ID: 5, MetricID: 2, Value: 9000, Timestamp: 1719672800},
	{ID: 6, MetricID: 3, Value: 2200, Timestamp: 1719500000},
}
