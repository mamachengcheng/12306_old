package consumer

func PushConsumerFactory(groupName, topic, address string) (PushConsumerInterface, error) {

	if topic == "refund_ticket" {
		return newRefundTicketPushConsumer(groupName, topic, address), nil
	}

	if topic == "refund_money" {
		return newRefundMoneyPushConsumer(groupName, topic, address), nil
	}

	return nil, nil
}
