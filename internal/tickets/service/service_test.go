package service

import (
	"testing"

	"github.com/amandaanjolin/desafio-go-web/internal/tickets/repository"
	"github.com/amandaanjolin/desafio-go-web/model"
	"github.com/stretchr/testify/assert"
)

var mockTickets = []model.Ticket{
	{ID: "1", Name: "Ana", Email: "ana@email.com", Destination: "Brasil", Hour: "08:00", Price: 500},
	{ID: "2", Name: "Carlos", Email: "carlos@email.com", Destination: "Brasil", Hour: "04:00", Price: 450},
	{ID: "3", Name: "Julia", Email: "julia@email.com", Destination: "Argentina", Hour: "13:00", Price: 400},
	{ID: "4", Name: "Rafael", Email: "rafael@email.com", Destination: "Argentina", Hour: "22:00", Price: 300},
}

func setupService() TicketService {
	repo := repository.NewRepository(mockTickets)
	return NewService(repo)
}

func TestCountByCountry(t *testing.T) {
	service := setupService()
	count, _ := service.CountByCountry("Brasil")
	assert.Equal(t, 2, count)
	count, _ = service.CountByCountry("Argentina")
	assert.Equal(t, 2, count)
	count, _ = service.CountByCountry("Chile")
	assert.Equal(t, 0, count)
}

func TestCountByTimePeriod(t *testing.T) {
	service := setupService()
	result, _ := service.CountByTimePeriod()
	assert.Equal(t, 1, result["madrugada"]) // 04:00
	assert.Equal(t, 1, result["manha"])     // 08:00
	assert.Equal(t, 1, result["tarde"])     // 13:00
	assert.Equal(t, 1, result["noite"])     // 22:00
}

func TestPercentageByCountry(t *testing.T) {
	service := setupService()
	percent, _ := service.PercentageByCountry("Brasil")
	assert.InDelta(t, 50.0, percent, 0.01)
	percent, _ = service.PercentageByCountry("Argentina")
	assert.InDelta(t, 50.0, percent, 0.01)
	percent, _ = service.PercentageByCountry("Chile")
	assert.InDelta(t, 0.0, percent, 0.01)
}
