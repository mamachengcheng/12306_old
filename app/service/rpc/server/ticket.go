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
	var user models.User
	var schedule models.Schedule
	var train models.Ticket
	var seats []models.Seat

	utils.MysqlDB.Where("username = ?", in.UserID).First(&user)
	utils.MysqlDB.Where("id = ?", in.ScheduleID).First(&schedule)
	utils.MysqlDB.Where("id = ?", schedule.TrainRefer).First(&train)
	utils.MysqlDB.Where("seat_type = ?", in.SeatType).Find(&seats)

	scheduleCode, _ := strconv.ParseUint(strings.Split(schedule.TrainNo, "_")[0], 10, 64)

	var code int64 = -1

	for _, seat := range seats {
		if scheduleCode&seat.SeatStatus == 0 {
			result := utils.MysqlDB.Model(&seat).Where("updated_at", seat.UpdatedAt).Update("seat_status", seat.SeatStatus|scheduleCode)
			if result.RowsAffected == 0 {
				code = int64(seat.ID)
			}
			break
		}
	}

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
