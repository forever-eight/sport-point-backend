package app

import (
	"net/http"
	"time"

	"github.com/spf13/viper"
)

const (
	HTTPTimeout = 10 * time.Second
)

func (a *App) initHTTPServer() error {
	a.HTTPServer = &http.Server{
		Addr:         ":" + viper.GetString("port"),
		Handler:      a.handler.InitRoutes(),
		ReadTimeout:  HTTPTimeout,
		WriteTimeout: HTTPTimeout,
	}

	return nil
}
