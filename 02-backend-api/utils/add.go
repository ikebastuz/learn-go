package utils

import (
	"log/slog"
	"net/http"
)

func DoAdd(w http.ResponseWriter, r *http.Request) {
	if errRequest := validateRequest(r); errRequest != nil {
		writeJSON(w, http.StatusBadRequest, errRequest)
		return
	}

	var num1, num2, errPayload = validateRequest2Nums(r)

	if errPayload != nil {
		writeJSON(w, http.StatusBadRequest, errPayload)
		return
	}

	slog.Info("Adding", "number1", num1, "number2", num2)

	result := num1 + num2

	writeJSON(w, http.StatusOK, &ResponseData{Result: result})
}
