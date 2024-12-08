package routes

import (
	"pz-3/handlers"

	"github.com/gorilla/mux"
)

func initClubRotes(r *mux.Router, handler handlers.ClubHandler) {
	s := r.PathPrefix("/clubs").Subrouter()

	s.StrictSlash(true)
	s.HandleFunc("/create", handler.Create).Methods("GET", "POST")
	s.HandleFunc("/edit/{id:[0-9]+}", handler.EditPage).Methods("GET")
	s.HandleFunc("/edit/{id:[0-9]+}", handler.Edit).Methods("POST")
	s.HandleFunc("/delete/{id:[0-9]+}", handler.Delete)
	s.HandleFunc("/", handler.GetAll).Methods("GET")
}
