package services

import (
	"pz-3/database"
	"pz-3/models"
	"time"
)

type CompetitionService interface {
	GetAll() ([]models.Competition, error)
	GetOne(string) (models.Competition, error)
	Create(models.CompetitionToCreate) error
	Edit(string, models.CompetitionToCreate) error
	Delete(string) error
}

type competitionService struct {
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

func (s *competitionService) GetOne(id string) (models.Competition, error) {
	db, err := database.Db()

	row := db.QueryRow("SELECT * FROM competitions WHERE id = ?", id)

	var temp string
	competition := models.Competition{}
	err = row.Scan(&competition.Id,
		&competition.Name,
		&temp,
		&competition.Duration,
		&competition.City)
	competition.StartDate, err = time.Parse("2006-01-02", temp)

	if err != nil {
		return competition, err
	}

	return competition, nil
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
