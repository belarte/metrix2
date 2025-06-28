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
