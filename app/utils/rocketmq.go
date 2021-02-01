package utils

import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"os"
	"strconv"
)

//type RocketMQ struct {
//	Topics []string
//	Producer rocketmq.Producer
//	Consumer rocketmq.PushConsumer
//}
//
//var RMQ RocketMQ
//
//func init() {
//	RMQ.Producer , _ = rocketmq.NewProducer(
//		producer.WithNsResovler(primitive.NewPassthroughResolver([]string{"47.101.141.193:9876"})),
//		producer.WithRetry(2),
//	)
//
//	RMQ.Consumer, _ = rocketmq.NewPushConsumer(
//		consumer.WithGroupName("testGroup"),
//		consumer.WithNameServer([]string{"47.101.141.193:9876"}),
//	)
//
//	err := RMQ.Producer.Start()
//	if err != nil {
//		log.Printf("start producer error: %s", err.Error())
//		os.Exit(1)
//	}
//
//	err = RMQ.Consumer.Start()
//	if err != nil {
//		log.Printf("start consumer error: %s", err.Error())
//		os.Exit(1)
//	}
//}

func Producer() {
	p, _ := rocketmq.NewProducer(
		producer.WithNsResovler(primitive.NewPassthroughResolver([]string{"47.101.141.193:9876"})),
		producer.WithRetry(2),
	)
	err := p.Start()

	if err != nil {
		fmt.Printf("start producer error: %s", err.Error())
		os.Exit(1)
	}
	topic := "test"

	for i := 0; i < 10; i++ {
		msg := &primitive.Message{
			Topic: topic,
			Body:  []byte("Hello RocketMQ Go Client! " + strconv.Itoa(i)),
		}
		res, err := p.SendSync(context.Background(), msg)

		if err != nil {
			fmt.Printf("send message error: %s\n", err)
		} else {
			fmt.Printf("send message success: result=%s\n", res.String())
		}
	}
	err = p.Shutdown()
	if err != nil {
		fmt.Printf("shutdown producer error: %s", err.Error())
	}
}

func Consumer() {
	c, _ := rocketmq.NewPushConsumer(
		consumer.WithGroupName("testGroup"),
		consumer.WithNameServer([]string{"47.101.141.193:9876"}),
	)
	err := c.Subscribe("test", consumer.MessageSelector{},
		func(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
			for i := range msgs {
				fmt.Printf("subscribe callback: %v \n", msgs[i])
			}

			return consumer.ConsumeSuccess, nil
		})
	if err != nil {
		fmt.Println(err.Error())
	}
	// Note: start after subscribe
	err = c.Start()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
	err = c.Shutdown()
	if err != nil {
		fmt.Printf("shutdown Consumer error: %s", err.Error())
	}
}
