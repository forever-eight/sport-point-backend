package service

import (
	"github.com/forever-eight/sport-point-backend.git/internal/app/ds"
	"github.com/forever-eight/sport-point-backend.git/internal/app/repository"
)

type Authorization interface {
	CreateUser(user *ds.UserInput) (*ds.UserOutput, error)

	GenerateToken(email, password string) (string, error)
}

// todo new interfaces

type Classes interface {
	CreateClass(input *ds.ClassInput) (*ds.ClassOutput, error)
}

type Studios interface {
	CreateStudio(input *ds.ClassInput) (*ds.ClassOutput, error)
}

type Service struct {
	Authorization
	repos *repository.Repository
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		repos: repos,
	}
}
