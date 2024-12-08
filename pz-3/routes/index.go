package routes

import (
	"html/template"
	"net/http"
	"pz-3/handlers"
	"pz-3/services"

	"github.com/gorilla/mux"
)

func InitRoutes(r *mux.Router, templates *template.Template) {
	var s = services.GetAppServices()
	var h = handlers.GetAppHandlers(templates, &s)

	initClubRotes(r, h.ClubHandler)
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := templates.ExecuteTemplate(w, "index", nil)

		if err != nil {
			http.Error(w, err.Error(),
				http.StatusInternalServerError)
		}
	}).Methods("GET")
}
