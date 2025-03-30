package utils

import (
	"encoding/json"
	"net/http"
)

type RequestData struct {
	Number1 *float64 `json:"number1"`
	Number2 *float64 `json:"number2"`
}

type ResponseData struct {
	Result float64 `json:"result,omitempty"`
	Error  string  `json:"error,omitempty"`
}

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

func DoAdd(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSON(w, http.StatusMethodNotAllowed, ResponseData{
			Error: "Method not allowed. Only POST is supported",
		})
		return
	}

	var data RequestData
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		writeJSON(w, http.StatusBadRequest, ResponseData{
			Error: "Invalid JSON format in request body",
		})
		return
	}

	num1, num2, errResp := validateNumbers(data)
	if errResp != nil {
		writeJSON(w, http.StatusBadRequest, errResp)
		return
	}

	result := num1 + num2
	writeJSON(w, http.StatusOK, ResponseData{Result: result})
}
