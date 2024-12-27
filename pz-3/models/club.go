package models

type Club struct {
	Id      int
	Name    string
	City    string
	Address string
}

type ClubToCreate struct {
	Name    string
	City    string
	Address string
}

type ClubStats struct {
	Name   string
	Gold   int
	Silver int
	Bronze int
	Score  int
}
