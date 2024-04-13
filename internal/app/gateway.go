package app

import (
	"context"
	"gateway/internal/adapters/rest"
	"gateway/internal/env"
	"net/http"
)

type App struct {
	httpServer *http.Server
}

func (a *App) Run() error {
	h := rest.New()
	router := h.Router()

	a.httpServer = &http.Server{
		Addr:    env.HttpPort,
		Handler: router,
	}

	if err := a.httpServer.ListenAndServe(); err != nil {
		return err
	}

	return nil
}

func (a *App) Shutdown(ctx context.Context) error {
	if err := a.httpServer.Shutdown(ctx); err != nil {
		return err
	}
	return nil
}
