package repository

import (
	"github.com/jmoiron/sqlx"

	"github.com/forever-eight/sport-point-backend.git/internal/app/ds"
)

type Authorization interface {
	CreateUser(user *ds.User) (uint32, error)
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
