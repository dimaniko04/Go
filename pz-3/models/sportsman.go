package models

import "time"

type Sportsman struct {
	Id        int
	FirstName string
	LastName  string
	BirthDate string
	Sex       string
	ClubName  string
	ClubId    int
}

type SportsmanToCreate struct {
	FirstName string
	LastName  string
	BirthDate time.Time
	Sex       string
	ClubId    int
}
