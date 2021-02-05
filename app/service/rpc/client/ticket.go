package client

import (
	"context"
	pb "github.com/mamachengcheng/12306/app/service/rpc/message"
	"google.golang.org/grpc"
)

type QueryExecutor struct {
}

func (e *QueryExecutor) exec(coon *grpc.ClientConn, ctx context.Context, request map[string]interface{}) (map[string]interface{}, error) {
	var err error
	var data map[string]interface{}
	data = make(map[string]interface{})

	c := pb.NewTicketClient(coon)
	if r, err := c.Query(ctx, &pb.QueryRequest{ScheduleID: request["ScheduleID"].(uint64)}); err == nil {

		data["code"] = r.Code
		data["result"] = r.Result
	}
	return data, err
}

type BookExecutor struct {
}

func (e *BookExecutor) exec(coon *grpc.ClientConn, ctx context.Context, request map[string]interface{}) (map[string]interface{}, error) {

	var err error
	var data map[string]interface{}
	data = make(map[string]interface{})

	c := pb.NewTicketClient(coon)
	if r, err := c.Book(ctx, &pb.BookRequest{ScheduleID: request["ScheduleID"].(uint64)}); err == nil {
		data["code"] = r.Code
	}
	return data, err
}

type RefundExecutor struct {
}

func (e *RefundExecutor) exec(coon *grpc.ClientConn, ctx context.Context, request map[string]interface{}) (map[string]interface{}, error) {

	var err error
	var data map[string]interface{}
	data = make(map[string]interface{})

	c := pb.NewTicketClient(coon)
	if r, err := c.Refund(ctx, &pb.RefundRequest{OrderID: request["ScheduleID"].(uint64)}); err == nil {
		data["code"] = r.Code
	}
	return data, err
}
