package consumer

// Implement refund money push consumer interface.
type RefundMoneyPushConsumer struct {
	PushConsumer
}

func newRefundMoneyPushConsumer(groupName, topic, address string) PushConsumerInterface {

	f := func(msg string) {
		//	TODO: Connect grpc.
	}

	return &RefundMoneyPushConsumer{
		PushConsumer{
			F:         f,
			GroupName: groupName,
			Topic:     topic,
			Address:   address,
		},
	}
}
