package main

import (
	"encoding/json"
	"github.com/mamachengcheng/12306/app/models"
	"github.com/mamachengcheng/12306/app/static"
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
		StationNo   uint   `json:"stationNo"`
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
	//_ = utils.MysqlDB.AutoMigrate(
	//	&models.User{},
	//	&models.Passenger{},
	//	&models.Order{},
	//	&models.Ticket{},
	//	&models.Train{},
	//	&models.Station{},
	//	&models.Train{},
	//	&models.Schedule{},
	//	&models.Stop{},
	//	&models.Seat{},
	//)
	//InitStation(utils.MysqlDB)
	//InitSchedule(utils.MysqlDB)
	//InitStop(utils.MysqlDB)
	InitTrainAndScheduleAndStopAndSeat(utils.MysqlDB)
}

func InitStation(MysqlDB *gorm.DB) {

	var Data Result

	bytes, _ := ioutil.ReadFile("script/spider/data/station.txt")

	_ = json.Unmarshal(bytes, &Data)
	for _, val := range Data.Dc {
		MysqlDB.Create(&models.Station{
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

func InitSchedule(MysqlDB *gorm.DB) {

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

		var startStation, endStation models.Station
		utils.MysqlDB.Where("station_name = ?", val.Train.FmCity).First(&startStation)
		utils.MysqlDB.Where("station_name = ?", val.Train.ToCity).First(&endStation)

		res1B, _ := json.Marshal(val.Train.TicketStatusList)
		//fmt.Println(string(res1B))

		MysqlDB.Create(&models.Schedule{
			Model:          gorm.Model{},
			TrainNo:        val.Train.TrainNo,
			TrainType:      val.Train.Sort,
			ScheduleStatus: string(res1B),
			StartTime:      StartTime,
			EndTime:        EndTime,
			Duration:       val.Train.UsedTimeps,
			StartStation:   startStation,
			EndStation:     endStation,
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
			var station models.Station
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
			} else {
				duration1, _ := strconv.ParseUint(string([]rune(val1.OverTime)[:len([]rune(val1.OverTime))-2]), 10, 32)
				duration = uint(duration1)
			}
			var schedule models.Schedule
			utils.MysqlDB.Where("train_no = ?", val.Train.TrainNo).First(&schedule)
			stop := models.Stop{
				No:           val1.StationNo,
				StartStation: station,
				StartTime:    StartTime1,
				EndTime:      EndTime1,
				Duration:     duration,
			}
			utils.MysqlDB.Create(&stop)
			_ = utils.MysqlDB.Model(schedule).Association("stops").Append(stop)
		}
	}
}

type (
	TrainSeatNumStatus struct {
		TicketType int      `json:"type"`
		Num        uint     `json:"num"`
		Status     [65]uint `json:"status"`
	}

	TrainSeatPrice struct {
		TicketType     int     `json:"type"`
		Price          float32 `json:"price"`
		InitialSeatNum int     `json:"initial_seat_num"`
		LeftSeatNum    int     `json:"left_seat_num"`
	}
)

func InitTrainAndScheduleAndStopAndSeat(MysqlDB *gorm.DB) {
	var Data Result1

	bytes, _ := ioutil.ReadFile("script/spider/data/北京+上海+2021-01-31.json")

	_ = json.Unmarshal(bytes, &Data)

	date := Data.Date
	var uint_65_100 [65]uint
	for i := 0; i < 65; i++ {
		uint_65_100[i] = 100
	}
	for _, val := range Data.TrainLists {
		if val.Train.UsedTimeps == 5999 {
			continue
		}
		var a []TrainSeatNumStatus
		for _, val2 := range val.Train.TicketStatusList {
			for i, val3 := range static.SeatType {
				if val2.Cn == val3 {
					a = append(a, TrainSeatNumStatus{
						TicketType: i,
						Num:        100,
						Status:     uint_65_100,
					})
					break
				}
			}
		}
		train := models.Train{
			Model: gorm.Model{},
		}
		MysqlDB.Create(&train)

		for _, val1 := range val.Stopovers {
			var station models.Station
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
			} else {
				duration1, _ := strconv.ParseUint(string([]rune(val1.OverTime)[:len([]rune(val1.OverTime))-2]), 10, 32)
				duration = uint(duration1)
			}
			stop := models.Stop{
				No:           val1.StationNo,
				StartStation: station,
				StartTime:    StartTime1,
				EndTime:      EndTime1,
				Duration:     duration,
				TrainRefer:   train.ID,
			}
			utils.MysqlDB.Create(&stop)
			_ = utils.MysqlDB.Model(&train).Association("stops").Append(&stop)
		}

		var StartTime time.Time
		if val.Train.FmTime[0] == '-' {
			StartTime, _ = time.ParseInLocation("2006-01-02 15:04", "2099-12-31 23:59", time.Local)
		} else {
			StartTime, _ = time.ParseInLocation("2006-01-02 15:04", date+" "+val.Train.FmTime, time.Local)
		}
		EndTime, _ := time.ParseInLocation("2006-1-2 15:04:00", val.Train.ToDateTime, time.Local)

		var startStation, endStation models.Station
		utils.MysqlDB.Where("station_name = ?", val.Train.FmCity).First(&startStation)
		utils.MysqlDB.Where("station_name = ?", val.Train.ToCity).First(&endStation)

		var b []TrainSeatPrice
		for _, val2 := range val.Train.TicketStatusList {
			for i, val3 := range static.SeatType {
				if val2.Cn == val3 {
					b = append(b, TrainSeatPrice{
						TicketType:     i,
						Price:          val2.Price,
						InitialSeatNum: 10,
						LeftSeatNum:    10,
					})
					break
				}
			}
		}
		res2B, _ := json.Marshal(b)

		var stops []models.Stop
		utils.MysqlDB.Where("train_refer = ?", train.ID).Find(&stops)

		var l, r int
		for k, stop := range stops {
			if stop.StartStationRefer == startStation.ID {
				l = k
			}
			if stop.StartStationRefer == endStation.ID {
				r = k
				break
			}
		}

		var scheduleCode uint64 = 1
		for i := 0; i < r-l; i++ {
			scheduleCode = scheduleCode<<1 + 1
		}

		for i := 0; i < l; i++ {
			scheduleCode = scheduleCode << 1
		}

		schedule := models.Schedule{
			Model:          gorm.Model{},
			TrainNo:        strconv.Itoa(l) + "_" + strconv.Itoa(r) + "_" + strconv.FormatUint(scheduleCode, 10) + "_" + val.Train.TrainNo,
			TrainType:      val.Train.Sort,
			ScheduleStatus: string(res2B),
			StartTime:      StartTime,
			EndTime:        EndTime,
			Duration:       val.Train.UsedTimeps,
			StartStation:   startStation,
			EndStation:     endStation,
			TrainRefer:     train.ID,
		}
		utils.MysqlDB.Create(&schedule)
		_ = utils.MysqlDB.Model(&train).Association("schedules").Append(&schedule)

		if val.Train.TrainNo[0] == 'G' || val.Train.TrainNo[0] == 'D' {
			pre := []string{"A", "B", "C", "D", "E"}
			for i := 1; i <= 1; i++ {
				for _, val2 := range pre {
					for j := 1; j <= 2; j++ {
						seat := models.Seat{
							Model:      gorm.Model{},
							SeatNo:     strconv.Itoa(j) + val2,
							CarNumber:  uint(i),
							SeatType:   static.SecondClass,
							SeatStatus: 0,
							TrainRefer: train.ID,
						}
						utils.MysqlDB.Create(&seat)
						_ = utils.MysqlDB.Model(&train).Association("seats").Append(&seat)
					}
				}
			}
		} else {
			for i := 1; i <= 2; i++ {
				for j := 1; j <= 5; j++ {
					seat := models.Seat{
						Model:      gorm.Model{},
						SeatNo:     strconv.Itoa(j),
						CarNumber:  uint(i),
						SeatType:   static.HardSeat,
						SeatStatus: 0,
						TrainRefer: train.ID,
					}
					utils.MysqlDB.Create(&seat)
					_ = utils.MysqlDB.Model(&train).Association("seats").Append(&seat)
				}
			}
		}
	}
}

func main() {
	InitModel()
}
