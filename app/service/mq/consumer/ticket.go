package consumer

import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	pb "github.com/mamachengcheng/12306/app/service/rpc/message"
	"github.com/mamachengcheng/12306/app/static"
	"google.golang.org/grpc"
	"log"
	"strconv"
	"time"
)

// Implement refund ticket push consumer interface.
//type RefundTicketPushConsumer struct {
//	PushConsumer
//}

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
//				if r, err := c.RefundTickets(ctx, &pb.RefundTicketsRequest{OrderID: orderID}); err == nil {
//					log.Printf("Greeting: %v", r.Result)
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

func StartRefundTicketConsumer() {
	c, err := rocketmq.NewPushConsumer(
		consumer.WithGroupName("CancelOrderConsumer"),
		consumer.WithNsResovler(primitive.NewPassthroughResolver([]string{static.RMQAddress})),
	)

	log.Printf("%v", static.RMQAddress)

	log.Printf("%v \n", err)

	err = c.Subscribe("CancelOrder", consumer.MessageSelector{}, func(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {

		// 退票
		conn, err := grpc.Dial(static.GrpcAddress, grpc.WithInsecure(), grpc.WithBlock())
		defer conn.Close()
		c := pb.NewTicketClient(conn)
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()

		if err == nil {
			for i := range msgs {
				orderID, err := strconv.ParseUint(string(msgs[i].Body), 10, 64)

				if err == nil {
					r, err := c.RefundTickets(ctx, &pb.RefundTicketsRequest{
						OrderID: orderID,
					})

					if err != nil || !r.Result {
						return consumer.ConsumeRetryLater, nil
					}
				}
			}

		}

		return consumer.Commit, nil
	})
	if err != nil {
		fmt.Println(err.Error())
	}
	err = c.Start()
	if err != nil {
		fmt.Println(err.Error())
	}
}
