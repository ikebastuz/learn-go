package utils

import (
	"log/slog"
	"net/http"
)

func DoSum(w http.ResponseWriter, r *http.Request) {
	if errRequest := validateRequest(r); errRequest != nil {
		writeJSON(w, http.StatusBadRequest, errRequest)
		return
	}

	var items, errPayload = validateRequestNumsSlice(r)

	if errPayload != nil {
		writeJSON(w, http.StatusBadRequest, errPayload)
		return
	}

	slog.Info("Adding", "items", items)

	var result float64 = 0
	for _, num := range items {
		result += num
	}

	writeJSON(w, http.StatusOK, &ResponseData{Result: result})
}
