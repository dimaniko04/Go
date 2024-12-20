package handlers

import (
	"database/sql"
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"pz-3/models"
	"pz-3/services"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

const competition_base_route = "/competitions"

type CompetitionHandler interface {
	GetAll(http.ResponseWriter, *http.Request)
	Create(http.ResponseWriter, *http.Request)
	CreatePage(http.ResponseWriter, *http.Request)
	EditPage(http.ResponseWriter, *http.Request)
	Edit(http.ResponseWriter, *http.Request)
	Delete(http.ResponseWriter, *http.Request)
}

type competitionHandler struct {
	templates          *template.Template
	competitionService services.CompetitionService
}

func (h *competitionHandler) GetAll(w http.ResponseWriter, _ *http.Request) {
	competitions, err := h.competitionService.GetAll()

	if err != nil {
		http.Error(w, err.Error(),
			http.StatusInternalServerError)
		return
	}

	data := struct {
		Competitions []models.Competition
	}{
		Competitions: competitions,
	}

	err = h.templates.ExecuteTemplate(w, "competitions", data)

	if err != nil {
		http.Error(w, err.Error(),
			http.StatusInternalServerError)
	}
}

func parseCompetitionFormData(r *http.Request) (*models.CompetitionToCreate, error) {
	name := r.FormValue("name")
	startDateStr := r.FormValue("start_date")
	durationStr := r.FormValue("duration")
	city := r.FormValue("city")

	startDate, err := time.Parse("2006-01-02", startDateStr)
	if err != nil {
		return nil, errors.New("invalid date")
	}
	duration, err := strconv.Atoi(durationStr)
	if err != nil {
		return nil, errors.New("invalid duration")
	}

	return &models.CompetitionToCreate{
		Name:      name,
		StartDate: startDate,
		Duration:  sql.NullInt16{Int16: int16(duration | 0), Valid: err == nil},
		City:      city,
	}, nil
}

func (h *competitionHandler) CreatePage(w http.ResponseWriter, r *http.Request) {
	err := h.templates.ExecuteTemplate(w, "createCompetition", nil)

	if err != nil {
		http.Error(w, err.Error(),
			http.StatusInternalServerError)
	}
}

func (h *competitionHandler) Create(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		http.Error(w, err.Error(),
			http.StatusInternalServerError)
		return
	}
	competition, err := parseCompetitionFormData(r)

	if err != nil {
		http.Error(w, err.Error(),
			http.StatusBadRequest)
		return
	}

	err = h.competitionService.Create(*competition)

	http.Redirect(w, r, competition_base_route, 301)
}

func (h *competitionHandler) EditPage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	competition, err := h.competitionService.GetOne(id)

	if err != nil {
		http.Error(w, http.StatusText(404), http.StatusNotFound)
		return
	}

	data := struct {
		Competition models.Competition
	}{
		Competition: competition,
	}

	err = h.templates.ExecuteTemplate(w, "editCompetition", data)

	if err != nil {
		http.Error(w, err.Error(),
			http.StatusInternalServerError)
	}
}

func (h *competitionHandler) Edit(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	err := r.ParseForm()

	if err != nil {
		http.Error(w, err.Error(),
			http.StatusInternalServerError)
		return
	}
	competition, err := parseCompetitionFormData(r)

	if err != nil {
		http.Error(w, err.Error(),
			http.StatusBadRequest)
		return
	}

	err = h.competitionService.Edit(id, *competition)

	if err != nil {
		http.Error(w, err.Error(),
			http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, competition_base_route, 301)
}

func (h *competitionHandler) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	err := h.competitionService.Delete(id)

	if err != nil {
		fmt.Println(err)
	}

	http.Redirect(w, r, competition_base_route, 301)
}
