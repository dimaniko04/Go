package services

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/dimaniko04/Go/lb-4/server/models"
	"github.com/dimaniko04/Go/lb-4/server/requests"
)

type ProductService interface {
	GetAll() ([]models.Product, error)
	GetOne(string) (*models.Product, error)
	Add(requests.AddProductRequest, string) error
	Delete(string) error
}

type productService struct {
	db *sql.DB
}

func (s *productService) GetAll() ([]models.Product, error) {
	rows, err := s.db.Query("SELECT * FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	products := []models.Product{}

	for rows.Next() {
		p := models.Product{}
		err := rows.Scan(&p.Id,
			&p.Name,
			&p.Description,
			&p.Price,
			&p.Stock,
			&p.ImagePath)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil
}

func (s *productService) GetOne(id string) (*models.Product, error) {
	row := s.db.QueryRow("SELECT * FROM products WHERE id = ?", id)

	product := models.Product{}
	err := row.Scan(&product.Id,
		&product.Name,
		&product.Description,
		&product.Price,
		&product.Stock,
		&product.ImagePath)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (s *productService) Add(product requests.AddProductRequest, imagePath string) error {
	query := "INSERT INTO products (name, image_path, description, price, stock) VALUES (?, ?, ?, ?, ?)"
	_, err := s.db.Exec(query, product.Name, imagePath, product.Description, product.Price, product.Stock)
	if err != nil {
		return err
	}

	return nil
}

func (s *productService) Delete(id string) error {
	product, err := s.GetOne(id)

	if product == nil {
		return fmt.Errorf("No product with id %s", id)
	}

	query := "DELETE FROM products WHERE id = ?"

	_, err = s.db.Exec(query, id)
	if err != nil {
		return err
	}

	err = os.Remove(product.ImagePath)
	if err != nil {
		return err
	}

	return nil
}
