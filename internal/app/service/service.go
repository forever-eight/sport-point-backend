package service

import "github.com/forever-eight/sport-point-backend.git/internal/app/repository"

type Authorization interface {
}

// todo new interfaces

type Service struct {
	Authorization
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
