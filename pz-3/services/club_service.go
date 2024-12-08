package services

import (
	"pz-3/database"
	"pz-3/models"
)

type ClubService interface {
	GetAll() ([]models.Club, error)
	GetOne(string) (models.Club, error)
	Create(models.ClubToCreate) error
	Edit(string, models.ClubToCreate) error
	Delete(string) error
}

type clubService struct {
}

func (s *clubService) GetAll() ([]models.Club, error) {
	db, err := database.Db()

	rows, err := db.Query("SELECT * FROM clubs")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	clubs := []models.Club{}

	for rows.Next() {
		c := models.Club{}
		err := rows.Scan(&c.Id, &c.Name, &c.City, &c.Address)
		if err != nil {
			return nil, err
		}
		clubs = append(clubs, c)
	}

	return clubs, nil
}

func (s *clubService) GetOne(id string) (models.Club, error) {
	db, err := database.Db()

	row := db.QueryRow("SELECT * FROM clubs WHERE id = ?", id)

	club := models.Club{}
	err = row.Scan(&club.Id, &club.Name, &club.City, &club.Address)

	if err != nil {
		return club, err
	}

	return club, nil
}

func (s *clubService) Create(club models.ClubToCreate) error {
	db, err := database.Db()

	_, err = db.Exec("INSERT INTO clubs (name, city, address) VALUES (?, ?, ?)",
		club.Name, club.City, club.Address)

	if err != nil {
		return err
	}

	return nil
}

func (s *clubService) Edit(id string, club models.ClubToCreate) error {
	db, err := database.Db()

	_, err = db.Exec("UPDATE clubs SET name = ?, city = ?, address = ? WHERE id = ?",
		club.Name, club.City, club.Address, id)

	if err != nil {
		return err
	}

	return nil
}

func (s *clubService) Delete(id string) error {
	db, err := database.Db()

	_, err = db.Exec("DELETE FROM clubs WHERE id = ?", id)

	if err != nil {
		return err
	}

	return nil
}
