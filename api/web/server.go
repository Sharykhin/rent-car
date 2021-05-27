package web

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/sirupsen/logrus"

	"Sharykhin/rent-car/di"
)

type (
	// Server is a http server that implements REST api presentation
	Server struct {
		port string
		log  *logrus.Entry
		http *http.Server
	}
)

// NewServer creates a new server instance
func NewServer(port string) *Server {
	s := Server{
		port: port,
		log:  di.Container.Logger,
		http: &http.Server{
			Handler:      router(),
			Addr:         fmt.Sprintf(":%s", port),
			WriteTimeout: 30 * time.Second,
			ReadTimeout:  30 * time.Second,
			IdleTimeout:  15 * time.Second,
		},
	}

	return &s
}

// ListenAndServe starts web server
func (s *Server) ListenAndServe() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	postgresConn := di.Container.PostgresConn
	err := postgresConn.Connect()
	if err != nil {
		s.log.Fatalf("[Server][ListenAndServe] failed to connect to postgres: %v", err)
	}
	// TODO: @improve check deferred error
	defer postgresConn.Close()

	go func() {
		s.log.Infof("[Server][ListenAndServe] started http server on port %s", s.port)
		err := s.http.ListenAndServe()
		if err != http.ErrServerClosed {
			s.log.Errorf("[Server][ListenAndServe] failed to start http server: %v", err)
		}
	}()

	sig := <-interrupt
	s.log.Infof("[Server][ListenAndServe] got interrupt signal %s, going to gracefully shutdown the server", sig)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = s.http.Shutdown(ctx)
	if err != nil {
		s.log.Errorf("[Server][ListenAndServe] failed to gracefully shutdown the server; %v", err)
	}

	s.log.Info("[Server][ListenAndServe] the server stopped")
}
