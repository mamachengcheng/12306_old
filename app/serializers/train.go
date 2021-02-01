package serializers


type StationList struct {
	StationName string `json:"station_name"`
	InitialName string `json:"initial_name"`
	Pinyin      string `json:"pinyin"`
	CityNo      string `json:"city_no"`
	CityName    string `json:"city_name"`
	ShowName    string `json:"show_name"`
	NameType    string `json:"name_type"`
}

type GetStation struct {
	InitialName	string `json:"initial_name" validate:"required,len=1,VerifyInitialNameFormat"`
}

type SearchStation struct {
	CityName	string `json:"city_name" validate:"required,VerifyCityNameFormat"`
}

type GetScheduleDetail struct {
	TrainNo      string    `json:"train_no"`
}