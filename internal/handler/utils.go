package handler

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type Error struct {
	Message string `json:"message"`
}

func FromJSON(i interface{}, r io.Reader) error {
	return json.NewDecoder(r).Decode(i)
}

func Headers(w http.ResponseWriter, statusCode int) {
	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")
}

func Response(w http.ResponseWriter, statusCode int, data interface{}) {
	Headers(w, statusCode)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Fatalf(err.Error())
	}
	return
}
func ErrorResponse(w http.ResponseWriter, statusCode int, error string) {
	Headers(w, statusCode)
	if err := json.NewEncoder(w).Encode(Error{Message: error}); err != nil {
		log.Fatalf(err.Error())
	}
	return
}
