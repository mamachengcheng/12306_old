package models

import (
	"encoding/json"
	"github.com/mamachengcheng/12306/app/utils"
	"gorm.io/gorm"
	"io/ioutil"
	"time"
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

type (
	TrainList struct {
		Stopovers	[]Stopover	`json:"stopovers"`
		Train		Train		`json:"train"`
	}
	Stopover struct {
		StationNo   string `json:"stationNo"`
		StationName string `json:"stationName"`
		Runtime 	string `json:"runtime"`
		OverTime    string `json:"overTime"`
		Kilometer   string `json:"kilometer"`
		EndTime     string `json:"endTime"`
		StartTime   string `json:"startTime"`
	}
	Train struct {
		TicketStatusList 	[]TicketStatus	`json:"TicketStatus"`
		TrainNo				string			`json:"TrainNo"`
		Sort				string			`json:"Sort"`
		FmTime				string			`json:"FmTime"`
		ToDateTime			string			`json:"ToDateTime"`
		FmCity				string			`json:"FmCity"`
		ToCity				string			`json:"ToCity"`
		UsedTimeps			uint			`json:"UsedTimeps"`
	}
	TicketStatus struct {
		Cn			string		`json:"Cn"`
		Price		float32		`json:"Price"`
	}
	Result1 struct {
		Date       	string 		`json:"date"`
		To         	string 		`json:"to"`
		From       	string 		`json:"from"`
		TrainLists 	[]TrainList `json:"train_list"`
	}
)

func InitModel() {
	utils.MysqlDB.AutoMigrate(
		&User{},
		&Passenger{},
		&Order{},
		&Schedule{},
		&Stop{},
		&Seat{},
	)
	//InitStation(utils.MysqlDB)
	InitSchedule(utils.MysqlDB)
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

func InitSchedule(MysqlDB *gorm.DB)  {

	var Data Result1

	bytes, _ := ioutil.ReadFile("script/spider/data/北京+上海+2021-01-31.json")

	_ = json.Unmarshal(bytes, &Data)

	date := Data.Date
	for _, val := range Data.TrainLists {
		StartTime,_ := time.Parse("2006-01-02 15:04", date+" "+val.Train.FmTime)
		EndTime,_ 	:= time.Parse("2006-1-2 15:04:00", val.Train.ToDateTime)

		MysqlDB.Create(&Schedule{
			Model:             gorm.Model{},
			TrainNo:           val.Train.TrainNo,
			TrainType:         val.Train.Sort,
			TicketStatus:      "",
			StartTime:         StartTime,
			EndTime:           EndTime,
			Duration:          val.Train.UsedTimeps,
			StartStation:      val.Train.FmCity,
			EndStation:        val.Train.ToCity,
			StartStationRefer: 1,
			EndStationRefer:   2,
			Seats:             nil,
			Stops:             nil,
		})
	}
}


