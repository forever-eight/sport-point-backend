package repository

import (
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
