package api

import (
	"net/http"
	"time"

	"log/slog"

	"github.com/gorilla/mux"
)

// LoggingMiddleware logs request details
func LoggingMiddleware(logger *slog.Logger) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			// Create a custom response writer to capture the status code
			wrapped := wrapResponseWriter(w)
			next.ServeHTTP(wrapped, r)

			logger.Info("request completed",
				"method", r.Method,
				"path", r.URL.Path,
				"status", wrapped.status,
				"duration", time.Since(start),
				"remote_addr", r.RemoteAddr,
			)
		})
	}
}

// responseWriter is a minimal wrapper for http.ResponseWriter that allows us
// to track the status code
type responseWriter struct {
	http.ResponseWriter
	status int
}

func wrapResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{ResponseWriter: w, status: http.StatusOK}
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
}

// RecoveryMiddleware recovers from panics and logs the error
func RecoveryMiddleware(logger *slog.Logger) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if err := recover(); err != nil {
					logger.Error("panic recovered",
						"error", err,
						"method", r.Method,
						"path", r.URL.Path,
						"remote_addr", r.RemoteAddr,
					)
					http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				}
			}()
			next.ServeHTTP(w, r)
		})
	}
} 