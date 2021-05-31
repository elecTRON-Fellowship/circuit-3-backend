package db

import "database/sql"

type Repo struct {
	*Queries
	DB *sql.DB
}

func NewRepo(db *sql.DB) *Repo {
	return &Repo{
		Queries: New(db),
		DB:      db,
	}
}
