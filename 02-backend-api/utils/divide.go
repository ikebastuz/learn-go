package utils

import (
	"net/http"
)

func DoDivide(w http.ResponseWriter, r *http.Request) {
	if errRequest := validateRequest(r); errRequest != nil {
		writeJSON(w, http.StatusBadRequest, errRequest)
		return
	}

	var num1, num2, errPayload = validateRequest2Nums(&w, r)

	if errPayload != nil {
		writeJSON(w, http.StatusBadRequest, errPayload)
		return
	}

	if num2 == 0 {
		writeJSON(w, http.StatusBadRequest, ResponseData{
			Error: "Can not divide by 0",
		})

		return
	}

	result := num1 / num2
	writeJSON(w, http.StatusOK, ResponseData{Result: result})
}
