package contacts

import "gorm.io/gorm"

type Contacts struct {
	gorm.Model

	Name    string `json:"name"`
	PhoneNo string `json:"phone_no"`
}
