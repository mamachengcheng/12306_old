package client

type RPCClientInterface interface {
	Query()
}

//var (
//	cfg, _  = ini.Load(resource.ConfFilePath)
//	rpcCfg  = cfg.Section("rpc")
//	address = rpcCfg.Key("host").String() + ":" + rpcCfg.Key("port").String()
//)
//
//func Query() uint32 {
//
//	if c, err := grpc.Dial(address, grpc.WithInsecure()); err == nil {
//		ticketClient := pb.NewTicketClient(c)
//		defer c
//		ctx, _ := context.WithTimeout(context.Background(), time.Second*3)
//		clientDeadline := time.Now().Add(3 * time.Second)
//		ctx, cancel := context.WithDeadline(ctx, clientDeadline)
//		defer cancel()
//
//		response, _ := ticketClient.Query(ctx, &pb.QueryRequest{ScheduleID: 1})
//		return response.Code
//	}
//}
//
//func Book() int64 {
//
//	if c, err := grpc.Dial(address, grpc.WithInsecure()); err == nil {
//		ticketClient := pb.NewTicketClient(c)
//		defer c
//		ctx, _ := context.WithTimeout(context.Background(), time.Second*3)
//		clientDeadline := time.Now().Add(3 * time.Second)
//		ctx, cancel := context.WithDeadline(ctx, clientDeadline)
//		defer cancel()
//
//		response, _ := ticketClient.Book(ctx, &pb.BookRequest{ScheduleID: 1, UserID: 1, SeatType: 1})
//		return response.Code
//	}
//}
//
//func Refund() int64 {
//
//	if c, err := grpc.Dial(address, grpc.WithInsecure()); err == nil {
//		ticketClient := pb.NewTicketClient(c)
//		defer c
//		ctx, _ := context.WithTimeout(context.Background(), time.Second*3)
//		clientDeadline := time.Now().Add(3 * time.Second)
//		ctx, cancel := context.WithDeadline(ctx, clientDeadline)
//		defer cancel()
//
//		response, _ := ticketClient.Refund(ctx, &pb.RefundRequest{OrderID: 1})
//		return response.Code
//	}
//}