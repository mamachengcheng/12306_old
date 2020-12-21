package models

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	TradeNo     string    `gorm:"not null;unique" json:"trade_no"`
	Passenger   Passenger `gorm:"not null;unique" json:"passenger"`
	User        User      `gorm:"not null;unique" json:"user"`
	Schedule    Schedule  `gorm:"not null;unique" json:"schedule"`
	OrderStatus uint      `gorm:"not null;unique" json:"order_status"`
	Seat        Seat      `gorm:"not null;unique" json:"seat"`
}
