package services

type AppServices struct {
	ClubService      ClubService
	SportsmanService SportsmanService
}

func GetAppServices() AppServices {
	return AppServices{
		ClubService:      &clubService{},
		SportsmanService: &sportsmanService{},
	}
}
