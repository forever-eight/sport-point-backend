package repository

import (
	"database/sql"
	"fmt"

	"github.com/forever-eight/sport-point-backend.git/internal/app/ds"
)

func (r *Repository) CreateUser(user *ds.User) (uint32, error) {
	var id uint32
	query := `
		INSERT INTO users(name, email, password)
		VALUES ($1, $2, $3)
		RETURNING id;
	`
	err := r.db.QueryRow(query, user.Name, user.Email, user.Password).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("failed to insert: %w", err)
	}

	return id, nil
}

func (r *Repository) GetUser(email, password string) (*ds.User, error) {
	query := `
			SELECT id, name, email, password
			FROM users
			WHERE email = $1 AND password = $2;
	`
	var user ds.User
	err := r.db.Get(&user, query, email, password)
	if err != nil {
		return nil, fmt.Errorf("failed to select: %w", err)
	}
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("not found in storage")
	}

	return &user, nil
}
