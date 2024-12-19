package handlers

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"pz-3/models"
	"pz-3/services"
	"pz-3/util"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

const sportsman_base_route = "/sportsmen"

type SportsmanHandler interface {
	GetAll(http.ResponseWriter, *http.Request)
	Create(http.ResponseWriter, *http.Request)
	CreatePage(http.ResponseWriter, *http.Request)
	EditPage(http.ResponseWriter, *http.Request)
	Edit(http.ResponseWriter, *http.Request)
	Delete(http.ResponseWriter, *http.Request)
}

type sportsmanHandler struct {
	templates        *template.Template
	sportsmanService services.SportsmanService
}

func (h *sportsmanHandler) GetAll(w http.ResponseWriter, _ *http.Request) {
	sportsmen, err := h.sportsmanService.GetAll()

	if err != nil {
		http.Error(w, err.Error(),
			http.StatusInternalServerError)
		return
	}

	data := struct {
		Sportsmen []models.Sportsman
	}{
		Sportsmen: sportsmen,
	}

	err = h.templates.ExecuteTemplate(w, "sportsmen", data)

	if err != nil {
		http.Error(w, err.Error(),
			http.StatusInternalServerError)
	}
}

func parseFormData(r *http.Request) (*models.SportsmanToCreate, error) {
	firstName := r.FormValue("first_name")
	lastName := r.FormValue("last_name")
	birthDateStr := r.FormValue("birth_date")
	sex := r.FormValue("sex")
	clubIdSrt := r.FormValue("club_id")

	birthDate, err := time.Parse("2006-01-02", birthDateStr)
	if err != nil {
		return nil, errors.New("invalid date")
	}
	clubId, err := strconv.Atoi(clubIdSrt)
	if err != nil {
		return nil, errors.New("invalid club name")
	}

	return &models.SportsmanToCreate{
		FirstName: firstName,
		LastName:  lastName,
		BirthDate: birthDate,
		Sex:       sex,
		ClubId:    clubId,
	}, nil
}

type dropdownOptions struct {
	SexOptions  []util.Option[string]
	ClubOptions []util.Option[int]
}

func getDropDownOptions(h *sportsmanHandler) (dropdownOptions, error) {
	clubOptions, err := h.sportsmanService.GetClubOptions()

	if err != nil {
		return dropdownOptions{}, err
	}

	return dropdownOptions{
		SexOptions: []util.Option[string]{
			{Value: "male", Title: "Male"},
			{Value: "female", Title: "Female"},
		},
		ClubOptions: clubOptions,
	}, nil
}

func (h *sportsmanHandler) CreatePage(w http.ResponseWriter, r *http.Request) {
	options, err := getDropDownOptions(h)

	if err != nil {
		http.Error(w, err.Error(),
			http.StatusInternalServerError)
		return
	}

	err = h.templates.ExecuteTemplate(w, "createSportsman", options)

	if err != nil {
		http.Error(w, err.Error(),
			http.StatusInternalServerError)
	}
}

func (h *sportsmanHandler) Create(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		http.Error(w, err.Error(),
			http.StatusInternalServerError)
		return
	}
	sportsman, err := parseFormData(r)

	if err != nil {
		http.Error(w, err.Error(),
			http.StatusBadRequest)
		return
	}

	err = h.sportsmanService.Create(*sportsman)

	http.Redirect(w, r, sportsman_base_route, 301)
}

func (h *sportsmanHandler) EditPage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	sportsman, err := h.sportsmanService.GetOne(id)

	if err != nil {
		http.Error(w, http.StatusText(404), http.StatusNotFound)
		return
	}

	options, err := getDropDownOptions(h)

	if err != nil {
		http.Error(w, err.Error(),
			http.StatusInternalServerError)
		return
	}

	data := struct {
		dropdownOptions
		Sportsman models.Sportsman
	}{
		Sportsman:       sportsman,
		dropdownOptions: options,
	}

	err = h.templates.ExecuteTemplate(w, "editSportsman", data)

	if err != nil {
		http.Error(w, err.Error(),
			http.StatusInternalServerError)
	}
}

func (h *sportsmanHandler) Edit(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	err := r.ParseForm()

	if err != nil {
		http.Error(w, err.Error(),
			http.StatusInternalServerError)
		return
	}
	sportsman, err := parseFormData(r)

	if err != nil {
		http.Error(w, err.Error(),
			http.StatusBadRequest)
		return
	}

	err = h.sportsmanService.Edit(id, *sportsman)

	http.Redirect(w, r, sportsman_base_route, 301)
}

func (h *sportsmanHandler) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	err := h.sportsmanService.Delete(id)

	if err != nil {
		fmt.Println(err)
	}

	http.Redirect(w, r, sportsman_base_route, 301)
}
