package models

import (
	"encoding/json"
	"github.com/mamachengcheng/12306/app/utils"
	"gorm.io/gorm"
	"io/ioutil"
	"strconv"
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
		Dc []Doc `json:"doc"`
	}
)

type (
	TrainList struct {
		Stopovers []Stopover `json:"stopovers"`
		Train     Traintrain `json:"train"`
	}
	Stopover struct {
		StationNo   uint 	`json:"stationNo"`
		StationName string `json:"stationName"`
		Runtime     string `json:"runtime"`
		OverTime    string `json:"overTime"`
		Kilometer   string `json:"kilometer"`
		EndTime     string `json:"endTime"`
		StartTime   string `json:"startTime"`
	}
	Traintrain struct {
		TicketStatusList []TicketStatus `json:"TicketStatus"`
		TrainNo          string         `json:"TrainNo"`
		Sort             string         `json:"Sort"`
		FmTime           string         `json:"FmTime"`
		ToDateTime       string         `json:"ToDateTime"`
		FmCity           string         `json:"FmCity"`
		ToCity           string         `json:"ToCity"`
		UsedTimeps       uint           `json:"UsedTimeps"`
	}
	TicketStatus struct {
		Cn    string  `json:"Cn"`
		Price float32 `json:"Price"`
	}
	Result1 struct {
		Date       string      `json:"date"`
		To         string      `json:"to"`
		From       string      `json:"from"`
		TrainLists []TrainList `json:"train_list"`
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
	//InitSchedule(utils.MysqlDB)
	//InitStop(utils.MysqlDB)
}

func InitStation(MysqlDB *gorm.DB) {

	var Data Result

	bytes, _ := ioutil.ReadFile("script/spider/data/station.txt")

	_ = json.Unmarshal(bytes, &Data)
	for _, val := range Data.Dc {
		MysqlDB.Create(&Station{
			StationName: val.Name,
			InitialName: val.InitialName,
			Pinyin:      val.Pinyin,
			CityNo:      val.Cityid,
			CityName:    val.Cityname,
			ShowName:    val.Showname,
			NameType:    val.NameType,
		})
	}
}

// 北京南 天津南 南京南 上海
func InitScheduleTest(MysqlDB *gorm.DB) {
	var startStation, endStation Station
	var seats []Seat
	var stops []Stop
	utils.MysqlDB.Where("station_name = ?", "北京南").First(&startStation)
	utils.MysqlDB.Where("station_name = ?", "上海").First(&endStation)

	stopStations := []string{"北京南", "天津南", "南京南", "上海"}
	for idx, stopStation := range stopStations {
		var stop Station
		utils.MysqlDB.Where("station_name = ?", stopStation).First(&stop)
		stops = append(stops, Stop{
			No:           uint(idx),
			StartStation: stop,
			StartTime:    time.Now(),
			EndTime:      time.Now(),
			Duration:     1,
		})
	}

	MysqlDB.Create(&Schedule{
		TrainNo:      "G5",
		TrainType:    "",
		TicketStatus: "1",
		StartTime:    time.Now(),
		EndTime:      time.Now(),
		Duration:     1,
		StartStation: startStation,
		EndStation:   endStation,
		Seats:        seats,
		Stops:        stops,
	})
}

func InitSchedule(MysqlDB *gorm.DB) {

	var seats []Seat
	var stops []Stop

	var Data Result1

	bytes, _ := ioutil.ReadFile("script/spider/data/北京+上海+2021-01-31.json")

	_ = json.Unmarshal(bytes, &Data)

	date := Data.Date
	for _, val := range Data.TrainLists {
		var StartTime time.Time
		if val.Train.FmTime[0] == '-' {
			StartTime, _ = time.ParseInLocation("2006-01-02 15:04", "2099-12-31 23:59", time.Local)
		} else {
			StartTime, _ = time.ParseInLocation("2006-01-02 15:04", date+" "+val.Train.FmTime, time.Local)
		}
		EndTime, _ := time.ParseInLocation("2006-1-2 15:04:00", val.Train.ToDateTime, time.Local)

		var startStation, endStation Station
		utils.MysqlDB.Where("station_name = ?", val.Train.FmCity).First(&startStation)
		utils.MysqlDB.Where("station_name = ?", val.Train.ToCity).First(&endStation)

		res1B, _ := json.Marshal(val.Train.TicketStatusList)
		//fmt.Println(string(res1B))

		MysqlDB.Create(&Schedule{
			Model:             gorm.Model{},
			TrainNo:           	val.Train.TrainNo,
			TrainType:         	val.Train.Sort,
			TicketStatus:      	string(res1B),
			StartTime:         	StartTime,
			EndTime:           	EndTime,
			Duration:          	val.Train.UsedTimeps,
			StartStation: 	   	startStation,
			EndStation:   	   	endStation,
			Seats:        		seats,
			Stops:       		stops,
		})
	}
}

func InitStop(MysqlDB *gorm.DB) {
	var Data Result1

	bytes, _ := ioutil.ReadFile("script/spider/data/北京+上海+2021-01-31.json")

	_ = json.Unmarshal(bytes, &Data)

	date := Data.Date
	for _, val := range Data.TrainLists {
		for _, val1 := range val.Stopovers {
			var station Station
			utils.MysqlDB.Where("station_name = ?", val1.StationName).First(&station)
			var StartTime1 time.Time
			var EndTime1 time.Time
			if val1.StartTime[0] == '-' {
				StartTime1, _ = time.ParseInLocation("2006-01-02 15:04", "2099-12-31 23:59", time.Local)
			} else {
				StartTime1, _ = time.ParseInLocation("2006-01-02 15:04", date+" "+val1.StartTime, time.Local)
			}
			if val1.EndTime[0] == '-' {
				EndTime1, _ = time.ParseInLocation("2006-01-02 15:04", "2099-12-31 23:59", time.Local)
			} else {
				EndTime1, _ = time.ParseInLocation("2006-01-02 15:04", date+" "+val1.EndTime, time.Local)
			}
			var duration uint
			if val1.OverTime[0] == '-' {
				duration = 0
			} else{
				duration1, _ := strconv.ParseUint(string([]rune(val1.OverTime)[:len([]rune(val1.OverTime))-2]), 10, 32)
				duration = uint(duration1)
			}
			var schedule Schedule
			utils.MysqlDB.Where("train_no = ?", val.Train.TrainNo).First(&schedule)
			stop := Stop{
				No:            val1.StationNo,
				StartStation:  station,
				StartTime:     StartTime1,
				EndTime:       EndTime1,
				Duration:      duration,
				ScheduleRefer: schedule.ID,
			}
			utils.MysqlDB.Create(&stop)
			_ = utils.MysqlDB.Model(schedule).Association("stops").Append(stop)
		}
	}
}
