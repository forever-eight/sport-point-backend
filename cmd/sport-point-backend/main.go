package main

import (
	"log"

	"github.com/forever-eight/sport-point-backend.git/internal/pkg/app"
)

func main() {
	ap, err := app.New()
	if err != nil {
		log.Fatalf("app creating error: %s", err.Error())
	}

	err = ap.Run()
	if err != nil {
		log.Fatalf("app run error: %s", err.Error())
	}
}
