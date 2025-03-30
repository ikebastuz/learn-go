package utils

import (
	"net/http"
)

func DoMultiply(w http.ResponseWriter, r *http.Request) {
	if errRequest := validateRequest(r); errRequest != nil {
		writeJSON(w, http.StatusBadRequest, errRequest)
		return
	}

	var num1, num2, errPayload = validateRequest2Nums(&w, r)

	if errPayload != nil {
		writeJSON(w, http.StatusBadRequest, errPayload)
		return
	}

	result := num1 * num2
	writeJSON(w, http.StatusOK, ResponseData{Result: result})
}
