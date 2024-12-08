package handlers

import (
	"html/template"
	"pz-3/services"
)

type AppHandlers struct {
	ClubHandler ClubHandler
}

func GetAppHandlers(templates *template.Template, services *services.AppServices) AppHandlers {
	return AppHandlers{
		ClubHandler: &clubHandler{templates: templates, clubService: services.ClubService},
	}
}
