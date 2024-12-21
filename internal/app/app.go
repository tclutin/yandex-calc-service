package app

import (
	"log/slog"
	"net/http"
)

type App struct {
	logger     *slog.Logger
	httpServer *http.Server
}

func New() *App {
	return &App{}
}

func (a *App) Run() {

}

func (a *App) Stop() {

}
