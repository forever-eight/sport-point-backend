package repository

import (
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
}

// todo new interfaces

type Repository struct {
	Authorization
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Close() error {
	return r.db.Close()
}
