package services

import (
	"hotel/internal/models"
	"hotel/internal/repositories"
)

type TicketService struct {
	repo *repositories.Repo
}

func NewTicketService(repo *repositories.Repo) *TicketService {
	return &TicketService{repo: repo}
}

func (s *TicketService) CreateTicket(ticket models.TicketCreate) (int, error) {
	return s.repo.Ticket.CreateTicket(ticket)
}

func (s *TicketService) GetTicketById(id int) (models.Ticket, error) {
	return s.repo.Ticket.GetTicketById(id)
}

func (s *TicketService) GetAllTickets() ([]models.Ticket, error) {
	return s.repo.Ticket.GetAllTickets()
}

func (s *TicketService) UpdateTicket(id int, ticket models.TicketCreate) error {
	return s.repo.Ticket.UpdateTicket(id, ticket)
}

func (s *TicketService) DeleteTicket(id int) error {
	return s.repo.Ticket.DeleteTicket(id)
}
