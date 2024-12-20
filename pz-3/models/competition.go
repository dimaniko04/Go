package models

import (
	"database/sql"
	"time"
)

type Competition struct {
	Id        int
	Name      string
	StartDate time.Time
	Duration  int
	City      string
}

type CompetitionToCreate struct {
	Name      string
	StartDate time.Time
	Duration  sql.NullInt16
	City      string
}

type Competitor struct {
	Sportsman
	Id              int
	SportsmanId     int
	WeightingResult sql.NullFloat64
	DivisionName    sql.NullString
	LapNum          int16
}

type CompetitionDetails struct {
	Competition
	Competitors []Competitor
}
