package controllers

import "github.com/gin-gonic/gin"

func GetStationListAPI(c *gin.Context)  {
	// TODO: 获取车站列表接口 @韦俊朗

}

func SearchStationAPI(c *gin.Context)  {
	// TODO: 搜索车站接口 @韦俊朗

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
