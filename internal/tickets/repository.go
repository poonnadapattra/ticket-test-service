package tickets

import "gorm.io/gorm"

type Repository interface {
	GetTicket() (res []Ticket, err error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return repository{db}
}

func (r repository) GetTicket() (res []Ticket, err error) {

	if err = r.db.Model(&Ticket{}).Find(&res).Error; err != nil {
		return
	}

	return
}
