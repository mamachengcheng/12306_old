package server

import (
	"context"
	"github.com/mamachengcheng/12306/app/models"
	pb "github.com/mamachengcheng/12306/app/service/rpc/message"
	"github.com/mamachengcheng/12306/app/utils"
	"strconv"
	"strings"
)

type TicketServer struct {
	pb.UnimplementedTicketServer
}

func (s *TicketServer) Book(ctx context.Context, in *pb.BookRequest) (*pb.BookReply, error) {
	var code int64 = 0

	var seat models.Seat

	result := utils.MysqlDB.Model(&seat).Where(" id = ? AND updated_at = ?", in.UpdatedAt).Update("seat_status", seat.SeatStatus&in.ScheduleCode)
	code = result.RowsAffected

	return &pb.BookReply{
		Code: code,
	}, nil
}

func (s *TicketServer) Refund(ctx context.Context, in *pb.RefundRequest) (*pb.RefundReply, error) {
	var order models.Order
	var code int64 = 0

	err := utils.MysqlDB.Where("id = ?", in.OrderID).First(&order).Error

	if err == nil {
		scheduleCode, _ := strconv.ParseUint(strings.Split(order.Schedule.TrainNo, "_")[0], 10, 64)

		for _, ticket := range order.Tickets {
			result := utils.MysqlDB.Model(&ticket.Seat).Update("seat_status", ticket.Seat.SeatStatus&(^scheduleCode))
			if result.Error != nil {
				code = 1
			}
		}
	} else {
		code = 2
	}

	return &pb.RefundReply{
		Code: code,
	}, nil
}
