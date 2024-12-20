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
