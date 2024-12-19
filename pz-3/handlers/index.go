package handlers

import (
	"html/template"
	"pz-3/services"
)

type AppHandlers struct {
	ClubHandler      ClubHandler
	SportsmanHandler SportsmanHandler
}

func GetAppHandlers(templates *template.Template, services *services.AppServices) AppHandlers {
	return AppHandlers{
		ClubHandler:      &clubHandler{templates: templates, clubService: services.ClubService},
		SportsmanHandler: &sportsmanHandler{templates: templates, sportsmanService: services.SportsmanService},
	}
}
