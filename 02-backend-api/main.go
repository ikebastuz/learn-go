package main

import (
	"calculator/utils"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/add", utils.DoAdd)
	http.HandleFunc("/subtract", utils.DoSubtract)
	port := ":1337"
	log.Println("Server is running on http://localhost" + port)
	log.Fatal(http.ListenAndServe(port, nil))
}
