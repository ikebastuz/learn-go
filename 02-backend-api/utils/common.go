package utils

import (
	"encoding/json"
	"net/http"
)

func writeJSON(w http.ResponseWriter, status int, data *ResponseData) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func validateRequest(r *http.Request) *ResponseData {
	if r.Method != http.MethodPost {
		return &ResponseData{
			Error: "Method not allowed. Only POST is supported",
		}
	}
	return nil
}

func validateRequest2Nums(r *http.Request) (float64, float64, *ResponseData) {
	var data RequestData2Nums
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		return 0, 0, &ResponseData{
			Error: "Invalid JSON format in request body",
		}
	}

	if data.Number1 == nil {
		return 0, 0, &ResponseData{
			Error: "number1 is required and must be a valid number",
		}
	}
	if data.Number2 == nil {
		return 0, 0, &ResponseData{
			Error: "number2 is required and must be a valid number",
		}
	}

	return *data.Number1, *data.Number2, nil
}

func validateRequestNumsSlice(r *http.Request) ([]float64, *ResponseData) {
	var data RequestDataNumsSlice
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		return []float64{}, &ResponseData{
			Error: "Invalid JSON format in request body",
		}
	}

	if data.Items == nil {
		return nil, &ResponseData{
			Error: "items is required and must be a valid number array",
		}
	}

	return *data.Items, nil
}
