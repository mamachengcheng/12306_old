package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/mamachengcheng/12306/app/middlewares"
	"github.com/mamachengcheng/12306/app/models"
	"github.com/mamachengcheng/12306/app/serializers"
	"github.com/mamachengcheng/12306/app/utils"
	"net/http"
	"time"
)

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func ReadyPayAPI(c *gin.Context)  {
	claims := c.MustGet("claims").(*middlewares.Claims)

	response := utils.Response{
		Code: 200,
		Data: make(map[string]interface{}),
		Msg:  "待支付",
	}

	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)

	if err == nil {
		defer ws.Close()
		for err = utils.RedisDB.Get(utils.RedisDBCtx, claims.Username).Err(); err != nil; {
			response.Data.(map[string]interface{})["remaining_time"] = utils.RedisDB.TTL(utils.RedisDBCtx, claims.Username)
			_ = ws.WriteJSON(response)
			err = utils.RedisDB.Get(utils.RedisDBCtx, claims.Username).Err()
		}

		var user models.User{}
		var order models.Order{}

		utils.MysqlDB.Where("username = ?", claims.Username).First(&user)
		utils.MysqlDB.Where("user_refer = ? AND order_Status = ?", user.ID, 0).First(&order)
		utils.MysqlDB.Select("Tickets").Delete(&order)

		response.Data = make(map[string]interface{})
		response.Data.(map[string]interface{})["tickets"] = order.Tickets
		response.Msg = "取消订单"

		utils.StatusOKResponse(response, c)
	}
}

func BookTicketAPI(c *gin.Context) {
	response := utils.Response{
		Code: 200,
		Data: make(map[string]interface{}),
		Msg:  "",
	}

	data := serializers.BookTicket{}
	c.BindJSON(&data)
	validate := serializers.GetValidate()
	err := validate.Struct(data)

	if err != nil {
		response.Code = 201
		response.Msg = "参数不合法"

		utils.StatusOKResponse(response, c)
	} else {
		claims := c.MustGet("claims").(*middlewares.Claims)

		user := models.User{}
		schedule := models.Schedule{}
		passenger := models.Passenger{}
		train := models.Train{}

		utils.MysqlDB.Where("username = ?", claims.Username).First(&user)
		utils.MysqlDB.Where("id = ?", data.ScheduleID).First(&schedule)
		utils.MysqlDB.Where("id = ?", schedule.TrainRefer).First(&train)
		utils.MysqlDB.Where("id = ?", data.PassengerID).First(&passenger)

		// TODO: 车次表生成时添入 scheduleCode  TrainNo: scheduleCode_scheduleNo
		var l, r int
		for k, stop := range train.Stops {
			if stop.TrainRefer == schedule.StartStationRefer {
				l = k - 1
			}
			if stop.TrainRefer == schedule.EndStationRefer {
				r = k - 1
				break
			}
		}

		var scheduleCode uint64 = 1
		for i := 0; i < r - l + 1; i++ {
			scheduleCode = scheduleCode<<1 + 1
		}

		for i := 0; i < l; i++ {
			scheduleCode = scheduleCode << 1
		}

		for _, seat := range train.Seats {
			if seat.SeatStatus & scheduleCode == 1 {

				utils.MysqlDB.Model(&passenger).Association("Orders").Append(&models.Order{
					Seat:     seat,
					Schedule: schedule,
				})
				break
			}
		}

		if err := utils.RedisDB.Get(utils.RedisDBCtx, claims.Username).Err(); err == nil {
			utils.RedisDB.Set(utils.RedisDBCtx, claims.Username, true, time.Minute*30)
		}

		ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)

		if err == nil {
			defer ws.Close()
			for {
				response.Data.(map[string]interface{})["remaining_time"] = utils.RedisDB.TTL(utils.RedisDBCtx, claims.Username)
				_ = ws.WriteJSON(response)
			}
		}
	}
}

func RefundTicketAPI(c *gin.Context) {

}

func PayOrderAPI(c *gin.Context) {

}

func RefundMoneyAPI(c *gin.Context) {

}
