package app

import (
	"context"
	"errors"
	"github.com/tclutin/yandex-calc-service/internal/config"
	"github.com/tclutin/yandex-calc-service/internal/handler"
	"github.com/tclutin/yandex-calc-service/pkg/calc"
	"github.com/tclutin/yandex-calc-service/pkg/logger"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type App struct {
	logger     *slog.Logger
	httpServer *http.Server
}

func New() *App {
	logger := logger.New()

	cfg := config.New()

	calculator := calc.New()

	router := handler.New(logger, calculator)

	return &App{
		logger: logger,
		httpServer: &http.Server{
			Addr:    net.JoinHostPort(cfg.Host, cfg.Port),
			Handler: router.Init(),
		},
	}
}

func (a *App) Run() {
	a.logger.Info("Starting app")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	go func() {
		a.logger.Info("Starting HTTP server")
		err := a.httpServer.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			a.logger.Error(
				"A completion error has occurred", slog.Any("error", err))
			os.Exit(1)
		}
	}()

	<-quit
	a.logger.Info("Received shutdown signal....")
	a.Stop(context.Background())
}

func (a *App) Stop(ctx context.Context) {
	a.logger.Info("Shutting down app...")

	if err := a.httpServer.Shutdown(ctx); err != nil {
		a.logger.Error("Error during server shutdown", slog.Any("error", err))
		os.Exit(1)
	}

	a.logger.Info("App shutdown completed")
}
