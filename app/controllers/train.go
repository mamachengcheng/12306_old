package controllers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/mamachengcheng/12306/app/models"
	"github.com/mamachengcheng/12306/app/serializers"
	"github.com/mamachengcheng/12306/app/utils"
	"strings"
	"time"
)

func GetStationListAPI(c *gin.Context) {
	response := utils.Response{
		Code: 200,
		Msg:  "获取车站列表成功",
		Data: make(map[string]interface{}),
	}

	data := serializers.GetStation{}
	data.InitialName = strings.ToLower(data.InitialName)

	_ = c.BindJSON(&data)

	validate := serializers.GetValidate()
	err := validate.Struct(data)

	if err != nil {
		response.Code = 201
		response.Msg = "参数不合法"
	} else {
		var result []serializers.StationList
		err1 := utils.RedisDB.Get(utils.RedisDBCtx, data.InitialName).Err()
		if err1 == redis.Nil {
			var stations []models.Station
			utils.MysqlDB.Where("initial_name = ?", data.InitialName).Find(&stations)
			for _, station := range stations {
				result = append(result, serializers.StationList{
					StationID:   station.ID,
					StationName: station.StationName,
					InitialName: station.InitialName,
					Pinyin:      station.Pinyin,
					CityNo:      station.CityNo,
					CityName:    station.CityName,
					ShowName:    station.CityName,
					NameType:    station.NameType,
				})
			}
			res, _ := json.Marshal(result)
			utils.RedisDB.Set(utils.RedisDBCtx, data.InitialName, res, 0)
		} else {
			val := utils.RedisDB.Get(utils.RedisDBCtx, data.InitialName).Val()
			_ = json.Unmarshal([]byte(val), &result)
		}
		response.Data = result
	}

	utils.StatusOKResponse(response, c)

}

func SearchStationAPI(c *gin.Context) {
	response := utils.Response{
		Code: 200,
		Data: make(map[string]interface{}),
		Msg:  "获取城市相关车站列表成功",
	}
	data := serializers.SearchStation{}
	c.BindJSON(&data)
	validate := serializers.GetValidate()

	err := validate.Struct(data)
	if err != nil {
		response.Code = 201
		response.Msg = "参数不合法"
	} else {
		var stations []models.Station
		key := data.Key + "%"
		err = utils.MysqlDB.Where("station_name LIKE ? OR pinyin LIKE ?", key, key).Find(&stations).Error
		if err != nil {
			response.Code = 201
			response.Msg = "无法查询到车站"
		}
		var result []serializers.StationList
		for _, val := range stations {
			result = append(result, serializers.StationList{
				StationName: val.StationName,
				InitialName: val.InitialName,
				Pinyin:      val.Pinyin,
				CityNo:      val.CityNo,
				CityName:    val.CityName,
				ShowName:    val.CityName,
				NameType:    val.NameType,
			})
		}
		response.Data = result
	}

	utils.StatusOKResponse(response, c)

}

func GetScheduleListAPI(c *gin.Context) {
	response := utils.Response{
		Code: 200,
		Data: make(map[string]interface{}),
		Msg:  "获取车次列表成功",
	}
	data := serializers.GetScheduleList{}
	_ = c.BindJSON(&data)
	validate := serializers.GetValidate()
	err := validate.Struct(data)
	if err != nil {
		response.Code = 201
		response.Msg = "参数不合法"
	} else {
		var schedules []models.Schedule
		startTime, _ := time.ParseInLocation("2006-01-02", data.StartDate, time.Local)

		err = utils.MysqlDB.Preload("StartStation").Preload("EndStation").Where("start_time >= ? AND end_time < ?", startTime, startTime.Add(time.Hour*24)).Where("start_station_refer = ? AND end_station_refer = ?", data.StartStationID, data.EndStationID).Find(&schedules).Error

		response.Data = schedules
	}
	utils.StatusOKResponse(response, c)
}

func GetStopAPI(c *gin.Context) {
	response := utils.Response{
		Code: 200,
		Data: make(map[string]interface{}),
		Msg:  "获取列车经停站列表成功",
	}
	data := serializers.GetStop{}
	_ = c.BindJSON(&data)
	validate := serializers.GetValidate()
	err := validate.Struct(data)
	if err != nil {
		response.Code = 201
		response.Msg = "参数不合法"
	} else {
		var stops []serializers.StopList
		var schedule models.Schedule
		var train models.Train
		utils.MysqlDB.Where("id = ?", data.ScheduleID).First(&schedule)
		utils.MysqlDB.Preload("Stops").Preload("Stops.StartStation").Where("id = ?", schedule.TrainRefer).Find(&train)
		for _, stop := range train.Stops {
			stops = append(stops, serializers.StopList{
				No:          stop.No,
				StationName: stop.StartStation.StationName,
				StartTime:   stop.StartTime,
				Duration:    stop.Duration,
			})

		}
		response.Data = stops
	}

	utils.StatusOKResponse(response, c)
}
