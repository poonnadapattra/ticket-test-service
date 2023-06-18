package tickets

import "github.com/poonnadapattra/ticket-test-service/internal/entity"

type Service interface {
	GetTicketCount() (res []TicketStatus, err error)
	GetTicket(req ReqTicket) (res ResponseTicket, err error)
	CreateTicket(req entity.Tickets) (err error)
	UpdateTicket(req Tickets) (err error)
	DeleteTicket(req Tickets) (err error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return service{repository}
}

func (s service) GetTicketCount() (res []TicketStatus, err error) {
	return s.repository.GetTicketCount()
}

func (s service) GetTicket(req ReqTicket) (res ResponseTicket, err error) {

	return s.repository.GetTicket(req)
}

func (s service) UpdateTicket(req Tickets) (err error) {
	return s.repository.UpdateTicket(req)
}

func (s service) DeleteTicket(req Tickets) (err error) {
	return s.repository.DeleteTicket(req)
}

func (s service) CreateTicket(req entity.Tickets) (err error) {
	return s.repository.CreateTicket(req)
}
