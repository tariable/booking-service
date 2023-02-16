package http

import (
	"bookingService/config"
	"bookingService/pkg/logger"
	"context"
	"fmt"
	"net/http"
	"time"
)

const (
	shutdownTime = 3 * time.Second
)

type Server struct {
	logger logger.Logger
	server http.Server
	notify chan error
}

func StartServer(config config.HTTP, l logger.Logger, router *http.ServeMux) *Server {
	server := Server{
		logger: l,
		server: http.Server{
			Handler: router,
			Addr:    getAddress(config.Host, config.Port),
		},
		notify: make(chan error),
	}

	go func() {
		server.notify <- server.server.ListenAndServe()
		close(server.notify)
	}()

	return &server
}

func (s *Server) Notify() chan error {
	return s.notify
}

func (s *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), shutdownTime)
	defer cancel()

	s.logger.Info("gracefully shutdown of http server")
	return s.server.Shutdown(ctx)
}

func getAddress(host string, port string) string {
	return fmt.Sprintf("%s:%s", host, port)
}
