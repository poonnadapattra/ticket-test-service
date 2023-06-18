package entity

import "gorm.io/gorm"

type Tickets struct {
	gorm.Model

	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
	ContactID   int    `json:"contact_id"`
}
