package server

import (
	"context"
	"errors"
	"github.com/mamachengcheng/12306/app/models"
	pb "github.com/mamachengcheng/12306/app/service/rpc/message"
	"github.com/mamachengcheng/12306/app/utils"
	"gorm.io/gorm"
	"log"
	"strconv"
	"strings"
)

type TicketServer struct {
	pb.UnimplementedTicketServer
}

func (s *TicketServer) BookTickets(ctx context.Context, in *pb.BookTicketsRequest) (*pb.BookTicketsReply, error) {
	res := true

	var order models.Order
	var schedule models.Schedule
	utils.MysqlDB.Where("id = ?", in.OrderID).First(&order)
	utils.MysqlDB.Where("id = ?", in.ScheduleID).First(&schedule)

	scheduleCode, _ := strconv.ParseUint(strings.Split(schedule.TrainNo, "_")[2], 10, 64)

	var seats []models.Seat
	var passengers []models.Passenger
	var tickets []models.Ticket

	// 开启事务
	err := utils.MysqlDB.Transaction(func(tx *gorm.DB) error {
		utils.MysqlDB.Where("seat_type = ? AND train_refer = ?", in.SeatType, schedule.TrainRefer).Find(&seats)
		utils.MysqlDB.Where("id = ?", in.PassengerID).Find(&passengers)

		log.Printf("seats %v \n", len(seats))
		log.Printf("passengers %v \n", len(passengers))

		// 出票
		i, j := 0, 0
		for ; i < len(seats) && j < len(passengers); {
			if seats[i].SeatStatus&scheduleCode == 0 {
				tickets = append(tickets, models.Ticket{
					Seat:      seats[i],
					Passenger: passengers[j],
				})
			}
			i++
			j++
		}

		// 余票不足
		if j != len(passengers) {
			return errors.New("InsufficientNumberOfTickets")
		} else {
			err := utils.MysqlDB.Model(&order).Association("Tickets").Append(&tickets)

			if err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		res = false
	}

	return &pb.BookTicketsReply{
		Result: res,
	}, nil
}

func (s *TicketServer) RefundTickets(ctx context.Context, in *pb.RefundTicketsRequest) (*pb.RefundTicketsReply, error) {
	res := true

	log.Printf("%v", "Hello world!")

	var order models.Order
	err := utils.MysqlDB.Transaction(func(tx *gorm.DB) error {
		utils.MysqlDB.Preload("Tickets.Seat").Preload("Tickets").Where("id = ?", in.OrderID).First(&order)
		for _, ticket := range order.Tickets {
			log.Printf("%v", ticket.Seat.SeatNo)
		}
		return nil
	})

	if err != nil {
		res = false
	}

	return &pb.RefundTicketsReply{
		Result: res,
	}, nil
}
