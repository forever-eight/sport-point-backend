package app

import (
	"fmt"
	"net/http"

	"github.com/jmoiron/sqlx"

	"github.com/forever-eight/sport-point-backend.git/internal/app/handlers"
	"github.com/forever-eight/sport-point-backend.git/internal/app/repository"
	"github.com/forever-eight/sport-point-backend.git/internal/app/service"
)

type App struct {
	repos      *repository.Repository
	svc        *service.Service
	handler    *handlers.Handler
	HTTPServer *http.Server
	db         *sqlx.DB
}

func New() (*App, error) {
	a := &App{}

	err := a.initConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to init config: %w", err)
	}

	a.db, err = a.initDB()
	if err != nil {
		return nil, err
	}

	a.repos = repository.NewRepository(a.db)
	a.svc = service.NewService(a.repos)
	a.handler = handlers.NewHandler(a.svc)

	err = a.initHTTPServer()
	if err != nil {
		return nil, fmt.Errorf("failed to init http: %w", err)
	}

	return a, nil
}

func (a *App) Run() error {
	return a.HTTPServer.ListenAndServe()
}
