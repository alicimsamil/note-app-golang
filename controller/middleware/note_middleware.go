package middleware

import (
	"fmt"
	"net/http"
	"time"
)

// Logging This method written to log response time of routes
func Logging(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		handler.ServeHTTP(w, r)
		fmt.Printf("%s %s %s \n", r.Method, r.RequestURI, time.Since(start))
	}
}
