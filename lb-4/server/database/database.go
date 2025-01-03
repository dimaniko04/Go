package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/dimaniko04/Go/lb-4/server/config"
	"github.com/go-sql-driver/mysql"
)

func Db(env *config.Env) (*sql.DB, error) {
	cfg := mysql.Config{
		User:                 env.DbUser,
		Passwd:               env.DbPassword,
		DBName:               env.DbName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return db, fmt.Errorf("failed to connect to the database: %w", err)
	}
	err = db.Ping()
	if err != nil {
		return db, err
	}

	log.Println("DB: Successfully connected!")

	return db, nil
}
