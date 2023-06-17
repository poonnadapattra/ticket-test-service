package tickets

import (
	"math"
	"time"

	"gorm.io/gorm"
)

type Repository interface {
	GetTicketCount() (res []TicketStatus, err error)
	GetTicket(req ReqTicket) (res ResponseTicket, err error)
	CreateTicket(req Ticket) (err error)
	UpdateTicket(req Ticket) (err error)
	DeleteTicket(req Ticket) (err error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return repository{db}
}

func (r repository) GetTicketCount() (res []TicketStatus, err error) {

	query := `Select 'all' as status, count(status) FROM tickets
		union all
		Select 'pending' as status, count(status) FROM tickets
		where status = 'pending'
			union all
		Select 'resolved' as status, count(status) as count FROM tickets
		where status = 'resolved'
			union all
		Select 'accepted' as status, count(status) as count FROM tickets
		where status = 'accepted'
			union all
		Select 'rejected' as status, count(status) as count FROM tickets
		where status = 'rejected'`

	err = r.db.Raw(query).Scan(&res).Error

	return
}

func (r repository) GetTicket(req ReqTicket) (res ResponseTicket, err error) {

	var total int64
	err = r.db.Model(&Ticket{}).Where("status = ? or (? = '' or ? = 'all')", req.Status, req.Status, req.Status).Count(&total).Error
	if err != nil {
		return
	}

	err = r.db.Where("status = ? or (? = '' or ? = 'all')", req.Status, req.Status, req.Status).Order("id asc").Offset((req.Pagging.Size * (req.Pagging.Page - 1))).Limit(req.Pagging.Size).Find(&res.Data).Error
	if err != nil {
		return
	}

	res.Pagging = req.Pagging
	res.Pagging.Total = int(total)
	res.Pagging.TotalPage = int(math.Ceil(float64(total) / float64(req.Pagging.Size)))
	return
}

func (r repository) CreateTicket(req Ticket) (err error) {
	req.CreatedDate = time.Now()
	req.UpdatedDate = time.Now()
	req.Status = "pending"
	err = r.db.Create(&req).Error
	return
}

func (r repository) UpdateTicket(req Ticket) (err error) {
	err = r.db.Model(&Ticket{}).Where("id = ?", req.ID).Updates(Ticket{Status: req.Status, UpdatedDate: time.Now()}).Error
	return
}

func (r repository) DeleteTicket(req Ticket) (err error) {
	err = r.db.Delete(&req).Error
	return
}
