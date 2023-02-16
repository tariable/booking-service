package app

import (
	"bookingService/config"
	"bookingService/internal/adapter/http"
	"bookingService/internal/adapter/repository"
	"bookingService/internal/service"
	httpS "bookingService/pkg/http"
	"bookingService/pkg/logger"
	"bookingService/pkg/metrics"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func Run(config config.Config) {
	// initialize logger
	l := logger.New(config.Logger)

	appName := fmt.Sprintf("%s | %s", config.App.Name, config.App.Version)

	l.Infof("APP: %s IS STARTING", appName)

	// initialize metrics
	m, err := metrics.New(config.Metrics)
	if err != nil {
		l.Errorf("initializing metrics: %s", err)
	}
	l.Info("Metric UP")

	// initialize repository
	r, err := repository.New(l, config.Storage)
	if err != nil {
		l.Errorf("initializing storage: %s", err)
	}
	l.Info("Storage UP")

	// initialize service
	s := service.NewOrderService(l, r, m)

	// initialize router
	router := http.NewRouter(s)

	// initialize http server
	httpServer := httpS.StartServer(config.HTTP, l, router)
	l.Info("HTTP server UP")

	// handle os signal
	interruption := make(chan os.Signal, 1)
	signal.Notify(interruption, os.Interrupt, syscall.SIGTERM)

	// waiting interruption or httpServer error
	select {
	case s := <-interruption:
		l.Errorf("APP INTERRUPTED: %s", fmt.Errorf("error - %s", s.String()))
	case err = <-httpServer.Notify():
		l.Errorf("HTTP LISTEN: %s", err)
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		l.Errorf("HTTP SERVER SHUT DOWN: %s", err)
	}
}
