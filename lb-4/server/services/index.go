package services

import "database/sql"

type Services struct {
	ProductService ProductService
}

func GetServices(db *sql.DB) *Services {
	return &Services{
		ProductService: &productService{db},
	}
}
