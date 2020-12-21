package models

import (
	"gorm.io/gorm"
	"time"
)

type Train struct {
	gorm.Model
	TrainNo   string `gorm:"not null" json:"train_no"`
	TrainType uint `gorm:"not null" json:"train_type"`
	Seats     []Seat `gorm:"not null" json:"seats"`
}

type Seat struct {
	SeatNo     string  `gorm:"not null" json:"seat_no"`
	CarNumber  uint    `gorm:"not null" json:"car_number"`
	Train      Train   `gorm:"not null" json:"train"`
	Price      float32 `gorm:"not null" json:"price"`
	SeatType   uint    `gorm:"not null" json:"seat_type"`
	SeatStatus uint    `gorm:"not null" json:"seat_status"`
}

type Station struct {
	gorm.Model
	StationName string `gorm:"not null" json:"station_name"`
	InitialName string `gorm:"not null;unique" json:"initial_name"`
	Pinyin      string `gorm:"not null" json:"pinyin"`
	CityNo      string `gorm:"not null" json:"city_no"`
	CityName    string `gorm:"not null" json:"city_name"`
	ShowName    string `gorm:"not null" json:"show_name"`
	NameType    string `gorm:"not null" json:"name_type"`
}

type Schedule struct {
	gorm.Model
	Train        Train     `gorm:"not null" json:"train"`
	StartStation Station   `gorm:"not null" json:"start_station"`
	EndStation   Station   `gorm:"not null" json:"end_station"`
	StartTime    time.Time `gorm:"not null" json:"start_time"`
	EndTime      time.Time `gorm:"not null" json:"end_time"`
	Duration     uint      `gorm:"not null" json:"duration"`
	Stops        []Stop    `gorm:"not null" json:"stops"`
}

type Stop struct {
	gorm.Model
	No           uint    `gorm:"not null" json:"no"`
	StartStation Station `gorm:"not null" json:"start_station"`
	EndStation   Station `gorm:"not null" json:"end_station"`
	StartTime    Station `gorm:"not null" json:"start_time"`
	EndTime      Station `gorm:"not null" json:"end_time"`
	Duration     uint    `gorm:"not null" json:"duration"`
}
