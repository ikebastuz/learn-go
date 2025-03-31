package main

import (
	"calculator/utils"
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
	"golang.org/x/time/rate"
	"time"
)

var limiter = rate.NewLimiter(1, 5)

func rateLimitMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !limiter.Allow() {
			http.Error(w, "Too many requests", http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	r := mux.NewRouter()
	r.Use(rateLimitMiddleware)

	r.HandleFunc("/add", utils.DoAdd).Methods(http.MethodPost)
	r.HandleFunc("/subtract", utils.DoSubtract).Methods(http.MethodPost)
	r.HandleFunc("/multiply", utils.DoMultiply).Methods(http.MethodPost)
	r.HandleFunc("/divide", utils.DoDivide).Methods(http.MethodPost)
	r.HandleFunc("/sum", utils.DoSum).Methods(http.MethodPost)

	port := ":1337"
	slog.Info("Server is running on http://localhost" + port)

	go func() {
		if err := http.ListenAndServe(port, r); err != nil {
			slog.Error("Server Error", "Err:", err)
		}
	}()

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	slog.Info("Shutting down server")
	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
}
