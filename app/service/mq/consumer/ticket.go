package consumer

// Implement refund ticket push consumer interface.
//type RefundTicketPushConsumer struct {
//	PushConsumer
//}
//
//func newRefundTicketPushConsumer(groupName, topic, address string) PushConsumerInterface {
//
//	f := func(msg string) {
//
//		if conn, err := grpc.Dial(static.GrpcAddress, grpc.WithInsecure(), grpc.WithBlock()); err == nil {
//			defer conn.Close()
//			c := pb.NewTicketClient(conn)
//
//			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
//			defer cancel()
//
//			if orderID, err := strconv.ParseUint(msg, 10, 64); err == nil {
//				if r, err := c.Refund(ctx, &pb.RefundRequest{OrderID: orderID}); err == nil {
//					log.Printf("Greeting: %s", r.Code)
//
//				}
//			}
//		}
//	}
//
//	return &RefundTicketPushConsumer{
//		PushConsumer{
//			F:         f,
//			GroupName: groupName,
//			Topic:     topic,
//			Address:   address,
//		},
//	}
//}
