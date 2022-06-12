package main

import (
	"log"

	sportpoint "github.com/forever-eight/sport-point-backend.git"
	handler "github.com/forever-eight/sport-point-backend.git/internal/app/handler"
)

func main() {
	srv := new(sportpoint.Server)
	handlers := new(handler.Handler)
	if err := srv.Start("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("server error: %s", err.Error())
	}
}
