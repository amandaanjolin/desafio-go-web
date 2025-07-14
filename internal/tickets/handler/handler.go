package handler

import (
	"encoding/json"
	"net/http"

	"github.com/amandaanjolin/desafio-go-web/internal/tickets/service"
	"github.com/go-chi/chi/v5"
)

type TicketHandler struct {
	service service.TicketService
}

func NewHandler(service service.TicketService) *TicketHandler {
	return &TicketHandler{service: service}
}

func (h *TicketHandler) CountByCountry(w http.ResponseWriter, r *http.Request) {
	country := chi.URLParam(r, "country")
	if country == "" {
		http.Error(w, "country param is required", http.StatusBadRequest)
		return
	}
	count, err := h.service.CountByCountry(country)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]int{"count": count})
}

func (h *TicketHandler) CountByTimePeriod(w http.ResponseWriter, r *http.Request) {
	result, err := h.service.CountByTimePeriod()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(result)
}

func (h *TicketHandler) PercentageByCountry(w http.ResponseWriter, r *http.Request) {
	country := chi.URLParam(r, "country")
	if country == "" {
		http.Error(w, "country param is required", http.StatusBadRequest)
		return
	}
	percent, err := h.service.PercentageByCountry(country)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]float64{"percentage": percent})
}
