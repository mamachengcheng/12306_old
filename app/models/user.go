package models

import (
	"github.com/mamachengcheng/12306/app/utils"
	"gorm.io/gorm"
	"time"
)

func init() {
	if utils.MysqlDBErr == nil && !utils.MysqlDB.Migrator().HasTable(&User{}) {
		_ = utils.MysqlDB.Migrator().CreateTable(&User{})
	}
}

type User struct {
	gorm.Model
	UID         string
	UserName    string `gorm:"not null;unique" json:"user_name"`
	Password    string `gorm:"not null" json:"password"`
	PassengerID uint
	Passengers  []Passenger `gorm:"foreignKey:UserRefer" json:"regular_passengers"`
	Orders      []Order     `json:"orders"`

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
}

type Passenger struct {
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
	UserRefer           uint
}
