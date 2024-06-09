package server

import (
	"github.com/hameedhub/pismo/internal/handler"
	"github.com/hameedhub/pismo/internal/service"
	"net/http"
)

func (s *Server) Routes(service service.Service) {
	accounts := handler.NewAccountHandler(service.AccountService)

	s.router.HandleFunc("/health", healthCheck).Methods(http.MethodGet)
	s.router.HandleFunc("/account", accounts.Create).Methods(http.MethodPost)
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}
