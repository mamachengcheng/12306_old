package consumer

import (
	"context"
	pb "github.com/mamachengcheng/12306/app/service/rpc/message"
	"google.golang.org/grpc"
	"log"
	"strconv"
	"time"
)

// Implement refund ticket push consumer interface.
type RefundTicketPushConsumer struct {
	PushConsumer
}

func newRefundTicketPushConsumer(groupName, topic, address string) PushConsumerInterface {

	f := func(msg string) error {

		var err error
		if conn, err := grpc.Dial("0.0.0.0:50501", grpc.WithInsecure(), grpc.WithBlock()); err == nil {
			defer conn.Close()
			c := pb.NewTicketClient(conn)

			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()

			if orderID, err := strconv.ParseUint(msg, 10, 64); err == nil {
				if r, err := c.Refund(ctx, &pb.RefundRequest{OrderID: orderID}); err == nil {
					log.Printf("Greeting: %s", r.Code)

				}
			}
		}
		return err
	}

	return &RefundTicketPushConsumer{
		PushConsumer{
			F:         f,
			GroupName: groupName,
			Topic:     topic,
			Address:   address,
		},
	}
}
