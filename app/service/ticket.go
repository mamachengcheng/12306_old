package service

import (
	"context"
	"github.com/mamachengcheng/12306/app/models"
	pb "github.com/mamachengcheng/12306/app/service/message"
	"github.com/mamachengcheng/12306/app/utils"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	pb.UnimplementedTicketServer
}

func (s *server) Query(ctx context.Context, in *pb.QueryRequest) (*pb.QueryReply, error) {
	// TODO: 查询余票
	var schedule models.Schedule
	var seats []models.Seat
	utils.MysqlDB.Where("id = ?", in.ScheduleID).First(&schedule)

	utils.MysqlDB.Select("seat_status").Where("schedule_refer", in.ScheduleID).Find(&seats)

	return &pb.QueryReply{
		Result: nil,
		Code: 0,
	}, nil
}


func (s *server) Pay(ctx context.Context, in *pb.PayRequest) (*pb.PayReply, error) {
	// TODO:支付
	return &pb.PayReply{
		Code: 0,
	}, nil
}

func (s *server) Book(ctx context.Context, in *pb.BookRequest) (*pb.BookReply, error) {
	// TODO: 订票
	return &pb.BookReply{
		Code: 0,
	}, nil
}

func (s *server) Refund(ctx context.Context, in *pb.RefundRequest) (*pb.RefundReply, error) {
	// TODO: 退票
	return &pb.RefundReply{
		Code: 0,
	}, nil
}

const (
	host = "0.0.0.0:50051"
)

func Server() {
	listen, err := net.Listen("tcp", host)

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterTicketServer(s, &server{})

	if err := s.Serve(listen); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func init() {
	Server()
}
