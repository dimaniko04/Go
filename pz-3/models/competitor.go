package models

import "database/sql"

type Competitor struct {
	Sportsman
	Id              int
	WeightingResult sql.NullFloat64
	DivisionName    sql.NullString
	LapNum          int16
}
