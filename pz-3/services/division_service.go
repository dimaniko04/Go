package services

import (
	"pz-3/database"
	"pz-3/models"
)

type DivisionService interface {
	GetAll() ([]models.Division, error)
	GetOne(string) (models.Division, error)
	Create(models.DivisionToCreate) error
	Edit(string, models.DivisionToCreate) error
	Delete(string) error
}

type divisionService struct {
}

func (s *divisionService) GetAll() ([]models.Division, error) {
	db, err := database.Db()

	rows, err := db.Query("SELECT * FROM divisions d ORDER BY id DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	divisions := []models.Division{}

	for rows.Next() {
		d := models.Division{}
		err := rows.Scan(&d.Id,
			&d.Name,
			&d.Sex,
			&d.MinWeight,
			&d.MaxWeight,
			&d.MinAge,
			&d.MaxAge)
		if err != nil {
			return nil, err
		}
		divisions = append(divisions, d)
	}

	return divisions, nil
}

func (s *divisionService) GetOne(id string) (models.Division, error) {
	db, err := database.Db()

	row := db.QueryRow("SELECT * FROM divisions d WHERE d.id = ?", id)

	division := models.Division{}
	err = row.Scan(&division.Id,
		&division.Name,
		&division.Sex,
		&division.MinWeight,
		&division.MaxWeight,
		&division.MinAge,
		&division.MaxAge)

	if err != nil {
		return division, err
	}

	return division, nil
}

func (s *divisionService) Create(division models.DivisionToCreate) error {
	db, err := database.Db()

	_, err = db.Exec("INSERT INTO divisions (name, sex, min_age, max_age, min_weight, max_weight) VALUES (?, ?, ?, ?, ?, ?)",
		division.Name, division.Sex,
		division.MinAge, division.MaxAge,
		division.MinWeight, division.MaxWeight)

	if err != nil {
		return err
	}

	return nil
}

func (s *divisionService) Edit(id string, division models.DivisionToCreate) error {
	db, err := database.Db()

	_, err = db.Exec("UPDATE divisions d SET name = ?, sex = ?, min_age = ?, max_age = ?, min_weight = ?, max_weight = ? WHERE id = ?",
		division.Name, division.Sex,
		division.MinAge, division.MaxAge,
		division.MinWeight, division.MaxWeight, id)

	if err != nil {
		return err
	}

	return nil
}

func (s *divisionService) Delete(id string) error {
	db, err := database.Db()

	_, err = db.Exec("DELETE FROM divisions WHERE id = ?", id)

	if err != nil {
		return err
	}

	return nil
}
