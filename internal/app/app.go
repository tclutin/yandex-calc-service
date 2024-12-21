package app

import (
	"github.com/tclutin/yandex-calc-service/internal/config"
	"github.com/tclutin/yandex-calc-service/pkg/logger"
	"log/slog"
	"net"
	"net/http"
)

type App struct {
	logger     *slog.Logger
	httpServer *http.Server
}

func New() *App {
	cfg := config.New()

	return &App{
		logger: logger.New(),
		httpServer: &http.Server{
			Addr:    net.JoinHostPort(cfg.Host, cfg.Port),
			Handler: nil,
		},
	}
}

func (a *App) Run() {

}

func (a *App) Stop() {

}
