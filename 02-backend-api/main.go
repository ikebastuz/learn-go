package main

import (
	"calculator/utils"
	"log/slog"
	"net/http"
)

func main() {
	http.HandleFunc("/add", utils.DoAdd)
	http.HandleFunc("/subtract", utils.DoSubtract)
	http.HandleFunc("/multiply", utils.DoMultiply)
	http.HandleFunc("/divide", utils.DoDivide)
	http.HandleFunc("/sum", utils.DoSum)

	port := ":1337"
	slog.Info("Server is running on http://localhost" + port)
	http.ListenAndServe(port, nil)
}
