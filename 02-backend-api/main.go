package main

import (
	"calculator/server"
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
	"time"
)

func main() {
	r := mux.NewRouter()

	r.Use(server.RateLimitMiddleware)
	r.Use(server.AuthMiddleware)

	r.HandleFunc("/add", server.DoAdd).Methods(http.MethodPost)
	r.HandleFunc("/subtract", server.DoSubtract).Methods(http.MethodPost)
	r.HandleFunc("/multiply", server.DoMultiply).Methods(http.MethodPost)
	r.HandleFunc("/divide", server.DoDivide).Methods(http.MethodPost)
	r.HandleFunc("/sum", server.DoSum).Methods(http.MethodPost)

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
