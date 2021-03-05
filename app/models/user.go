package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model

	Username    string `gorm:"not null;unique" json:"username"`
	Email       string `gorm:"not null;unique" json:"email"`
	MobilePhone string `gorm:"not null;unique" json:"mobile_phone"`

	Password string `gorm:"not null" json:"password"`

	UserInformationID uint

	Passengers []Passenger `gorm:"foreignKey:UserRefer" json:"regular_passengers"` // Has Many Passenger
	Orders     []Order     `gorm:"foreignKey:UserRefer" json:"orders"`             // Has Many Order
}

type Passenger struct {
	gorm.Model

	Name                string    `gorm:"not null" json:"name"`
	CertificateType     uint8     `gorm:"default:0"  json:"certificate_type"`
	Sex                 bool      `gorm:"not null" json:"sex"`
	Birthday            time.Time `gorm:"not null" json:"birthday"`
	Country             string    `gorm:"default:中国CHINA" json:"country"`
	CertificateDeadline time.Time `gorm:"default:'9999-12-31 23:59:59'" json:"certificate_deadline"`
	Certificate         string    `gorm:"not null" json:"certificate"`
	PassengerType       uint8     `gorm:"default:0" json:"passenger_type"`
	MobilePhone         string    `gorm:"not null" json:"mobile_phone"`
	CheckStatus         uint8     `gorm:"default:0" json:"check_status"`
	UserStatus          uint8     `gorm:"default:0" json:"user_status"`

	UserRefer uint64
}


