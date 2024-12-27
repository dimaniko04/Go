package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"pz-3/models"
	"pz-3/services"

	"github.com/gorilla/mux"
)

type ClubHandler interface {
	GetAll(http.ResponseWriter, *http.Request)
	Create(http.ResponseWriter, *http.Request)
	EditPage(http.ResponseWriter, *http.Request)
	Edit(http.ResponseWriter, *http.Request)
	Delete(http.ResponseWriter, *http.Request)
	GetAllClubsStats(http.ResponseWriter, *http.Request)
}

type clubHandler struct {
	templates   *template.Template
	clubService services.ClubService
}

func (h *clubHandler) GetAll(w http.ResponseWriter, _ *http.Request) {
	clubs, err := h.clubService.GetAll()

	if err != nil {
		http.Error(w, err.Error(),
			http.StatusInternalServerError)
		return
	}

	data := struct {
		Clubs []models.Club
	}{
		Clubs: clubs,
	}

	err = h.templates.ExecuteTemplate(w, "clubs", data)

	if err != nil {
		http.Error(w, err.Error(),
			http.StatusInternalServerError)
	}
}

func (h *clubHandler) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		err := h.templates.ExecuteTemplate(w, "createClub", nil)

		if err != nil {
			http.Error(w, err.Error(),
				http.StatusInternalServerError)
		}
		return
	}

	err := r.ParseForm()

	if err != nil {
		http.Error(w, err.Error(),
			http.StatusInternalServerError)
	}
	name := r.FormValue("name")
	city := r.FormValue("city")
	address := r.FormValue("address")

	err = h.clubService.Create(models.ClubToCreate{
		Name:    name,
		City:    city,
		Address: address,
	})

	http.Redirect(w, r, "/clubs", 301)
}

func (h *clubHandler) EditPage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	club, err := h.clubService.GetOne(id)

	if err != nil {
		http.Error(w, http.StatusText(404), http.StatusNotFound)
		return
	}

	data := struct {
		Club models.Club
	}{
		Club: club,
	}

	err = h.templates.ExecuteTemplate(w, "editClub", data)

	if err != nil {
		http.Error(w, err.Error(),
			http.StatusInternalServerError)
	}
}

func (h *clubHandler) Edit(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	err := r.ParseForm()

	if err != nil {
		http.Error(w, err.Error(),
			http.StatusInternalServerError)
	}
	name := r.FormValue("name")
	city := r.FormValue("city")
	address := r.FormValue("address")

	err = h.clubService.Edit(id, models.ClubToCreate{
		Name:    name,
		City:    city,
		Address: address,
	})

	http.Redirect(w, r, "/clubs", 301)
}

func (h *clubHandler) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	err := h.clubService.Delete(id)

	if err != nil {
		fmt.Println(err)
	}

	http.Redirect(w, r, "/clubs", 301)
}

func (h *clubHandler) GetAllClubsStats(w http.ResponseWriter, _ *http.Request) {
	clubsStats, err := h.clubService.GetClubsStats()

	if err != nil {
		http.Error(w, err.Error(),
			http.StatusInternalServerError)
		return
	}

	data := struct {
		ClubsStats []models.ClubStats
	}{
		ClubsStats: clubsStats,
	}

	err = h.templates.ExecuteTemplate(w, "leaderboard", data)

	if err != nil {
		http.Error(w, err.Error(),
			http.StatusInternalServerError)
	}
}
