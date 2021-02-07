package controllers

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/mamachengcheng/12306/app/middlewares"
	"github.com/mamachengcheng/12306/app/models"
	"github.com/mamachengcheng/12306/app/serializers"
	"github.com/mamachengcheng/12306/app/service/mq/producer"
	pb "github.com/mamachengcheng/12306/app/service/rpc/message"
	"github.com/mamachengcheng/12306/app/static"
	"github.com/mamachengcheng/12306/app/utils"
	"google.golang.org/grpc"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"strings"
	"time"
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

	if err != nil && len(data.Passengers) <= 5 && len(data.Passengers) > 0 {
		response.Code = 201
		response.Msg = "参数不合法"
	} else {
		claims := c.MustGet("claims").(*middlewares.Claims)

		user := models.User{}
		schedule := models.Schedule{}
		var seats []models.Seat

		utils.MysqlDB.Where("username = ?", claims.Username).First(&user)
		utils.MysqlDB.Where("id = ?", data.ScheduleID).First(&schedule)

		utils.MysqlDB.Where("seat_type = ? AND trainRefer = ?", data.SeatType, schedule.TrainRefer).Find(&seats)
		scheduleCode, _ := strconv.ParseUint(strings.Split(schedule.TrainNo, "_")[2], 10, 64)

		var tickets []models.Ticket
		var passenger models.Passenger

		advanceTicketCount := 0

		for _, seat := range seats {
			if advanceTicketCount == len(data.Passengers) {
				break
			}
			utils.MysqlDB.Where("id = ?", data.Passengers[advanceTicketCount].PassengerID).First(&passenger)
			if seat.SeatStatus&scheduleCode == 0 {
				tickets = append(tickets, models.Ticket{
					Seat:      seat,
					Passenger: passenger,
				})
			}
			advanceTicketCount++
		}

		if advanceTicketCount == len(data.Passengers) {
			conn, _ := grpc.Dial(static.GrpcAddress, grpc.WithInsecure(), grpc.WithBlock())
			defer conn.Close()
			c := pb.NewTicketClient(conn)
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()

			for _, ticket := range tickets {

				r, _ := c.Book(ctx, &pb.BookRequest{
					SeatID:       uint64(ticket.Seat.ID),
					ScheduleCode: scheduleCode,
					UpdatedAt:    ticket.Seat.UpdatedAt.String(),
				})
				if r.Code == 0 {
					// 退票
					producer.SendMsg("refund_ticket", "")
					break
				}
			}
		} else {
			response.Msg = "出票失败"
		}
	}

	utils.StatusOKResponse(response, c)
}

func RefundTicketAPI(c *gin.Context) {
	response := utils.Response{
		Code: 200,
		Data: make(map[string]interface{}),
		Msg:  "退票成功",
	}

	producer.SendMsg("refund_ticket", "")
	utils.StatusOKResponse(response, c)
}

func CancelOrderAPI(c *gin.Context) {

}

func PayOrderAPI(c *gin.Context) {

}
