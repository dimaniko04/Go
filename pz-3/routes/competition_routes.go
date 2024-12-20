package routes

import (
	"pz-3/handlers"

	"github.com/gorilla/mux"
)

func initCompetitionRoutes(r *mux.Router, handler handlers.CompetitionHandler) {
	s := r.PathPrefix("/competitions").Subrouter()

	s.StrictSlash(true)
	s.HandleFunc("/create", handler.CreatePage).Methods("GET")
	s.HandleFunc("/create", handler.Create).Methods("POST")
	s.HandleFunc("/edit/{id:[0-9]+}", handler.EditPage).Methods("GET")
	s.HandleFunc("/edit/{id:[0-9]+}", handler.Edit).Methods("POST")
	s.HandleFunc("/delete/{id:[0-9]+}", handler.Delete)
	s.HandleFunc("/", handler.GetAll).Methods("GET")
}
