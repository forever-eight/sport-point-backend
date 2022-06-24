package repository

import (
	"github.com/jmoiron/sqlx"

	"github.com/forever-eight/sport-point-backend.git/internal/app/ds"
)

type Authorization interface {
	CreateUser(user *ds.User) (uint32, error)
	GetUser(email, password string) (*ds.User, error)
}

// todo new interfaces

type Classes interface {
	CreateClass(class *ds.Class) (uint32, error)
}

type Repository struct {
	Classes
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
