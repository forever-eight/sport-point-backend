package service

import (
	"fmt"

	"github.com/forever-eight/sport-point-backend.git/internal/app/ds"
)

// todo create studio

func (s *Service) CreateClass(input *ds.ClassInput) (*ds.ClassOutput, error) {
	err := input.Validate()
	if err != nil {
		return nil, fmt.Errorf("service err: CreateClass: %s", err)
	}
	class := input.ToDB()

	id, err := s.repos.CreateClass(class)
	if err != nil {
		return nil, fmt.Errorf("service err: CreateClass: %s", err)
	}
	//class.ToOutput(id)
	return &ds.ClassOutput{ID: 1}, nil
}
