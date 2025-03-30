package utils

import (
	"encoding/json"
	"net/http"
)

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func validateNumbers(data RequestData) (float64, float64, *ResponseData) {
	if data.Number1 == nil {
		return 0, 0, &ResponseData{Error: "number1 is required and must be a valid number"}
	}
	if data.Number2 == nil {
		return 0, 0, &ResponseData{Error: "number2 is required and must be a valid number"}
	}
	return *data.Number1, *data.Number2, nil
}
