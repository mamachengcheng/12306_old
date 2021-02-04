package consumer

import (
	"context"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
)

// Define push consumer interface.
type PushConsumerInterface interface {
	Start() error
	Shutdown() error
}

// Implement push consumer.
type PushConsumer struct {
	F         func(msg string)
	GroupName string
	Topic     string
	Address   string
	consumer  rocketmq.PushConsumer
}

func (c *PushConsumer) Start() error {
	var err error
	c.consumer, err = rocketmq.NewPushConsumer(
		consumer.WithGroupName(c.GroupName),
		consumer.WithNameServer([]string{c.Address}),
	)
	if err == nil {
		err = c.consumer.Subscribe(c.Topic, consumer.MessageSelector{},
			func(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
				for i := range msgs {
					c.F(string(msgs[i].Body))
				}
				return consumer.ConsumeSuccess, nil
			})
	}
	if err == nil {
		err = c.consumer.Start()
	}
	return err
}

func (c *PushConsumer) Shutdown() error {
	err := c.consumer.Shutdown()
	return err
}

