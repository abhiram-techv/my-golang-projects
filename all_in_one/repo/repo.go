package repo

import (
	"database/sql"
)

type Repo struct {
	db *sql.DB
}

type RepoImply interface {
}

func NewRepo(db *sql.DB) RepoImply {
	return &Repo{db: db}
}
