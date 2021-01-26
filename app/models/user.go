package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model

	UserName    string `gorm:"not null;unique" json:"user_name"`

	Password string `gorm:"not null" json:"password"`

	UserInformation Passenger `gorm:"foreignKey:UserRefer;not null;unique" json:"user_information"`

	Passengers []Passenger `gorm:"foreignKey:UserRefer" json:"regular_passengers"`
}

type Passenger struct {
	gorm.Model

	Name                string    `gorm:"not null" json:"name"`
	CertificateType     uint      `json:"certificate_type"`
	Sex                 bool      `json:"sex"`
	Birthday            time.Time `json:"birthday"`
	Country             string    `json:"country"`
	CertificateDeadline time.Time `json:"certificate_deadline"`
	Certificate         string    `json:"certificate"`
	PassengerType       uint      `json:"passenger_type"`
	MobilePhone         string    `json:"mobile_phone"`
	Email               string    `gorm:"not null;unique" json:"email"`
	CheckStatus         uint      `json:"check_status"`
	UserStatus          uint      `json:"user_status"`

	UserRefer uint
	Orders    []Order `gorm:"foreignKey:PassengerRefer" json:"orders"`
}
