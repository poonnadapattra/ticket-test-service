package tickets

type Service interface {
	GetTicketCount() (res []TicketStatus, err error)
	GetTicket(req ReqTicket) (res ResponseTicket, err error)
	CreateTicket(req Ticket) (err error)
	UpdateTicket(req Ticket) (err error)
	DeleteTicket(req Ticket) (err error)
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

func (s service) UpdateTicket(req Ticket) (err error) {
	return s.repository.UpdateTicket(req)
}

func (s service) DeleteTicket(req Ticket) (err error) {
	return s.repository.DeleteTicket(req)
}

func (s service) CreateTicket(req Ticket) (err error) {
	return s.repository.CreateTicket(req)
}
