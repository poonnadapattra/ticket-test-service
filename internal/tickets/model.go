package tickets

import "time"

type Ticket struct {
	ID          int
	Title       string `json:"title"`
	Description string `json:"description"`

	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-"`
}
