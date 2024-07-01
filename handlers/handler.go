package handlers

import "database/sql"

type Handlers struct {
	db *sql.DB
}

func NewHandler(db *sql.DB) *Handlers {
	return &Handlers{
		db: db,
	}
}
