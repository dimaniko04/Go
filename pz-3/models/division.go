package models

import "database/sql"

type Division struct {
	Id        int
	Name      string
	Sex       string
	MinWeight sql.NullFloat64
	MaxWeight sql.NullFloat64
	MinAge    int16
	MaxAge    int16
}

type DivisionToCreate struct {
	Name      string
	Sex       string
	MinWeight sql.NullFloat64
	MaxWeight sql.NullFloat64
	MinAge    int16
	MaxAge    int16
}
