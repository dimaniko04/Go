package models

import "time"

type Sportsman struct {
	Id        int
	FirstName string
	LastName  string
	BirthDate time.Time
	Sex       string
	ClubName  string
}
