package model

type Metric struct {
	Title       string
	Unit        string
	Description string
}

var Metrics = []Metric{
	{Title: "Weight", Unit: "kg", Description: "Body weight in kilograms"},
	{Title: "Steps", Unit: "steps", Description: "Daily step count"},
	{Title: "Calories", Unit: "kcal", Description: "Calories burned"},
}
