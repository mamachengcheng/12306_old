package models

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model

	OrderStatus uint     `gorm:"default:0;not null" json:"order_status"`
	Tickets     []Ticket `gorm:"foreignKey:OrderRefer" json:"tickets"` // Has Many Ticket

	Schedule      Schedule `gorm:"foreignKey:ScheduleRefer;not null" json:"schedule"`
	ScheduleRefer uint     // Belongs to Schedule

	Price uint `gorm:"not null" json:"price"`

	UserRefer uint64
}
