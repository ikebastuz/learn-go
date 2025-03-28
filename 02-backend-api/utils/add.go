package utils

import (
	"encoding/json"
	"net/http"
)

type RequestData struct {
	Number1 float64 `json:"number1"`
	Number2 float64 `json:"number2"`
}

type ResponseData struct {
	Result float64 `json:"result"`
}

func DoAdd(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var data RequestData

	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		http.Error(w, "Failed to parse body", http.StatusBadRequest)
		return
	}

	result := data.Number1 + data.Number2
	response := ResponseData{Result: result}

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(response)
}
