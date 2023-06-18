package tickets

import (
	"github.com/poonnadapattra/ticket-test-service/internal/entity"
	"gorm.io/gorm"
)

type Tickets struct {
	gorm.Model

	Title          string `json:"title"`
	Description    string `json:"description"`
	Status         string `json:"status"`
	ContactID      int    `json:"contact_id"`
	ContactName    string `json:"contact_name"`
	ContactPhoneNo string `json:"contact_phone_no"`
}

type TicketStatus struct {
	Status string `json:"status"`
	Count  int    `json:"count"`
}

type ReqTicket struct {
	Status  string `json:"status"`
	OrderBy string `json:"order_by"`
	SortBy  string `json:"sort_by"`

	Pagging entity.Pagging `json:"pagging"`
}

type ResponseTicket struct {
	Data    []Tickets      `json:"data"`
	Pagging entity.Pagging `json:"pagging"`
}
