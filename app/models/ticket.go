package models

import (
	"gorm.io/gorm"
)

type Ticket struct {
	gorm.Model

	Seat           Seat      `gorm:"foreignKey:SeatRefer;not null" json:"seat"`
	Passenger      Passenger `gorm:"foreignKey:PassengerRefer;not null" json:"passenger"`
	SeatRefer      uint      // Belongs to Seat
	PassengerRefer uint      // Belongs to Passenger

	OrderRefer uint
}

