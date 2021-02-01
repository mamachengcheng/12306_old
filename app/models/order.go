package models

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model

	OrderStatus uint `gorm:"not null" json:"order_status"`

	Seat          Seat     `gorm:"foreignKey:SeatRefer;not null" json:"seat"`
	Schedule      Schedule `gorm:"foreignKey:ScheduleRefer;not null" json:"schedule"`
	SeatRefer     uint     // Belongs to Seat
	ScheduleRefer uint     // Belongs to Schedule

	PassengerRefer uint
	ScheduleRefer  uint
}
