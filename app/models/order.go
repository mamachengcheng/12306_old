package models

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	TradeNo        string `gorm:"not null;unique" json:"trade_no"`

	Schedule       Schedule `gorm:"foreignKey:ScheduleRefer;not null;unique" json:"schedule"`
	OrderStatus    uint     `gorm:"not null;unique" json:"order_status"`
	Seat           Seat     `gorm:"foreignKey:SeatID;not null;unique" json:"seat"`

	PassengerRefer uint
	ScheduleRefer  uint
}
