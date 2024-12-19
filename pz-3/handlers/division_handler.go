package handlers

import (
	"database/sql"
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"pz-3/models"
	"pz-3/services"
	"pz-3/util"
	"strconv"

	"github.com/gorilla/mux"
)

const division_base_route = "/divisions"

type DivisionHandler interface {
	GetAll(http.ResponseWriter, *http.Request)
	Create(http.ResponseWriter, *http.Request)
	CreatePage(http.ResponseWriter, *http.Request)
	EditPage(http.ResponseWriter, *http.Request)
	Edit(http.ResponseWriter, *http.Request)
	Delete(http.ResponseWriter, *http.Request)
}

type divisionHandler struct {
	templates       *template.Template
	divisionService services.DivisionService
}

func (h *divisionHandler) GetAll(w http.ResponseWriter, _ *http.Request) {
	divisions, err := h.divisionService.GetAll()

	if err != nil {
		http.Error(w, err.Error(),
			http.StatusInternalServerError)
		return
	}

	data := struct {
		Divisions []models.Division
	}{
		Divisions: divisions,
	}

	err = h.templates.ExecuteTemplate(w, "divisions", data)

	if err != nil {
		http.Error(w, err.Error(),
			http.StatusInternalServerError)
	}
}

func parseDivisionFormData(r *http.Request) (*models.DivisionToCreate, error) {
	name := r.FormValue("name")
	sex := r.FormValue("sex")
	minAgeStr := r.FormValue("min_age")
	maxAgeStr := r.FormValue("max_age")
	minWeightStr := r.FormValue("min_weight")
	maxWeightStr := r.FormValue("max_weight")

	minAge, err := strconv.Atoi(minAgeStr)
	if err != nil {
		return nil, errors.New("invalid min age")
	}
	maxAge, err := strconv.Atoi(maxAgeStr)
	if err != nil {
		return nil, errors.New("invalid max age")
	}
	minWeight, errMinWeight := strconv.ParseFloat(minWeightStr, 32)
	maxWeight, errMaxWeight := strconv.ParseFloat(maxWeightStr, 32)

	if errMinWeight != nil && errMaxWeight != nil {
		return nil, errors.New("invalid weight range")
	}

	return &models.DivisionToCreate{
		Name:      name,
		Sex:       sex,
		MinAge:    int16(minAge),
		MaxAge:    int16(maxAge),
		MinWeight: sql.NullFloat64{Float64: minWeight, Valid: errMinWeight == nil},
		MaxWeight: sql.NullFloat64{Float64: maxWeight, Valid: errMaxWeight == nil},
	}, nil
}

var sexOptions = struct {
	SexOptions []util.Option[string]
}{
	SexOptions: []util.Option[string]{
		{Value: "male", Title: "Male"},
		{Value: "female", Title: "Female"},
	},
}

func (h *divisionHandler) CreatePage(w http.ResponseWriter, r *http.Request) {
	err := h.templates.ExecuteTemplate(w, "createDivision", sexOptions)

	if err != nil {
		http.Error(w, err.Error(),
			http.StatusInternalServerError)
	}
}

func (h *divisionHandler) Create(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		http.Error(w, err.Error(),
			http.StatusInternalServerError)
		return
	}
	division, err := parseDivisionFormData(r)

	if err != nil {
		http.Error(w, err.Error(),
			http.StatusBadRequest)
		return
	}

	err = h.divisionService.Create(*division)

	if err != nil {
		http.Error(w, err.Error(),
			http.StatusBadRequest)
		return
	}

	http.Redirect(w, r, division_base_route, 301)
}

func (h *divisionHandler) EditPage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	division, err := h.divisionService.GetOne(id)

	if err != nil {
		http.Error(w, http.StatusText(404), http.StatusNotFound)
		return
	}

	data := struct {
		SexOptions []util.Option[string]
		Division   models.Division
	}{
		SexOptions: sexOptions.SexOptions,
		Division:   division,
	}

	err = h.templates.ExecuteTemplate(w, "editDivision", data)

	if err != nil {
		http.Error(w, err.Error(),
			http.StatusInternalServerError)
	}
}

func (h *divisionHandler) Edit(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	err := r.ParseForm()

	if err != nil {
		http.Error(w, err.Error(),
			http.StatusInternalServerError)
		return
	}
	division, err := parseDivisionFormData(r)

	if err != nil {
		http.Error(w, err.Error(),
			http.StatusBadRequest)
		return
	}

	err = h.divisionService.Edit(id, *division)

	http.Redirect(w, r, division_base_route, 301)
}

func (h *divisionHandler) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	err := h.divisionService.Delete(id)

	if err != nil {
		fmt.Println(err)
	}

	http.Redirect(w, r, division_base_route, 301)
}
