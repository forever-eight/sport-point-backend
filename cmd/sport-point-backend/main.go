package main

import (
	"log"

	"github.com/forever-eight/sport-point-backend.git/internal/app/handler"
	"github.com/forever-eight/sport-point-backend.git/internal/server"
)

func main() {
	srv := new(server.Server)
	handlers := new(handler.Handler)
	if err := srv.Start("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("server error: %s", err.Error())
	}
}
