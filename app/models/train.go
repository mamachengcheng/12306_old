package models

import (
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"github.com/mamachengcheng/12306/app/utils"
	"gorm.io/gorm"
	"strconv"
	"strings"
	"time"
)

type Station struct {
	gorm.Model
	StationName string `gorm:"not null" json:"station_name"`
	InitialName string `gorm:"not null" json:"initial_name"`
	Pinyin      string `gorm:"not null" json:"pinyin"`
	CityNo      string `gorm:"not null" json:"city_no"`
	CityName    string `gorm:"not null" json:"city_name"`
	ShowName    string `gorm:"not null" json:"show_name"`
	NameType    string `gorm:"not null" json:"name_type"`
}

type Schedule struct {
	gorm.Model
	TrainNo   string    `gorm:"not null" json:"train_no"`
	TrainType string    `gorm:"not null" json:"train_type"`
	StartTime time.Time `gorm:"not null" json:"start_time"`
	EndTime   time.Time `gorm:"not null" json:"end_time"`
	Duration  uint      `gorm:"not null" json:"duration"`

	PriceStatus string `gorm:"not null" json:"ticket_status"`

	StartStation      Station `gorm:"foreignKey:StartStationRefer;not null" json:"start_station"`
	EndStation        Station `gorm:"foreignKey:EndStationRefer;not null" json:"end_station"`
	StartStationRefer uint    // Belongs to Station
	EndStationRefer   uint    // Belongs to Station

	TrainRefer uint
}

type Seat struct {
	gorm.Model
	SeatNo     string `gorm:"not null" json:"seat_no"`
	CarNumber  uint   `gorm:"not null" json:"car_number"`
	SeatType   string `gorm:"not null" json:"seat_type"`
	SeatStatus uint64 `gorm:"not null" json:"seat_status"`

	TrainRefer uint
}

type Stop struct {
	gorm.Model
	No uint `gorm:"not null" json:"no"`

	StartTime time.Time `gorm:"not null" json:"start_time"`
	EndTime   time.Time `gorm:"not null" json:"end_time"`
	Duration  uint      `gorm:"not null" json:"duration"`

	StartStation      Station `gorm:"foreignKey:StartStationRefer;not null" json:"start_station"`
	StartStationRefer uint    // Belongs to Station

	TrainRefer uint
}

type Train struct {
	gorm.Model
	Schedules []Schedule `gorm:"foreignKey:TrainRefer" json:"schedules"` // Has Many Schedules
	Stops     []Stop     `gorm:"foreignKey:TrainRefer" json:"stops"`     // Has Many Stops
	Seats     []Seat     `gorm:"foreignKey:TrainRefer" json:"seats"`     // Has Many Seats

	TicketStatus string `gorm:"not null" json:"ticket_status"`
}

type priceStatus struct {
	Type  int     `json:"type"`
	Price float64 `json:"price"`
	Num   uint32  `json:"num"`
}

type ticketStatus struct {
	Type   int      `json:"type"`
	Num    uint32   `json:"num"`
	Status []uint32 `json:"status"`
}

func (schedule *Schedule) AfterFind(tx *gorm.DB) (err error) {

	var matrix [8][64]uint32
	var rt utils.RemainingTicket

	key := "remaining_ticket_" + strconv.Itoa(int(schedule.ID))
	val, err := utils.RedisDB.Get(utils.RedisDBCtx, key).Result()

	// 第一次查找，redis中不存在则需要生成
	if err == redis.Nil {
		var train Train
		utils.MysqlDB.Where("id = ?", schedule.TrainRefer).First(&train)

		var ticketStatus []ticketStatus
		if err := json.Unmarshal([]byte(train.TicketStatus), &ticketStatus); err == nil {
			for _, v := range ticketStatus {
				rt.Creat(v.Type, v.Num, &matrix)
			}

			res, _ := json.Marshal(matrix)
			utils.RedisDB.Set(utils.RedisDBCtx, key, res, 0)
		}

	} else {
		_ = json.Unmarshal([]byte(val), &matrix)
	}

	var priceStatus []priceStatus
	startStation, _ := strconv.ParseUint(strings.Split(schedule.TrainNo, "_")[0], 10, 64)
	endStation, _ := strconv.ParseUint(strings.Split(schedule.TrainNo, "_")[1], 10, 64)

	if err := json.Unmarshal([]byte(schedule.PriceStatus), &priceStatus); err == nil {
		for k, v := range priceStatus {
			num := rt.Find(uint32(startStation), uint32(endStation), uint32(v.Type), matrix)
			priceStatus[k].Num = num
		}
	}

	res, _ := json.Marshal(priceStatus)
	schedule.PriceStatus = string(res)

	return
}
