package server

import (
	"context"
	"errors"
	"github.com/mamachengcheng/12306/app/models"
	"github.com/mamachengcheng/12306/app/service/mq/producer"
	pb "github.com/mamachengcheng/12306/app/service/rpc/message"
	"github.com/mamachengcheng/12306/app/static"
	"github.com/mamachengcheng/12306/app/utils"
	"google.golang.org/grpc"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type OrderServer struct {
	pb.UnimplementedOrderServer
}

func (s *OrderServer) CreateOrder(ctx context.Context, in *pb.CreateOrderRequest) (*pb.CreateOrderReply, error) {

	// 获取Reds分布式锁
	res := utils.AcquireLockWithTimeout(in.Username)
	msg := ""

	if res {
		// 生成预订单
		var user models.User
		var schedule models.Schedule

		utils.MysqlDB.Where("username = ?", in.Username).First(&user)
		utils.MysqlDB.Where("id = ?", in.ScheduleID).First(&schedule)

		order := models.Order{
			ScheduleRefer: uint(in.ScheduleID),
			Price:         10,
		}

		err := utils.MysqlDB.Model(&user).Association("Orders").Append(&order)

		if err != nil {
			res = false
		}

		// 出票
		conn, err := grpc.Dial(static.GrpcAddress, grpc.WithInsecure(), grpc.WithBlock())
		if err == nil {
			defer conn.Close()
			c := pb.NewTicketClient(conn)
			ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
			defer cancel()

			r, err := c.BookTickets(ctx, &pb.BookTicketsRequest{
				OrderID:     uint64(order.ID),
				ScheduleID:  in.ScheduleID,
				SeatType:    in.SeatType,
				PassengerID: in.PassengerID,
			})

			if err != nil || !r.Result {
				res = false
			}
		} else {
			res = false
		}

		// 如果出票失败，则释放Reds分布式锁
		if !res {
			// 除订单
			utils.ReleaseLock(in.Username)
		}
	} else {
		msg = "尚有待支付订单"
	}

	return &pb.CreateOrderReply{
		Result: res,
		Msg: msg,
	}, nil
}

func (s *OrderServer) CancelOrder(ctx context.Context, in *pb.CancelOrderRequest) (*pb.CancelOrderReply, error) {
	// RocketMQ事务消息

	res := true

	var order models.Order
	result := utils.MysqlDB.Where("id", in.OrderID).First(&order)


	if !errors.Is(result.Error, gorm.ErrRecordNotFound) && order.OrderStatus == static.PendingOrder {

		err := producer.SendMsgWithTransaction(func() bool {
			// 软删除订单, 并发送订单号, 进行释放座位操作，并释放锁

			result := utils.MysqlDB.Model(&order).Where("id", in.OrderID).Update("order_status", static.CancelledOrder)

			if result.RowsAffected == 0 {
				return false
			}

			return true
		}, "CancelOrder", strconv.FormatUint(in.OrderID, 10))

		if err != nil {
			res = false
		}

	} else {
		res = false
	}

	if res {
		utils.ReleaseLock(in.Username)
	}

	return &pb.CancelOrderReply{
		Result: res,
	}, nil
}
