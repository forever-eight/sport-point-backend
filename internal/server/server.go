package server

import (
	"context"
	"net/http"
	"time"

	"github.com/forever-eight/sport-point-backend.git/internal/app/repository"
)

const timeout = 5 * time.Second

type Server struct {
	httpServer *http.Server
	r          *repository.Repository
}

func NewServer() *Server {
	return &Server{}

}

func (s *Server) Start(port string, handler http.Handler) error {

	s.httpServer = &http.Server{
		Addr:         ":" + port,
		Handler:      handler,
		ReadTimeout:  timeout,
		WriteTimeout: timeout,
	}

	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
