package services

type AppServices struct {
	ClubService      ClubService
	SportsmanService SportsmanService
	DivisionService  DivisionService
}

func GetAppServices() AppServices {
	return AppServices{
		ClubService:      &clubService{},
		SportsmanService: &sportsmanService{},
		DivisionService:  &divisionService{},
	}
}
