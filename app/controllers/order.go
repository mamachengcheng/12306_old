package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/mamachengcheng/12306/app/middlewares"
	"github.com/mamachengcheng/12306/app/models"
	"github.com/mamachengcheng/12306/app/serializers"
	"github.com/mamachengcheng/12306/app/utils"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"strings"
)

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func ReadyPayAPI(c *gin.Context) {
	claims := c.MustGet("claims").(*middlewares.Claims)

	response := utils.Response{
		Code: 200,
		Data: make(map[string]interface{}),
		Msg:  "待支付",
	}

	var user models.User
	var order models.Order

	utils.MysqlDB.Where("username = ?", claims.Username).First(&user)
	err := utils.MysqlDB.Where("user_refer = ? AND order_Status = ?", user.ID, 0).First(&order).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		response.Msg = "无待支付订单"
	} else {

		ws, _ := upGrader.Upgrade(c.Writer, c.Request, nil)
		defer ws.Close()
		for err = utils.RedisDB.Get(utils.RedisDBCtx, claims.Username).Err(); err != nil; {
			response.Data.(map[string]interface{})["remaining_time"] = utils.RedisDB.TTL(utils.RedisDBCtx, claims.Username)
			_ = ws.WriteJSON(response)
			err = utils.RedisDB.Get(utils.RedisDBCtx, claims.Username).Err()
		}

		utils.MysqlDB.Select("Tickets").Delete(&order)

		response.Data = make(map[string]interface{})
		response.Data.(map[string]interface{})["tickets"] = order.Tickets
		response.Msg = "取消订单"
	}

	utils.StatusOKResponse(response, c)
}

func BookTicketsAPI(c *gin.Context) {
	response := utils.Response{
		Code: 200,
		Data: make(map[string]interface{}),
		Msg:  "出票成功",
	}

	data := serializers.BookTickets{}
	c.BindJSON(&data)

	validate := serializers.GetValidate()
	err := validate.Struct(data)

	if err != nil && len(data.Tickets) <= 5 && len(data.Tickets) > 0 {
		response.Code = 201
		response.Msg = "参数不合法"
	} else {
		claims := c.MustGet("claims").(*middlewares.Claims)

		user := models.User{}
		schedule := models.Schedule{}
		train := models.Train{}
		var seats []models.Seat

		utils.MysqlDB.Where("username = ?", claims.Username).First(&user)
		utils.MysqlDB.Where("id = ?", data.ScheduleID).First(&schedule)
		utils.MysqlDB.Where("id = ?", schedule.TrainRefer).First(&train)

		utils.MysqlDB.Where("seat_type = ?", data.SeatType).Find(&seats)
		scheduleCode, _ := strconv.ParseUint(strings.Split(schedule.TrainNo, "_")[0], 10, 64)

		var tickets []models.Ticket
		var passenger models.Passenger

		var ticketCount int64 = 0
		for _, seat := range seats {
			if ticketCount == int64(len(data.Tickets)) {
				break
			}
			utils.MysqlDB.Where("id = ?", data.Tickets[ticketCount].PassengerID).First(&passenger)
			if seat.SeatStatus&scheduleCode == 0 {
				tickets = append(tickets, models.Ticket{
					Seat:      seat,
					Passenger: passenger,
				})
			}
			ticketCount++
		}

		if ticketCount == int64(len(data.Tickets)) {
			for _, ticket := range tickets {
				result := utils.MysqlDB.Model(&ticket.Seat).Where("updated_at", ticket.Seat.UpdatedAt).Update("seat_status", ticket.Seat.SeatStatus&scheduleCode)
				ticketCount -= result.RowsAffected
			}
			if ticketCount != 0 {
				// 出票失败
			}
		} else {
			// 出票失败
		}
	}

	utils.StatusOKResponse(response, c)
}

func CancelOrderAPI(c *gin.Context) {

}

func PayOrderAPI(c *gin.Context) {

}

func RefundTicketAPI(c *gin.Context) {

}
