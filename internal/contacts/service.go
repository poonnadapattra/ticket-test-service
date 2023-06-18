package contacts

type Service interface {
	GetContact() (res []Contacts, err error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return service{repository}
}

func (s service) GetContact() (res []Contacts, err error) {
	return s.repository.GetContact()
}
