package middleware

import (
	"log"
	"net/http"
)

func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.ContentLength)
		// log services request here...
		next.ServeHTTP(w, r)
	})
}
