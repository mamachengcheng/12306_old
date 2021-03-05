package serializers

import "time"

type StationList struct {
	StationID   uint   `json:"station_id"`
	StationName string `json:"station_name"`
	InitialName string `json:"initial_name"`
	Pinyin      string `json:"pinyin"`
	CityNo      string `json:"city_no"`
	CityName    string `json:"city_name"`
	ShowName    string `json:"show_name"`
	NameType    string `json:"name_type"`
}

type ScheduleList struct {
	TrainNo      string    `json:"train_no"`
	TrainType    string    `json:"train_type"`
	TicketStatus string    `json:"ticket_status"`
	StartTime    time.Time `json:"start_time"`
	EndTime      time.Time `json:"end_time"`
	Duration     uint      `json:"duration"`

	StartStation StationList `json:"start_station"`
	EndStation   StationList `json:"end_station"`
}

type StopList struct {
	No          uint      `json:"no"`
	StationName string    `json:"station_name"`
	StartTime   time.Time `json:"start_time"`
	Duration    uint      `json:"duration"`
}

type GetStation struct {
	InitialName string `json:"initial_name" validate:"required,len=1,VerifyInitialNameFormat"`
}

type SearchStation struct {
	Key string `json:"key" validate:"required"`
}

type GetScheduleDetail struct {
	TrainNo   string `json:"train_no" validate:"required,VerifyTrainNoFormat"`
	StartTime string `json:"start_time" validate:"required,VerifyTimeFormat"`
}

type GetScheduleList struct {
	StartDate      string `json:"start_date" validate:"required,VerifyTimeFormat"`
	StartStationID uint   `json:"start_station_id" validate:"required"`
	EndStationID   uint   `json:"end_station_id" validate:"required"`
}

type GetStop struct {
	ScheduleID uint `json:"schedule_id" validate:"required"`
}
