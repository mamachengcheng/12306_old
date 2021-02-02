package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/mamachengcheng/12306/app/models"
	"github.com/mamachengcheng/12306/app/serializers"
	"github.com/mamachengcheng/12306/app/utils"
)

func GetStationListAPI(c *gin.Context)  {
	// TODO: 获取车站列表接口 @韦俊朗
	response := utils.Response{
		Code: 200,
		Msg: "获取车站列表成功",
		Data: make(map[string]interface{}),
	}
	data := serializers.GetStation{}
	c.BindJSON(&data)

	validate := serializers.GetValidate()
	err := validate.Struct(data)

	if err != nil {
		response.Code = 201
		response.Msg = "参数不合法"
	} else {
		var stations []models.Station
		utils.MysqlDB.Where("initial_name = ?", data.InitialName).Find(&stations)
		var result []serializers.StationList
		for _, val := range stations {
			result = append(result, serializers.StationList{
				StationName: val.StationName,
				InitialName: val.InitialName,
				Pinyin: val.Pinyin,
				CityNo: val.CityNo,
				CityName: val.CityName,
				ShowName: val.CityName,
				NameType: val.NameType,
			})
		}
		response.Data = result
	}

	utils.StatusOKResponse(response, c)

}

func SearchStationAPI(c *gin.Context)  {
	// TODO: 搜索车站接口 @韦俊朗
	response := utils.Response{
		Code: 200,
		Data: make(map[string]interface{}),
		Msg: "获取城市相关车站列表成功",
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
				Pinyin: val.Pinyin,
				CityNo: val.CityNo,
				CityName: val.CityName,
				ShowName: val.CityName,
				NameType: val.NameType,
			})
		}
		response.Data = result
	}

	utils.StatusOKResponse(response, c)

}

func GetScheduleListAPI(c *gin.Context)  {
	// TODO: 获取车站列表接口 @韦俊朗

}

func GetStopAPI(c *gin.Context)  {
	// TODO: 获取车站列表接口 @韦俊朗

}

func GetScheduleDetailAPI(c *gin.Context)  {
	// TODO: 获取车次详情接口 @徐晓刚

}
