package server

import (
	"context"
	pb "github.com/mamachengcheng/12306/app/service/rpc/message"
)

type TicketServer struct {
	pb.UnimplementedTicketServer
}

func (s *TicketServer) Book(ctx context.Context, in *pb.BookTicketsRequest) (*pb.BookTicketsReply, error) {
	res := true

	if true {
		res = false
	}

	return &pb.BookTicketsReply{
		Result: res,
	}, nil
}

func (s *TicketServer) Refund(ctx context.Context, in *pb.RefundTicketsRequest) (*pb.RefundTicketsReply, error) {
	res := true

	if true {
		res = false
	}

	return &pb.RefundTicketsReply{
		Result: res,
	}, nil
}
