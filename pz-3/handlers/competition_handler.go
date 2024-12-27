package handlers

import (
	"database/sql"
	"errors"
	"fmt"
	"html/template"
	"math"
	"net/http"
	"pz-3/models"
	"pz-3/services"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

const competition_base_route = "/competitions"

func getDetailsBaseRoute(id string) string {
	return competition_base_route + "/" + id
}

func getShuffleBaseRoute(competitionId string, divisionId string) string {
	return competition_base_route + "/" + competitionId + "/divisions/" + divisionId
}

type CompetitionHandler interface {
	GetOne(http.ResponseWriter, *http.Request)
	AddCompetitors(http.ResponseWriter, *http.Request)
	AddCompetitorsPage(http.ResponseWriter, *http.Request)
	WeightCompetitor(http.ResponseWriter, *http.Request)
	RemoveCompetitor(http.ResponseWriter, *http.Request)

	GetWinners(http.ResponseWriter, *http.Request)
	GetAllDivisions(http.ResponseWriter, *http.Request)
	GetOneDivision(http.ResponseWriter, *http.Request)
	DeclareVictory(http.ResponseWriter, *http.Request)
	RevokeVictory(http.ResponseWriter, *http.Request)

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

func (h *competitionHandler) GetOne(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["competitionId"]
	competition, err := h.competitionService.GetOne(id)

	if err != nil {
		http.Error(w, err.Error(),
			http.StatusInternalServerError)
		return
	}

	err = h.templates.ExecuteTemplate(w, "competitionDetails", competition)

	if err != nil {
		http.Error(w, err.Error(),
			http.StatusInternalServerError)
	}
}

func (h *competitionHandler) AddCompetitorsPage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["competitionId"]
	sportsmen, err := h.competitionService.GetAvailableSportsmen(id)

	if err != nil {
		http.Error(w, err.Error(),
			http.StatusInternalServerError)
		return
	}

	err = h.templates.ExecuteTemplate(w, "addCompetitors", sportsmen)

	if err != nil {
		http.Error(w, err.Error(),
			http.StatusInternalServerError)
	}
}

func (h *competitionHandler) AddCompetitors(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	competitionId := vars["competitionId"]
	err := r.ParseForm()

	if err != nil {
		http.Error(w, err.Error(),
			http.StatusInternalServerError)
		return
	}
	ids := r.Form["selected_ids"]

	err = h.competitionService.AddCompetitors(competitionId, ids)

	http.Redirect(w, r, getDetailsBaseRoute(competitionId), 301)
}

func (h *competitionHandler) WeightCompetitor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	competitionId := vars["competitionId"]

	err := r.ParseForm()

	if err != nil {
		http.Error(w, err.Error(),
			http.StatusInternalServerError)
		return
	}
	weightingResultStr := r.FormValue("weighting_result")

	weightingResult, err := strconv.ParseFloat(weightingResultStr, 64)
	if err != nil {
		http.Error(w, "invalid weighing result value",
			http.StatusBadRequest)
		return
	}

	err = h.competitionService.WeightCompetitor(id, weightingResult)

	if err != nil {
		http.Error(w, err.Error(),
			http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, getDetailsBaseRoute(competitionId), 301)
}

func (h *competitionHandler) RemoveCompetitor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	competitionId := vars["competitionId"]

	err := h.competitionService.RemoveCompetitor(id)

	if err != nil {
		http.Error(w, err.Error(),
			http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, getDetailsBaseRoute(competitionId), 301)
}

func (h *competitionHandler) GetWinners(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	competitionId := vars["competitionId"]
	winners, err := h.competitionService.GetWinners(competitionId)

	if err != nil {
		http.Error(w, err.Error(),
			http.StatusInternalServerError)
		return
	}

	data := struct {
		Winners []models.Winner
	}{
		Winners: winners,
	}

	err = h.templates.ExecuteTemplate(w, "winners", data)

	if err != nil {
		http.Error(w, err.Error(),
			http.StatusInternalServerError)
	}
}

func (h *competitionHandler) GetAllDivisions(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	competitionId := vars["competitionId"]
	divisions, err := h.competitionService.GetAllDivisions(competitionId)

	if err != nil {
		http.Error(w, err.Error(),
			http.StatusInternalServerError)
		return
	}

	data := struct {
		Id        string
		Divisions []models.CompetitionDivision
	}{
		Id:        competitionId,
		Divisions: divisions,
	}

	err = h.templates.ExecuteTemplate(w, "competitionDivisions", data)

	if err != nil {
		http.Error(w, err.Error(),
			http.StatusInternalServerError)
	}
}

func setAction(i int, j int, laps [][]models.Shuffle) {
	//someone is ahead and current sportsmen already lost
	if laps[i+1][j/2].Id != "" {
		return
	}
	if j%2 == 0 {
		//sportsmen is the last person on the list, can only move him to next lap
		if j+1 >= len(laps[i]) {
			laps[i][j].Action = "declare"
			return
		}
		laps[i][j].Action = "revoke"
		return
	}
	//sportsmen won on previous lap and is waiting for new opponent
	if laps[i][j-1].Id == "" {
		laps[i][j].Action = "revoke"
		return
	}
	laps[i][j-1].Action = "declare"
	laps[i][j].Action = "declare"
}

func (h *competitionHandler) GetOneDivision(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	competitionId := vars["competitionId"]
	divisionId := vars["divisionId"]
	shuffles, err := h.competitionService.GetOneDivision(competitionId, divisionId)

	if err != nil {
		http.Error(w, err.Error(),
			http.StatusInternalServerError)
		return
	}

	lapSize := len(shuffles)
	lapQuantity := int(math.Ceil(math.Log2(float64(lapSize))))

	laps := make([][]models.Shuffle, lapQuantity+1)
	for i := 0; i <= lapQuantity; i++ {
		laps[i] = make([]models.Shuffle, lapSize)
		lapSize = (lapSize / 2) + (lapSize % 2)
	}

	for j, shuffle := range shuffles {
		temp := j
		for i := 0; i <= shuffle.LapNum; i++ {
			laps[i][temp] = shuffle
			temp = temp / 2
		}
		if shuffle.LapNum != 0 {
			prevLap := shuffle.LapNum - 1
			prevJ := j / int(math.Pow(2, float64(prevLap)))
			laps[prevLap][prevJ].Action = "revoke"
		}
		if shuffle.LapNum != lapQuantity {
			prevJ := j / int(math.Pow(2, float64(shuffle.LapNum)))
			setAction(shuffle.LapNum, prevJ, laps)
		}
	}

	data := struct {
		CompetitionId string
		DivisionId    string
		Laps          [][]models.Shuffle
	}{
		Laps:          laps,
		DivisionId:    divisionId,
		CompetitionId: competitionId,
	}

	err = h.templates.ExecuteTemplate(w, "shuffles", data)

	if err != nil {
		http.Error(w, err.Error(),
			http.StatusInternalServerError)
	}
}

func (h *competitionHandler) DeclareVictory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	divisionId := vars["divisionId"]
	competitionId := vars["competitionId"]

	err := h.competitionService.DeclareVictory(id)

	if err != nil {
		http.Error(w, err.Error(),
			http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, getShuffleBaseRoute(competitionId, divisionId), 301)
}

func (h *competitionHandler) RevokeVictory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	divisionId := vars["divisionId"]
	competitionId := vars["competitionId"]

	err := h.competitionService.RevokeVictory(id)

	if err != nil {
		http.Error(w, err.Error(),
			http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, getShuffleBaseRoute(competitionId, divisionId), 301)
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
		Competition: competition.Competition,
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
