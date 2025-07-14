package repository

import (
	"time"

	"github.com/amandaanjolin/desafio-go-web/model"
)

type TicketRepository interface {
	GetAll() ([]model.Ticket, error)
	CountByCountry(country string) (int, error)
	CountByTimePeriod() (map[string]int, error)
	PercentageByCountry(country string) (float64, error)
}

type ticketRepository struct {
	tickets []model.Ticket
}

func NewRepository(tickets []model.Ticket) TicketRepository {
	return &ticketRepository{tickets: tickets}
}

func (r *ticketRepository) GetAll() ([]model.Ticket, error) {
	return r.tickets, nil
}

func (r *ticketRepository) CountByCountry(country string) (int, error) {
	count := 0
	for _, t := range r.tickets {
		if t.Destination == country {
			count++
		}
	}
	return count, nil
}

func (r *ticketRepository) CountByTimePeriod() (map[string]int, error) {
	result := map[string]int{
		"madrugada": 0, // 00–06
		"manha":     0, // 07–12
		"tarde":     0, // 13–19
		"noite":     0, // 20–23
	}
	for _, t := range r.tickets {
		// Supondo que t.Hour está no formato "15:04"
		h, _ := time.Parse("15:04", t.Hour)
		switch {
		case h.Hour() >= 0 && h.Hour() <= 6:
			result["madrugada"]++
		case h.Hour() >= 7 && h.Hour() <= 12:
			result["manha"]++
		case h.Hour() >= 13 && h.Hour() <= 19:
			result["tarde"]++
		case h.Hour() >= 20 && h.Hour() <= 23:
			result["noite"]++
		}
	}
	return result, nil
}

func (r *ticketRepository) PercentageByCountry(country string) (float64, error) {
	total := len(r.tickets)
	if total == 0 {
		return 0, nil
	}
	count, _ := r.CountByCountry(country)
	return (float64(count) / float64(total)) * 100, nil
}
