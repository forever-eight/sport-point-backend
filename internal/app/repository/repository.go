package repository

type Authorization interface {
}

// todo new interfaces

type Repository struct {
	Authorization
}

func NewRepository() *Repository {
	return &Repository{}
}
