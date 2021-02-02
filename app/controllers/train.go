package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/mamachengcheng/12306/app/models"
	"github.com/mamachengcheng/12306/app/serializers"
	"github.com/mamachengcheng/12306/app/utils"
	"time"
)

func GetStationListAPI(c *gin.Context)  {
	// TODO: 获取车站列表接口 @韦俊朗
	response := utils.Response{
		Code: 200,
		Msg: "获取车站列表成功",
		Data: make(map[string]interface{}),
	}

	data := serializers.GetStation{}
	_ = c.BindJSON(&data)

	validate := serializers.GetValidate()
	err := validate.Struct(data)

	if err != nil {
		response.Code = 201
		response.Msg = "参数不合法"
	} else {
		var result []serializers.StationList
		err1 := utils.RedisDB.Get(utils.RedisDBCtx,data.InitialName).Err()
		if err1 == redis.Nil {
			var stations []models.Station
			utils.MysqlDB.Where("initial_name = ?", data.InitialName).Find(&stations)
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
	// TODO: 搜索车站接口 @韦俊朗
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
		err = utils.MysqlDB.Where("city_name = ?", data.CityName).Find(&stations).Error
		if err != nil {
			response.Code = 201
			response.Msg = "输入城市不存在"
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
	// TODO: 获取车站列表接口 @韦俊朗
	response := utils.Response{
		Code: 200,
		Data: make(map[string]interface{}),
		Msg:  "获取车次列表成功",
	}
	data := serializers.GetSchedule{}
	_ = c.BindJSON(&data)
	validate := serializers.GetValidate()
	err := validate.Struct(data)
	if err != nil {
		response.Code = 201
		response.Msg = "参数不合法"
	} else {
		var schedules []models.Schedule
		var result []serializers.ScheduleList
		utils.MysqlDB.Preload("StartStation", "station_name = ?", data.StartStationName).
			Preload("EndStation", "station_name = ?", data.EndStationName).Find(&schedules)

		for _, schedule := range schedules {
			if schedule.StartStation.StationName != data.StartStationName ||
				schedule.EndStation.StationName != data.EndStationName {
				continue
			}
			result = append(result, serializers.ScheduleList{
				TrainNo:      schedule.TrainNo,
				TrainType:    schedule.TrainType,
				TicketStatus: schedule.TicketStatus,
				StartTime:    schedule.StartTime,
				EndTime:      schedule.EndTime,
				Duration:     schedule.Duration,
				StartStation: serializers.StationList{
					StationName: schedule.StartStation.StationName,
					InitialName: schedule.StartStation.InitialName,
					Pinyin:      schedule.StartStation.Pinyin,
					CityNo:      schedule.StartStation.CityNo,
					CityName:    schedule.StartStation.CityName,
					ShowName:    schedule.StartStation.ShowName,
					NameType:    schedule.StartStation.NameType,
				},
				EndStation: serializers.StationList{
					StationName: schedule.EndStation.StationName,
					InitialName: schedule.EndStation.InitialName,
					Pinyin:      schedule.EndStation.Pinyin,
					CityNo:      schedule.EndStation.CityNo,
					CityName:    schedule.EndStation.CityName,
					ShowName:    schedule.EndStation.ShowName,
					NameType:    schedule.EndStation.NameType,
				},
			})
		}
		response.Data = result
	}
	utils.StatusOKResponse(response, c)
}

func GetStopAPI(c *gin.Context)  {
	// TODO: 获取车站列表接口 @韦俊朗
	response := utils.Response{
		Code: 200,
		Data: make(map[string]interface{}),
		Msg: "获取列车经停站列表成功",
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
		utils.MysqlDB.Where("train_no = ?", data.TrainNo).First(&schedule)
		utils.MysqlDB.Preload("Stops").Where("id = ?",schedule.TrainRefer).Find(&train)
		for _, stop := range train.Stops {
			stops = append(stops, serializers.StopList{
				No:				stop.No,
				StationName: 	stop.StartStation.StationName,
				StartTime: 		stop.StartTime,
				Duration:  		stop.Duration,
			})

		}
		response.Data = stops
	}

	utils.StatusOKResponse(response, c)
}

func GetScheduleDetailAPI(c *gin.Context) {
	// TODO: 获取车次详情接口 @徐晓刚
	response := utils.Response{
		Code: 200,
		Data: make(map[string]interface{}),
		Msg:  "获取车次详情接口成功",
	}
	data := serializers.GetScheduleDetail{}
	c.BindJSON(&data)
	validate := serializers.GetValidate()

	err := validate.Struct(data)
	if err != nil {
		response.Code = 201
		response.Msg = "参数不合法"
	} else {
		startTime, _ := time.ParseInLocation("2006-01-02", data.StartTime, time.Local)
		fmt.Println(startTime)
		var schedules []models.Schedule
		err = utils.MysqlDB.Where("train_no = ? and start_time >= ? and start_time < ?", data.TrainNo, startTime, startTime.Add(time.Hour*24)).Find(&schedules).Error
		if err != nil {
			response.Code = 201
			response.Msg = "输入车次不存在"
		}
		var result []serializers.ScheduleList
		for _, val := range schedules {
			var startStation models.Station
			err = utils.MysqlDB.Where("id = ?", val.StartStationRefer).Find(&startStation).Error
			var endStation models.Station
			err = utils.MysqlDB.Where("id = ?", val.EndStationRefer).Find(&endStation).Error
			result = append(result, serializers.ScheduleList{
				TrainNo:      val.TrainNo,
				TrainType:    val.TrainType,
				TicketStatus: val.TicketStatus,
				StartTime:    val.StartTime,
				EndTime:      val.EndTime,
				Duration:     val.Duration,
				StartStation: serializers.StationList{
					StationName: startStation.StationName,
					InitialName: startStation.InitialName,
					Pinyin:      startStation.Pinyin,
					CityNo:      startStation.CityNo,
					CityName:    startStation.CityName,
					ShowName:    startStation.ShowName,
					NameType:    startStation.NameType,
				},
				EndStation: serializers.StationList{
					StationName: endStation.StationName,
					InitialName: endStation.InitialName,
					Pinyin:      endStation.Pinyin,
					CityNo:      endStation.CityNo,
					CityName:    endStation.CityName,
					ShowName:    endStation.ShowName,
					NameType:    endStation.NameType,
				},
			})
		}
		response.Data = result
	}

	utils.StatusOKResponse(response, c)
}

//func TestRedis(c *gin.Context) {
//	//claims := c.MustGet("claims").(*middlewares.Claims)
//
//	//redisClient := utils.GetRedisClient()
//
//	//err := redisClient.Get(ctx, "key").Err()
//
//	//if err != redis.Nil{
//	//	redisClient.Incr(ctx, "key")
//	//} else {
//	//	redisClient.Set(ctx,"key", 0, 0)
//	//}
//	//
//	//val := redisClient.Get(ctx, "key").Val()
//	utils.DefaultResponse(resource.Success, nil, "", c)
//}
