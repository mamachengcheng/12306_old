package controllers

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/mamachengcheng/12306/app/middlewares"
	"github.com/mamachengcheng/12306/app/serializers"
	pb "github.com/mamachengcheng/12306/app/service/rpc/message"
	"github.com/mamachengcheng/12306/app/static"
	"github.com/mamachengcheng/12306/app/utils"
	"google.golang.org/grpc"
	"time"
)

//var upGrader = websocket.Upgrader{
//	CheckOrigin: func(r *http.Request) bool {
//		return true
//	},
//}

//func ReadyPayAPI(c *gin.Context) {
//	claims := c.MustGet("claims").(*middlewares.Claims)
//
//	response := utils.Response{
//		Code: 200,
//		Data: make(map[string]interface{}),
//		Msg:  "待支付",
//	}
//
//	var user models.User
//	var order models.Order
//
//	utils.MysqlDB.Where("username = ?", claims.Username).First(&user)
//	err := utils.MysqlDB.Where("user_refer = ? AND order_Status = ?", user.ID, 0).First(&order).Error
//	if errors.Is(err, gorm.ErrRecordNotFound) {
//		response.Msg = "无待支付订单"
//	} else {
//
//		ws, _ := upGrader.Upgrade(c.Writer, c.Request, nil)
//		defer ws.Close()
//		for err = utils.RedisDB.Get(utils.RedisDBCtx, claims.Username).Err(); err != nil; {
//			response.Data.(map[string]interface{})["remaining_time"] = utils.RedisDB.TTL(utils.RedisDBCtx, claims.Username)
//			_ = ws.WriteJSON(response)
//			err = utils.RedisDB.Get(utils.RedisDBCtx, claims.Username).Err()
//		}
//
//		utils.MysqlDB.Select("Tickets").Delete(&order)
//
//		response.Data = make(map[string]interface{})
//		response.Data.(map[string]interface{})["tickets"] = order.Tickets
//		response.Msg = "取消订单"
//	}
//
//	utils.StatusOKResponse(response, c)
//}

func CreateOrderAPI(c *gin.Context) {
	response := utils.Response{
		Code: 200,
		Data: make(map[string]interface{}),
		Msg:  "出票成功",
	}

	data := serializers.CreateOrder{}
	c.BindJSON(&data)

	validate := serializers.GetValidate()
	err := validate.Struct(data)

	if err != nil {
		response.Code = 201
		response.Msg = "参数不合法"
	} else {
		claims := c.MustGet("claims").(*middlewares.Claims)

		conn, err := grpc.Dial(static.GrpcAddress, grpc.WithInsecure(), grpc.WithBlock())
		if err == nil {
			defer conn.Close()
			c := pb.NewOrderClient(conn)
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()

			r, err := c.CreateOrder(ctx, &pb.CreateOrderRequest{
				Username:    claims.Username,
				ScheduleID:  data.ScheduleID,
				SeatType:    data.SeatType,
				PassengerID: data.Passengers,
			})
			if err != nil || !r.Result {
				response.Code = 202

				if r != nil {
					response.Msg = r.Msg
				}
			}
		} else {
			response.Code = 202
			response.Msg = "出票失败"
		}
	}

	utils.StatusOKResponse(response, c)
}

func CancelOrderAPI(c *gin.Context) {
	response := utils.Response{
		Code: 200,
		Data: make(map[string]interface{}),
		Msg:  "退票成功",
	}

	data := serializers.CancelOrder{}
	c.BindJSON(&data)

	validate := serializers.GetValidate()
	err := validate.Struct(data)

	if err != nil {
		response.Code = 201
		response.Msg = "参数不合法"
	} else {

		claims := c.MustGet("claims").(*middlewares.Claims)

		conn, err := grpc.Dial(static.GrpcAddress, grpc.WithInsecure(), grpc.WithBlock())


		if err == nil {
			defer conn.Close()
			c := pb.NewOrderClient(conn)
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()

			r, err := c.CancelOrder(ctx, &pb.CancelOrderRequest{
				Username: claims.Username,
				OrderID:  data.OrderID,
			})

			if err != nil || !r.Result {
				response.Code = 202
				response.Msg = r.Msg
			}
		} else {
			response.Code = 202
			response.Msg = "退票失败"
		}
	}

	utils.StatusOKResponse(response, c)
}

func PayMoneyAPI(c *gin.Context) {
	response := utils.Response{
		Code: 200,
		Data: make(map[string]interface{}),
		Msg:  "支付成功",
	}

	data := serializers.PayMoney{}
	c.BindJSON(&data)

	validate := serializers.GetValidate()
	err := validate.Struct(data)

	if err != nil {
		response.Code = 201
		response.Msg = "参数不合法"
	} else {
		claims := c.MustGet("claims").(*middlewares.Claims)

		conn, err := grpc.Dial(static.GrpcAddress, grpc.WithInsecure(), grpc.WithBlock())
		if err == nil {
			defer conn.Close()
			c := pb.NewPayClient(conn)
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()

			r, err := c.PayMoney(ctx, &pb.PayMoneyRequest{
				Username:    claims.Username,
				OrderID: data.OrderID,
			})
			if err != nil || !r.Result {
				response.Code = 202

				if r != nil {
					response.Msg = r.Msg
				}
			}
		} else {
			response.Code = 202
			response.Msg = "支付失败"
		}
	}

	utils.StatusOKResponse(response, c)
}

func RefundMoneyAPI(c *gin.Context) {
	response := utils.Response{
		Code: 200,
		Data: make(map[string]interface{}),
		Msg:  "退款成功",
	}

	data := serializers.RefundMoney{}
	c.BindJSON(&data)

	validate := serializers.GetValidate()
	err := validate.Struct(data)

	if err != nil {
		response.Code = 201
		response.Msg = "参数不合法"
	} else {
		claims := c.MustGet("claims").(*middlewares.Claims)

		conn, err := grpc.Dial(static.GrpcAddress, grpc.WithInsecure(), grpc.WithBlock())
		if err == nil {
			defer conn.Close()
			c := pb.NewPayClient(conn)
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()

			r, err := c.RefundMoney(ctx, &pb.RefundMoneyRequest{
				Username: claims.Username,
				OrderID: data.OrderID,
			})
			if err != nil || !r.Result {
				response.Code = 202

				if r != nil {
					response.Msg = r.Msg
				}
			}
		} else {
			response.Code = 202
			response.Msg = "退款失败"
		}
	}

	utils.StatusOKResponse(response, c)
}
