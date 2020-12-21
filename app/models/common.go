package models

import "github.com/mamachengcheng/12306/app/utils"

func InitModel() {
	utils.MysqlDB.AutoMigrate(
		&User{},
		&Passenger{},
		&Order{},
		&Train{},
		&Seat{},
		&Stop{},
		&Station{},
		&Schedule{},
	)
}
