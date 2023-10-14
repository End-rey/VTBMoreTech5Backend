// Package app configures and runs application.
package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/End-rey/VTBMoreTech5Backend/config"
	"github.com/End-rey/VTBMoreTech5Backend/internal/handler"
	"github.com/End-rey/VTBMoreTech5Backend/internal/repository"
	"github.com/End-rey/VTBMoreTech5Backend/internal/service"
	"github.com/End-rey/VTBMoreTech5Backend/pkg/httpserver"
	"github.com/End-rey/VTBMoreTech5Backend/pkg/logger"
	"github.com/End-rey/VTBMoreTech5Backend/pkg/postgres"
)

// Run creates objects via constructors.
func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)

	// Repository
	db, err := postgres.New(cfg.PG.URL)
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
	}

	// HTTP Server
	repos := repository.NewRepositories(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services, l)
	httpServer := httpserver.New(handlers.NewRouter(), httpserver.Port(cfg.HTTP.Port))

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}
}
