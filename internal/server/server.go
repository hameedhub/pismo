package server

import (
	"context"
	"errors"
	"github.com/gorilla/mux"
	"github.com/hameedhub/pismo/internal/config"
	"github.com/hameedhub/pismo/internal/server/middleware"
	"github.com/hameedhub/pismo/internal/service"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Server struct {
	cfg     config.Config
	router  *mux.Router
	server  *http.Server
	service service.Service
}

func NewServer(cfg config.Config, service *service.Service) *Server {
	router := mux.NewRouter()
	server := &Server{cfg: cfg, router: router}

	router.Use(middleware.LoggerMiddleware)
	router.Use(middleware.HeaderMiddleware)

	server.Routes(*service)
	return server
}

func (s *Server) Start() {
	s.server = &http.Server{
		Addr:    s.cfg.ServerPort,
		Handler: s.router,
	}

	go func() {
		log.Printf("Server is listening on %v", s.cfg.ServerPort)
		if err := s.server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			log.Fatal(err)
		}
	}()
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGINT, syscall.SIGTERM)
	<-stopChan

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := s.server.Shutdown(ctx); err != nil {
		log.Fatalf("shutdown failed: %+v", err)
	}
	log.Println("gracefully stopped")

}
