package models

import (
	"gorm.io/gorm"
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
	TrainNo      string    `gorm:"not null" json:"train_no"`
	TrainType    string    `gorm:"not null" json:"train_type"`
	TicketStatus string    `gorm:"not null" json:"ticket_status"`
	StartTime    time.Time `gorm:"not null" json:"start_time"`
	EndTime      time.Time `gorm:"not null" json:"end_time"`
	Duration     uint      `gorm:"not null" json:"duration"`

	StartStation Station `gorm:"foreignKey:StartStationRefer;not null" json:"start_station"`
	EndStation   Station `gorm:"foreignKey:EndStationRefer;not null" json:"end_station"`

	StartStationRefer uint   // Belongs to Station
	EndStationRefer   uint   // Belongs to Station
	Seats             []Seat `gorm:"foreignKey:ScheduleRefer" json:"seats"` // Has Many Seat
	Stops             []Stop `gorm:"foreignKey:ScheduleRefer" json:"stops"` // Has Many Stop
}

type Seat struct {
	gorm.Model
	SeatNo     string `gorm:"not null" json:"seat_no"`
	CarNumber  uint   `gorm:"not null" json:"car_number"`
	SeatType   uint   `gorm:"not null" json:"seat_type"`
	SeatStatus uint64 `gorm:"not null" json:"seat_status"`

	ScheduleRefer uint
}

type Stop struct {
	gorm.Model
	No uint `gorm:"not null" json:"no"`

	StartStationRefer uint    // Belongs to Station
	StartStation      Station `gorm:"foreignKey:StartStationRefer;not null" json:"start_station"`

	StartTime time.Time `gorm:"not null" json:"start_time"`
	EndTime   time.Time `gorm:"not null" json:"end_time"`
	Duration  uint      `gorm:"not null" json:"duration"`

	ScheduleRefer uint
}
