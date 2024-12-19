package handlers

import (
	"html/template"
	"pz-3/services"
)

type AppHandlers struct {
	ClubHandler      ClubHandler
	DivisionHandler  DivisionHandler
	SportsmanHandler SportsmanHandler
}

func GetAppHandlers(templates *template.Template, services *services.AppServices) AppHandlers {
	return AppHandlers{
		ClubHandler:      &clubHandler{templates: templates, clubService: services.ClubService},
		DivisionHandler:  &divisionHandler{templates: templates, divisionService: services.DivisionService},
		SportsmanHandler: &sportsmanHandler{templates: templates, sportsmanService: services.SportsmanService},
	}
}
