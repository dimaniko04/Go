package routes

import (
	"pz-3/handlers"

	"github.com/gorilla/mux"
)

func initDivisionRoutes(r *mux.Router, handler handlers.DivisionHandler) {
	s := r.PathPrefix("/divisions").Subrouter()

	s.StrictSlash(true)
	s.HandleFunc("/create", handler.CreatePage).Methods("GET")
	s.HandleFunc("/create", handler.Create).Methods("POST")
	s.HandleFunc("/edit/{id:[0-9]+}", handler.EditPage).Methods("GET")
	s.HandleFunc("/edit/{id:[0-9]+}", handler.Edit).Methods("POST")
	s.HandleFunc("/delete/{id:[0-9]+}", handler.Delete)
	s.HandleFunc("/", handler.GetAll).Methods("GET")
}
