package utils

import (
	"calculator/calculator"
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

	slog.Info("Summing", "numbers", items)
	result := calculator.Sum(items)

	writeJSON(w, http.StatusOK, result)
}
