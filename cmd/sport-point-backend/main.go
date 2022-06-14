package main

import (
	"log"

	"github.com/spf13/viper"

	"github.com/forever-eight/sport-point-backend.git/internal/app/handlers"
	"github.com/forever-eight/sport-point-backend.git/internal/app/repository"
	"github.com/forever-eight/sport-point-backend.git/internal/app/service"
	"github.com/forever-eight/sport-point-backend.git/internal/server"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("init config error: %s", err.Error())
	}

	repos := repository.NewRepository()
	svc := service.NewService(repos)
	handler := handlers.NewHandler(svc)

	srv := server.NewServer()
	if err := srv.Start(viper.GetString("port"), handler.InitRoutes()); err != nil {
		log.Fatalf("server error: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
