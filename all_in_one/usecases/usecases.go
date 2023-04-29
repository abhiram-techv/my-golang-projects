package Usecases

import (
	"database/sql"
)

type Usecases struct {
}

type UsecaseImply interface {
}

func NewUsecases(db *sql.DB) UsecaseImply {
	return &Usecases{}
}
