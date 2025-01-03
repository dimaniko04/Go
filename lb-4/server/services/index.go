package services

import (
	"database/sql"

	"github.com/dimaniko04/Go/lb-4/server/config"
)

type Services struct {
	ProductService ProductService
	AuthService    AuthService
}

func GetServices(db *sql.DB, env *config.Env) *Services {
	return &Services{
		ProductService: &productService{db},
		AuthService:    &authService{db, env.JwtSecret},
	}
}
