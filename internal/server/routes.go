package server

import (
	"github.com/hameedhub/pismo/internal/handler"
	"github.com/hameedhub/pismo/internal/service"
	"net/http"
)

func (s *Server) Routes(service service.Service) {
	accounts := handler.NewAccountHandler(service.AccountService)

	s.router.HandleFunc("/health", healthCheck).Methods(http.MethodGet)

	s.router.HandleFunc("/accounts", accounts.Create).Methods(http.MethodPost)
	s.router.HandleFunc("/accounts/{id:[0-9]+}", accounts.GetById).Methods(http.MethodGet)
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}
