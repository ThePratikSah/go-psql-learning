package middleware

import (
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
)

// RequireAuth ensures the request has a valid JWT token.
func RequireAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO: Validate JWT token in Phase 5
		token := r.Header.Get("Authorization")
		if token == "" {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// LoggerMiddleware logs request details.
func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Using chi's built-in request ID from the router middleware
		reqID := middleware.GetReqID(r.Context())
		_ = reqID // logged in Phase 5
		next.ServeHTTP(w, r)
	})
}
