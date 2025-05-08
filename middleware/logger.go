package middleware

import (
	"log"
	"net/http"	
	"time"
)	

func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// lanjutkan ke handler berikutnya
		next.ServeHTTP(w, r)

		duration := time.Since(start)
		log.Printf("[REQUEST] %s %s - %v\n", r.Method, r.URL.Path, duration)
	})
}