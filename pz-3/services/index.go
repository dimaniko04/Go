package services

type AppServices struct {
	ClubService ClubService
}

func GetAppServices() AppServices {
	return AppServices{
		ClubService: &clubService{},
	}
}
