package tickets

type Service interface {
	GetTicket() (res []Ticket, err error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return service{repository}
}

func (s service) GetTicket() (res []Ticket, err error) {
	return s.repository.GetTicket()
}
