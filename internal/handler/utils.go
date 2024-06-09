package handler

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func FromJSON(i interface{}, r io.Reader) error {
	return json.NewDecoder(r).Decode(i)
}

func Response(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Fatalf(err.Error())
	}
	return
}
