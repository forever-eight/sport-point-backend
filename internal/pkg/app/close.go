package app

import (
	"context"
	"log"
)

func (a *App) Close(ctx context.Context) error {
	err := a.HTTPServer.Shutdown(ctx)
	if err != nil {
		log.Fatalf("shutdown server error: %s", err)
	}

	err = a.db.Close()
	if err != nil {
		log.Fatalf("db close error: %s", err)
	}

	return nil
}
