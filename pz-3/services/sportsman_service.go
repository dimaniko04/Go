package services

import (
	"pz-3/database"
	"pz-3/models"
	"pz-3/util"
)

type SportsmanService interface {
	GetAll() ([]models.Sportsman, error)
	GetOne(string) (models.Sportsman, error)
	GetClubOptions() ([]util.Option[int], error)
	Create(models.SportsmanToCreate) error
	Edit(string, models.SportsmanToCreate) error
	Delete(string) error
}

type sportsmanService struct {
}

func (s *sportsmanService) GetAll() ([]models.Sportsman, error) {
	db, err := database.Db()

	rows, err := db.Query("SELECT s.*, c.name FROM sportsmen s INNER JOIN clubs c ON s.club_id = c.id ORDER BY id DESC")
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

func (s *sportsmanService) GetOne(id string) (models.Sportsman, error) {
	db, err := database.Db()

	row := db.QueryRow("SELECT s.*, c.name FROM sportsmen s INNER JOIN clubs c ON s.club_id = c.id WHERE s.id = ?", id)

	sportsman := models.Sportsman{}
	err = row.Scan(&sportsman.Id,
		&sportsman.FirstName,
		&sportsman.LastName,
		&sportsman.BirthDate,
		&sportsman.Sex,
		&sportsman.ClubId,
		&sportsman.ClubName)

	if err != nil {
		return sportsman, err
	}

	return sportsman, nil
}

func (s *sportsmanService) GetClubOptions() ([]util.Option[int], error) {
	db, err := database.Db()

	rows, err := db.Query("SELECT id, name FROM clubs")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	options := []util.Option[int]{}

	for rows.Next() {
		o := util.Option[int]{}
		err := rows.Scan(&o.Value, &o.Title)
		if err != nil {
			return nil, err
		}
		options = append(options, o)
	}

	return options, nil
}

func (s *sportsmanService) Create(sportsman models.SportsmanToCreate) error {
	db, err := database.Db()

	_, err = db.Exec("INSERT INTO sportsmen (first_name, last_name, birth_date, sex, club_id) VALUES (?, ?, ?, ?, ?)",
		sportsman.FirstName, sportsman.LastName,
		sportsman.BirthDate, sportsman.Sex, sportsman.ClubId)

	if err != nil {
		return err
	}

	return nil
}

func (s *sportsmanService) Edit(id string, sportsman models.SportsmanToCreate) error {
	db, err := database.Db()

	_, err = db.Exec("UPDATE sportsmen SET first_name = ?, last_name = ?, birth_date = ?, sex = ?, club_id = ? WHERE id = ?",
		sportsman.FirstName, sportsman.LastName,
		sportsman.BirthDate, sportsman.Sex,
		sportsman.ClubId, id)

	if err != nil {
		return err
	}

	return nil
}

func (s *sportsmanService) Delete(id string) error {
	db, err := database.Db()

	_, err = db.Exec("DELETE FROM sportsmen WHERE id = ?", id)

	if err != nil {
		return err
	}

	return nil
}
