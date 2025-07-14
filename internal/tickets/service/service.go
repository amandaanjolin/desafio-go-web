package service

import "github.com/amandaanjolin/desafio-go-web/internal/tickets/repository"

type TicketService interface {
	CountByCountry(country string) (int, error)
	CountByTimePeriod() (map[string]int, error)
	PercentageByCountry(country string) (float64, error)
}

type ticketService struct {
	repo repository.TicketRepository
}

func NewService(repo repository.TicketRepository) TicketService {
	return &ticketService{repo: repo}
}

func (s *ticketService) CountByCountry(country string) (int, error) {
	return s.repo.CountByCountry(country)
}

func (s *ticketService) CountByTimePeriod() (map[string]int, error) {
	return s.repo.CountByTimePeriod()
}

func (s *ticketService) PercentageByCountry(country string) (float64, error) {
	return s.repo.PercentageByCountry(country)
}
