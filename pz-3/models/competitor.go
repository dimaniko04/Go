package models

type Competitor struct {
	Sportsman
	Id              int
	WeightingResult float32
	LapNum          int16
	DivisionName    string
}
