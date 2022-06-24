package repository

import (
	"fmt"

	"github.com/forever-eight/sport-point-backend.git/internal/app/ds"
)

func (r *Repository) CreateClass(class *ds.Class) (uint32, error) {
	var id uint32
	query := `
		INSERT INTO classes( studio_id, title, type_id, description, weekday, 
		                    started_at, duration, amount)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id;
	`
	err := r.db.QueryRow(query, class.StudioID, class.Title, class.TypeID, class.Description, class.Weekday,
		class.StartedAt, class.Duration, class.Amount).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("failed to insert: %w", err)
	}

	return id, nil
}

func (r *Repository) GetStudios() ([]*ds.Studio, error) {
	var studios []*ds.Studio
	query := `
		SELECT id, title, address
		FROM studio;`
	err := r.db.Get(&studios, query)
	if err != nil {
		return nil, fmt.Errorf("failed to get: %w", err)
	}
	return studios, nil
}

func (r *Repository) GetStudio(id uint32) ([]*ds.Studio, error) {
	var studios []*ds.Studio
	query := `
		SELECT id, title, address
		FROM studio WHERE id = $1;`
	err := r.db.Get(&studios, query, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get: %w", err)
	}
	return studios, nil
}
