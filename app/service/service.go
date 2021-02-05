package service

import (
	"github.com/mamachengcheng/12306/app/service/mq/consumer"
	pb "github.com/mamachengcheng/12306/app/service/rpc/message"
	"github.com/mamachengcheng/12306/app/service/rpc/server"
	"github.com/mamachengcheng/12306/app/static"
	"google.golang.org/grpc"
	"log"
	"net"
)

func startGrpcServer() {
	listen, err := net.Listen("tcp", static.GrpcAddress)

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterTicketServer(s, &server.TicketServer{})

	log.Println("Start grpc server!")

	if err := s.Serve(listen); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func startMQConsumer() {

	if refundTicket, err := consumer.PushConsumerFactory("refund_ticket_consumer", "refund_ticket", static.RMQAddress); err == nil {
		refundTicket.Start()
	}

	if refundMoney, err := consumer.PushConsumerFactory("refund_money_consumer", "refund_money", static.RMQAddress); err == nil {
		refundMoney.Start()
	}
}

func Start() {
	startMQConsumer()
	startGrpcServer()
}
