package web

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/sirupsen/logrus"
)

type (
	Server struct {
		port string
		log  *logrus.Entry
		http *http.Server
	}
)

func NewServer(port string, log *logrus.Entry) *Server {
	s := Server{
		port: port,
		log:  log,
		http: &http.Server{
			Handler:      router(),
			Addr:         fmt.Sprintf(":%s", port),
			WriteTimeout: 15 * time.Second,
			ReadTimeout:  15 * time.Second,
			IdleTimeout:  10 * time.Second,
		},
	}

	return &s
}

func (s *Server) ListenAndServe() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

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
	err := s.http.Shutdown(ctx)
	if err != nil {
		s.log.Errorf("[Server][ListenAndServe] failed to gracefully shutdown the server; %v", err)
	}

	s.log.Info("[Server][ListenAndServe] the server stopped")
}

// Start starts a new web server on a provided port
func Start(serverPort string) {
	srv := &http.Server{
		Handler:      router(),
		Addr:         fmt.Sprintf(":%s", serverPort),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  10 * time.Second,
	}

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	go func() {
		log.Printf("Started http server on port %s", serverPort)
		err := srv.ListenAndServe()
		if err != http.ErrServerClosed {
			log.Fatalf("failed to start http server: %v", err)
		}
	}()

	sig := <-interrupt
	log.Printf("Got interrupt signal %s, going to gracefully shutdown the server\n", sig)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := srv.Shutdown(ctx)
	if err != nil {
		log.Fatalf("Failed to gracefully shutdwon the server; %v", err)
	}

	log.Println("Server gracefully shutdown")
}
