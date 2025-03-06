package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/alvarezjulia/fizzbuzz/internal/domain"
	"github.com/alvarezjulia/fizzbuzz/internal/service"
)

type Handler struct {
	service *service.FizzBuzzService
}

func NewHandler(service *service.FizzBuzzService) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) FizzBuzzHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req domain.Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	result, err := h.service.ProcessFizzBuzz(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(domain.Response{Result: result})
}

func (h *Handler) StatsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	stats := h.service.GetStats()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stats)
}
