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
	GetClubsStats() ([]models.ClubStats, error)
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

func (s *clubService) GetClubsStats() ([]models.ClubStats, error) {
	db, err := database.Db()

	rows, err := db.Query(`
		WITH winners AS (
			SELECT ctr.sportsman_id AS id,
				finals.final_lap - (lap_num) + 1 AS place
			FROM competitors ctr
			INNER JOIN (
				SELECT c.id AS competition, c2.division_id AS division, MAX(lap_num) AS final_lap
				FROM competitions c 
				INNER JOIN competitors c2
					ON c.id = c2.competition_id
				WHERE c.start_date + duration < CURDATE() 
				GROUP BY c.id, c2.division_id
			) finals 
				ON finals.competition = ctr.competition_id
				AND finals.division = ctr.division_id
			WHERE finals.final_lap - (lap_num) + 1 <= 3
		)
		SELECT cl.name, 
			COUNT(w1.id) AS gold, 
			COUNT(w2.id) AS silver, 
			COUNT(w3.id) AS bronze,
			COUNT(w1.id) * 3 + 
				COUNT(w2.id) * 2 + 
				COUNT(w3.id) AS score
		FROM clubs cl
		INNER JOIN sportsmen s
			ON s.club_id = cl.id
		LEFT JOIN winners w1
			ON w1.id = s.id
			AND w1.place = 1
		LEFT JOIN winners w2
			ON w2.id = s.id
			AND w2.place = 2
		LEFT JOIN winners w3
			ON w3.id = s.id
			AND w3.place = 3
		GROUP BY cl.name
		ORDER BY score DESC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	clubs := []models.ClubStats{}

	for rows.Next() {
		c := models.ClubStats{}
		err := rows.Scan(
			&c.Name,
			&c.Gold,
			&c.Silver,
			&c.Bronze,
			&c.Score)
		if err != nil {
			return nil, err
		}
		clubs = append(clubs, c)
	}

	return clubs, nil
}
