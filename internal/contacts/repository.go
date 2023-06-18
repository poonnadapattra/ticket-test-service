package contacts

import (
	"gorm.io/gorm"
)

type Repository interface {
	GetContact() (res []Contacts, err error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return repository{db}
}

func (r repository) GetContact() (res []Contacts, err error) {

	err = r.db.Order("name asc").Find(&res).Error
	if err != nil {
		return
	}
	return
}
