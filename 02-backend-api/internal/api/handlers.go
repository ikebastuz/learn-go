package api

import (
	"encoding/json"
	"net/http"

	"log/slog"

	"github.com/gorilla/mux"

	"calculator/internal/calculator"
)

type Handler struct {
	logger *slog.Logger
}

func NewHandler(logger *slog.Logger) *Handler {
	return &Handler{logger: logger}
}

func (h *Handler) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/add", h.handleAdd).Methods(http.MethodPost)
	r.HandleFunc("/subtract", h.handleSubtract).Methods(http.MethodPost)
	r.HandleFunc("/multiply", h.handleMultiply).Methods(http.MethodPost)
	r.HandleFunc("/divide", h.handleDivide).Methods(http.MethodPost)
	r.HandleFunc("/sum", h.handleSum).Methods(http.MethodPost)
}

func (h *Handler) handleAdd(w http.ResponseWriter, r *http.Request) {
	var req calculator.TwoNumberRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.handleError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	result, err := calculator.Add(req)
	if err != nil {
		h.handleError(w, http.StatusBadRequest, err.Error())
		return
	}

	h.handleSuccess(w, result)
}

func (h *Handler) handleSubtract(w http.ResponseWriter, r *http.Request) {
	var req calculator.TwoNumberRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.handleError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	result, err := calculator.Subtract(req)
	if err != nil {
		h.handleError(w, http.StatusBadRequest, err.Error())
		return
	}

	h.handleSuccess(w, result)
}

func (h *Handler) handleMultiply(w http.ResponseWriter, r *http.Request) {
	var req calculator.TwoNumberRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.handleError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	result, err := calculator.Multiply(req)
	if err != nil {
		h.handleError(w, http.StatusBadRequest, err.Error())
		return
	}

	h.handleSuccess(w, result)
}

func (h *Handler) handleDivide(w http.ResponseWriter, r *http.Request) {
	var req calculator.TwoNumberRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.handleError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	result, err := calculator.Divide(req)
	if err != nil {
		h.handleError(w, http.StatusBadRequest, err.Error())
		return
	}

	h.handleSuccess(w, result)
}

func (h *Handler) handleSum(w http.ResponseWriter, r *http.Request) {
	var req calculator.SumRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.handleError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	result, err := calculator.Sum(req)
	if err != nil {
		h.handleError(w, http.StatusBadRequest, err.Error())
		return
	}

	h.handleSuccess(w, result)
}

func (h *Handler) handleSuccess(w http.ResponseWriter, result calculator.Response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func (h *Handler) handleError(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(calculator.Response{Error: message})
} 