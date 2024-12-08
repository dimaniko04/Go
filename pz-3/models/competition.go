package models

import "time"

type Competition struct {
	Id        int
	Name      string
	StartDate time.Time
	Duration  int16
	City      string
}
