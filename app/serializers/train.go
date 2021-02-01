package serializers

import "time"

type StationList struct {
	StationName string `json:"station_name"`
	InitialName string `json:"initial_name"`
	Pinyin      string `json:"pinyin"`
	CityNo      string `json:"city_no"`
	CityName    string `json:"city_name"`
	ShowName    string `json:"show_name"`
	NameType    string `json:"name_type"`
}

type ScheduleList struct {
	TrainNo			string 		`json:"train_no"`
	TrainType		string		`json:"train_type"`
	TicketStatus	string		`json:"ticket_status"`
	StartTime		time.Time	`json:"start_time"`
	EndTime			time.Time	`json:"end_time"`
	Duration		uint		`json:"duration"`
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
	CityName string `json:"city_name" validate:"required,VerifyCityNameFormat"`
}

type GetScheduleDetail struct {
	TrainNo   string `json:"train_no" validate:"required,VerifyTrainNoFormat"`
	StartTime string `json:"start_time" validate:"required,VerifyTimeFormat"`
}

type GetSchedule struct {
	StartStationName string `json:"start_station_name" validate:"required,VerifyStationNameFormat"`
	EndStationName   string `json:"end_station_name" validate:"required,VerifyStationNameFormat"`
}

type GetStop struct {
	TrainNo string `json:"train_no" validate:"required,VerifyTrainNoFormat"`
}
