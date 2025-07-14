package middleware

import (
	"literary-lions-forum/internal/handlers"
	"log"
	"net/http"
)

func RecoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("🚨 Panic recovered: %v", err)

				// Render your custom 500 error page
				handlers.InternalServerErrorHandler(w, r)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
