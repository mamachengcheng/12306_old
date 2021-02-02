package models

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model

	OrderStatus uint `gorm:"default:0" json:"order_status"`

	Tickets []Ticket `gorm:"foreignKey:OrderRefer" json:"tickets"` // Has Many Ticket

	UserRefer uint64
}

type Ticket struct {
	gorm.Model

	Seat           Seat      `gorm:"foreignKey:SeatRefer;not null" json:"seat"`
	Schedule       Schedule  `gorm:"foreignKey:ScheduleRefer;not null" json:"schedule"`
	Passenger      Passenger `gorm:"foreignKey:PassengerRefer;not null" json:"passenger"`
	SeatRefer      uint      // Belongs to Seat
	ScheduleRefer  uint      // Belongs to Schedule
	PassengerRefer uint      // Belongs to Passenger

	OrderRefer uint
}
