package utils

//// Define push consumer interface.
//type PushConsumerInterface interface {
//	Start() error
//	Shutdown() error
//}
//
//// Implement push consumer interface.
//type PushConsumer struct {
//	F         func(msg string)
//	GroupName string
//	Topic     string
//	Address   string
//	consumer  rocketmq.PushConsumer
//}
//
//func (c *PushConsumer) Start() error {
//	var err error
//	c.consumer, err = rocketmq.NewPushConsumer(
//		consumer.WithGroupName(c.GroupName),
//		consumer.WithNameServer([]string{c.Address}),
//	)
//	if err == nil {
//		err = c.consumer.Subscribe(c.Topic, consumer.MessageSelector{},
//			func(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
//				for i := range msgs {
//					c.F(string(msgs[i].Body))
//				}
//				return consumer.ConsumeSuccess, nil
//			})
//	}
//	if err == nil {
//		err = c.consumer.Start()
//	}
//	return err
//}
//
//func (c *PushConsumer) Shutdown() error {
//	err := c.consumer.Shutdown()
//	return err
//}
//
//// Implement refund ticket push consumer interface.
//type RefundTicketPushConsumer struct {
//	PushConsumer
//}
//
//func newRefundTicketPushConsumer(address string) PushConsumerInterface {
//
//	f := func(msg string) {
//		//	TODO: Connect grpc.
//	}
//
//	return &RefundTicketPushConsumer{
//		PushConsumer{
//			F:         f,
//			GroupName: "",
//			Topic:     "",
//			Address:   address,
//		},
//	}
//}
//
//// Implement refund money push consumer interface.
//type RefundMoneyPushConsumer struct {
//	PushConsumer
//}
//
//func newRefundMoneyPushConsumer(address string) PushConsumerInterface {
//
//	f := func(msg string) {
//		//	TODO: Connect grpc.
//	}
//
//	return &RefundMoneyPushConsumer{
//		PushConsumer{
//			F:         f,
//			GroupName: "",
//			Topic:     "",
//			Address:   address,
//		},
//	}
//}
//
//func PushConsumerFactory(topic string) (rocketmq.PushConsumer, error) {
//	cfg, _ = ini.Load(static.ConfFilePath)
//	rpcCfg = cfg.Section("rpc")
//	address = rpcCfg.Key("host").String() + ":" + rpcCfg.Key("port").String()
//
//	if topic == "refund_ticket" {
//		return newRefundTicketPushConsumer(adress), nil
//	}
//	if topic == "refund_money" {
//		return newRefundMoneyPushConsumer(address), nil
//	}
//	return nil, nil
//}

//func NewProducer() iProducer {
//	cfg, _ := ini.Load(static.ConfFilePath)
//
//	rocketmqCfg := cfg.Section("rocketmq")
//	address := rocketmqCfg.Key("host").String() + ":" + rocketmqCfg.Key("port").String()
//	producer, _ := rocketmq.NewProducer(
//		producer.WithNsResovler(primitive.NewPassthroughResolver([]string{address})),
//		producer.WithRetry(3),
//	)
//
//	return &Producer{
//		Producer: producer,
//	}
//}

//func Consumer() {
//	c, _ := rocketmq.NewPushConsumer(
//		consumer.WithGroupName("testGroup"),
//		consumer.WithNameServer([]string{"47.101.141.193:9876"}),
//	)
//	err := c.Subscribe("test", consumer.MessageSelector{},
//		func(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
//			for i := range msgs {
//				fmt.Printf("subscribe callback: %v \n", msgs[i])
//			}
//
//			return consumer.ConsumeSuccess, nil
//		})
//	if err != nil {
//		fmt.Println(err.Error())
//	}
//	// Note: start after subscribe
//	err = c.Start()
//	if err != nil {
//		fmt.Println(err.Error())
//		os.Exit(-1)
//	}
//	err = c.Shutdown()
//	if err != nil {
//		fmt.Printf("shutdown Consumer error: %s", err.Error())
//	}
//}

//func Producer() {
//	p, _ := rocketmq.NewProducer(
//		producer.WithNsResovler(primitive.NewPassthroughResolver([]string{"47.101.141.193:9876"})),
//		producer.WithRetry(3),
//	)
//	err := p.Start()
//
//	if err != nil {
//		fmt.Printf("start producer error: %s", err.Error())
//		os.Exit(1)
//	}
//	topic := "test"
//
//	for i := 0; i < 10; i++ {
//		msg := &primitive.Message{
//			Topic: topic,
//			Body:  []byte("Hello RocketMQ Go Client! " + strconv.Itoa(i)),
//		}
//		res, err := p.SendSync(context.Background(), msg)
//
//		if err != nil {
//			fmt.Printf("send message error: %s\n", err)
//		} else {
//			fmt.Printf("send message success: result=%s\n", res.String())
//		}
//	}
//	err = p.Shutdown()
//	if err != nil {
//		fmt.Printf("shutdown producer error: %s", err.Error())
//	}
//}
