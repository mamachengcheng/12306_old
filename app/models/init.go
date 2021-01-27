package models

import (
	"encoding/json"
	"github.com/mamachengcheng/12306/app/utils"
	"gorm.io/gorm"
	"io/ioutil"
)


type (
	Doc struct {
		Resid       string `json:"resid"`
		Name        string `json:"name"`
		InitialName string `json:"initialName"`
		Count       string `json:"count"`
		Pinyin      string `json:"pinyin"`
		Cityid      string `json:"cityid"`
		Cityname    string `json:"cityname"`
		Showname    string `json:"showname"`
		NameType    string `json:"nameType"`
		Label       string `json:"label"`
		Tag         string `json:"tag"`
	}
	Result struct {
		Dc 	[]Doc `json:"doc"`
	}
)


func InitModel() {
	utils.MysqlDB.AutoMigrate(
		&User{},
		&Passenger{},
		&Order{},
		&Train{},
		&Schedule{},
		&Stop{},
		&Seat{},
	)
	//InitStation(utils.MysqlDB)
}

func InitStation(MysqlDB *gorm.DB)  {

	var Data Result

	bytes, _ := ioutil.ReadFile("script/spider/data/station.txt")

	_ = json.Unmarshal(bytes, &Data)
	for _, val := range Data.Dc {
		MysqlDB.Create(&Station{
			StationName: val.Name,
			InitialName: val.InitialName,
			Pinyin: val.Pinyin,
			CityNo: val.Cityid,
			CityName: val.Cityname,
			ShowName: val.Showname,
			NameType: val.NameType,
		})
	}
}




