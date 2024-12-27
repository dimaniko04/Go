package services

import (
	"pz-3/database"
	"pz-3/models"
	"time"
)

type CompetitionService interface {
	GetOne(id string) (models.CompetitionDetails, error)
	GetAvailableSportsmen(competitionId string) ([]models.Sportsman, error)
	AddCompetitors(competitionId string, ids []string) error
	WeightCompetitor(id string, weight float64) error
	RemoveCompetitor(id string) error
	GetWinners(competitionId string) ([]models.Winner, error)
	GetAllDivisions(competitionId string) ([]models.CompetitionDivision, error)
	GetOneDivision(competitionId string, divisionId string) ([]models.Shuffle, error)
	DeclareVictory(competitorId string) error
	RevokeVictory(competitorId string) error

	GetAll() ([]models.Competition, error)
	Create(models.CompetitionToCreate) error
	Edit(string, models.CompetitionToCreate) error
	Delete(string) error
}

type competitionService struct {
}

func getCompetitors(competitionId string) ([]models.Competitor, error) {
	db, err := database.Db()

	rows, err := db.Query("SELECT ctr.id, ctr.sportsman_id, s.first_name, s.last_name, s.sex, s.birth_date, cl.name, ctr.weighting_result, d.name, ctr.lap_num  FROM competitors ctr INNER JOIN sportsmen s ON s.id = ctr.sportsman_id LEFT JOIN divisions d ON ctr.division_id = d.id INNER JOIN clubs cl ON cl.id = s.club_id WHERE ctr.competition_id = ? ORDER BY id DESC",
		competitionId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	competitors := []models.Competitor{}

	for rows.Next() {
		c := models.Competitor{}
		err := rows.Scan(&c.Id,
			&c.SportsmanId,
			&c.FirstName,
			&c.LastName,
			&c.Sex,
			&c.BirthDate,
			&c.ClubName,
			&c.WeightingResult,
			&c.DivisionName,
			&c.LapNum)

		if err != nil {
			return nil, err
		}
		competitors = append(competitors, c)
	}

	return competitors, nil
}

func (s *competitionService) GetOne(id string) (models.CompetitionDetails, error) {
	db, err := database.Db()

	row := db.QueryRow("SELECT * FROM competitions WHERE id = ?", id)

	var temp string
	competition := models.CompetitionDetails{}
	err = row.Scan(&competition.Id,
		&competition.Name,
		&temp,
		&competition.Duration,
		&competition.City)
	competition.StartDate, err = time.Parse("2006-01-02", temp)

	if err != nil {
		return competition, err
	}

	competitors, err := getCompetitors(id)

	if err != nil {
		return competition, err
	}
	competition.Competitors = competitors

	return competition, nil
}

func (s *competitionService) GetAvailableSportsmen(competitionId string) ([]models.Sportsman, error) {
	db, err := database.Db()

	rows, err := db.Query("SELECT s.*, cl.name FROM sportsmen s LEFT JOIN competitors c ON c.sportsman_id = s.id AND c.competition_id = ? INNER JOIN clubs cl ON cl.id = s.club_id WHERE c.id IS NULL", competitionId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	sportsmen := []models.Sportsman{}

	for rows.Next() {
		sp := models.Sportsman{}
		err := rows.Scan(&sp.Id,
			&sp.FirstName,
			&sp.LastName,
			&sp.BirthDate,
			&sp.Sex,
			&sp.ClubId,
			&sp.ClubName)
		if err != nil {
			return nil, err
		}
		sportsmen = append(sportsmen, sp)
	}

	return sportsmen, nil
}

func (s *competitionService) AddCompetitors(competitionId string, ids []string) error {
	db, err := database.Db()

	queryStr := "INSERT INTO competitors (sportsman_id, competition_id) VALUES "
	args := []interface{}{}

	for _, id := range ids {
		queryStr += "(?, ?),"
		args = append(args, id, competitionId)
	}
	queryStr = queryStr[0 : len(queryStr)-1]

	_, err = db.Exec(queryStr, args...)

	if err != nil {
		return err
	}

	return nil
}

func (s *competitionService) RemoveCompetitor(id string) error {
	db, err := database.Db()

	_, err = db.Exec("DELETE FROM competitors WHERE id = ?", id)

	if err != nil {
		return err
	}

	return nil
}

func (s *competitionService) WeightCompetitor(id string, weight float64) error {
	db, err := database.Db()

	row := db.QueryRow("SELECT d.id FROM competitors c INNER JOIN competitions cn ON cn.id = c.competition_id INNER JOIN sportsmen s ON s.id = c.sportsman_id INNER JOIN divisions d ON d.sex = s.sex AND TIMESTAMPDIFF(YEAR, s.birth_date, cn.start_date) >= d.min_age AND TIMESTAMPDIFF(YEAR, s.birth_date, cn.start_date) <= d.max_age AND (d.min_weight IS NULL OR ? > d.min_weight) AND (d.max_weight IS NULL OR ? <= d.max_weight) WHERE c.id = ?",
		weight, weight, id)

	var divisionId string
	err = row.Scan(&divisionId)

	if err != nil {
		return err
	}

	_, err = db.Exec("UPDATE competitors SET weighting_result = ?, division_id = ? WHERE id = ?",
		weight, divisionId, id)

	if err != nil {
		return err
	}

	return nil
}

func (s *competitionService) GetAll() ([]models.Competition, error) {
	db, err := database.Db()

	rows, err := db.Query("SELECT * FROM competitions ORDER BY id DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	competitions := []models.Competition{}

	for rows.Next() {
		var temp string
		c := models.Competition{}
		err := rows.Scan(&c.Id,
			&c.Name,
			&temp,
			&c.Duration,
			&c.City)

		c.StartDate, err = time.Parse("2006-01-02", temp)

		if err != nil {
			return nil, err
		}
		competitions = append(competitions, c)
	}

	return competitions, nil
}

func (s *competitionService) Create(competition models.CompetitionToCreate) error {
	db, err := database.Db()

	_, err = db.Exec("INSERT INTO competitions (name, start_date, duration, city) VALUES (?, ?, ?, ?)",
		competition.Name, competition.StartDate,
		competition.Duration, competition.City)

	if err != nil {
		return err
	}

	return nil
}

func (s *competitionService) Edit(id string, competition models.CompetitionToCreate) error {
	db, err := database.Db()

	_, err = db.Exec("UPDATE competitions SET name = ?, start_date = ?, duration = ?, city = ? WHERE id = ?",
		competition.Name, competition.StartDate,
		competition.Duration, competition.City, id)

	if err != nil {
		return err
	}

	return nil
}

func (s *competitionService) Delete(id string) error {
	db, err := database.Db()

	_, err = db.Exec("DELETE FROM competitions WHERE id = ?", id)

	if err != nil {
		return err
	}

	return nil
}

func (s *competitionService) GetWinners(competitionId string) ([]models.Winner, error) {
	db, err := database.Db()

	rows, err := db.Query(`
		SELECT s.first_name,
			s.last_name,
			d.name,
			finals.final_lap - (lap_num) + 1 AS place
		FROM competitors ctr
		INNER JOIN sportsmen s 
			ON s.id = ctr.sportsman_id 
		INNER JOIN divisions d
			ON d.id = ctr.division_id
		INNER JOIN (
			SELECT c2.division_id AS division, MAX(lap_num) AS final_lap
			FROM competitions c 
			INNER JOIN competitors c2
				ON c.id = c2.competition_id
			WHERE c.id = ?
			GROUP BY c2.division_id
		) finals 
			ON finals.division = ctr.division_id
		WHERE finals.final_lap - (lap_num) + 1 <= 3
	`, competitionId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	winners := []models.Winner{}

	for rows.Next() {
		w := models.Winner{}
		err := rows.Scan(&w.FirstName,
			&w.LastName,
			&w.DivisionName,
			&w.Place)
		if err != nil {
			return nil, err
		}
		winners = append(winners, w)
	}

	return winners, nil
}

func (s *competitionService) GetAllDivisions(competitionId string) ([]models.CompetitionDivision, error) {
	db, err := database.Db()

	rows, err := db.Query(`
		SELECT d.id, d.name, COUNT(c2.id) AS participants
		FROM competitions c
		INNER JOIN competitors c2 
			ON c2.competition_id = c.id
		INNER JOIN divisions d
			ON d.id = c2.division_id
		WHERE c.id = ?
		GROUP BY d.name
		ORDER BY COUNT(c2.id)
	`, competitionId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	divisions := []models.CompetitionDivision{}

	for rows.Next() {
		cd := models.CompetitionDivision{}
		err := rows.Scan(&cd.Id,
			&cd.Name,
			&cd.SportsmenQuantity)
		if err != nil {
			return nil, err
		}
		divisions = append(divisions, cd)
	}

	return divisions, nil
}

func (s *competitionService) GetOneDivision(competitionId string, divisionId string) ([]models.Shuffle, error) {
	db, err := database.Db()

	rows, err := db.Query(`
		SELECT c.id, s.first_name, s.last_name, c2.name, c.lap_num 
		FROM competitors c 
		INNER JOIN sportsmen s 
			ON s.id = c.sportsman_id 
		INNER JOIN clubs c2 
			ON c2.id = s.club_id
		WHERE c.division_id = ?
			AND c.competition_id = ?
	`, divisionId, competitionId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	shuffles := []models.Shuffle{}

	for rows.Next() {
		s := models.Shuffle{}
		err := rows.Scan(&s.Id,
			&s.FirstName,
			&s.LastName,
			&s.ClubName,
			&s.LapNum)
		if err != nil {
			return nil, err
		}
		shuffles = append(shuffles, s)
	}

	return shuffles, nil
}

func (s *competitionService) DeclareVictory(competitorId string) error {
	db, err := database.Db()

	_, err = db.Exec("UPDATE competitors SET lap_num = lap_num + 1 WHERE id = ?",
		competitorId)

	if err != nil {
		return err
	}

	return nil
}

func (s *competitionService) RevokeVictory(competitorId string) error {
	db, err := database.Db()

	_, err = db.Exec("UPDATE competitors SET lap_num = lap_num -1 WHERE id = ?",
		competitorId)

	if err != nil {
		return err
	}

	return nil
}
