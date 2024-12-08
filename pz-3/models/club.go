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
