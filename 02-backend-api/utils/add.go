package utils

import (
	"encoding/json"
	"net/http"
)

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
