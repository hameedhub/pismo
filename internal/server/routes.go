package server

import "net/http"

func (s *Server) Routes() {
	s.router.HandleFunc("/health", healthCheck).Methods(http.MethodGet)
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}
