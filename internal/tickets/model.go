package tickets

import (
	"time"

	"github.com/poonnadapattra/ticket-test-service/internal/entity"
)

type Ticket struct {
	ID          int        `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      string     `json:"status"`
	CreatedDate time.Time  `json:"created_date"`
	UpdatedDate time.Time  `json:"updated_date"`
	DeletedDate *time.Time `json:"deleted_date"`
}
type TicketStatus struct {
	Status string `json:"status"`
	Count  int    `json:"count"`
}

type ReqTicket struct {
	Status string `json:"status"`

	Pagging entity.Pagging `json:"pagging"`
}

type ResponseTicket struct {
	Data    []Ticket       `json:"data"`
	Pagging entity.Pagging `json:"pagging"`
}
