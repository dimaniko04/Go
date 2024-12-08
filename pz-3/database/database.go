package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	user     = "dima"
	password = "dima_Mysql55"
	dbname   = "go"
)

var db *sql.DB

func initializeDb() error {
	connStr := fmt.Sprintf("%s:%s@/%s", user, password, dbname)

	temp, err := sql.Open("mysql", connStr)
	if err != nil {
		return fmt.Errorf("failed to connect to the database: %w", err)
	}

	db = temp

	return nil
}

func Db() (*sql.DB, error) {
	if db == nil {
		if err := initializeDb(); err != nil {
			return nil, err
		}
	}
	return db, nil
}
