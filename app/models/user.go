package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model

	UserName    string `gorm:"not null;unique" json:"user_name"`
	Email       string `gorm:"not null;unique" json:"email"`
	MobilePhone string `json:"mobile_phone"`

	Password string `gorm:"not null" json:"password"`

	UserInformation Passenger `gorm:"foreignKey:UserRefer;not null;unique" json:"user_information"`

	Passengers []Passenger `gorm:"foreignKey:UserRefer" json:"regular_passengers"`
}

type Passenger struct {
	gorm.Model

	Name                string    `gorm:"not null" json:"name"`
	CertificateType     uint8     `json:"certificate_type"`
	Sex                 bool      `json:"sex"`
	Birthday            time.Time `json:"birthday"`
	Country             string    `json:"country"`
	CertificateDeadline time.Time `json:"certificate_deadline"`
	Certificate         string    `json:"certificate"`
	PassengerType       uint8     `json:"passenger_type"`
	MobilePhone         string    `json:"mobile_phone"`
	CheckStatus         uint8     `json:"check_status"`
	UserStatus          uint8     `json:"user_status"`

	UserRefer uint64
	Orders    []Order `gorm:"foreignKey:PassengerRefer" json:"orders"`
}
