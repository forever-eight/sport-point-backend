package repository

import (
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

type Authorization interface {
}

// todo new interfaces

type Repository struct {
	Authorization
	db *sqlx.DB
}

func NewRepository() *Repository {
	db, err := NewPostgresDB(Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.dbpass"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		log.Fatalf("init db error: %s", err.Error())
	}
	return &Repository{
		db: db,
	}
}

func (r *Repository) Close() error {
	return r.db.Close()
}
