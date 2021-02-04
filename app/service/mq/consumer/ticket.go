package consumer

// Implement refund ticket push consumer interface.
type RefundTicketPushConsumer struct {
	PushConsumer
}

func newRefundTicketPushConsumer(groupName, topic, address string) PushConsumerInterface {

	f := func(msg string) {
		//	TODO: Connect grpc.
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
