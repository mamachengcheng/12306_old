package producer

import (
	"context"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"github.com/mamachengcheng/12306/app/static"
	"gopkg.in/ini.v1"
)

func SendMsg(topic, msgBody string) error {

	cfg, _ := ini.Load(static.ConfFilePath)
	rocketmqCfg := cfg.Section("rocketmq")
	address := rocketmqCfg.Key("host").String() + ":" + rocketmqCfg.Key("port").String()

	p, err := rocketmq.NewProducer(
		producer.WithNsResovler(primitive.NewPassthroughResolver([]string{address})),
		producer.WithRetry(3),
	)

	if err != nil {
		return err
	}

	err = p.Start()

	_, err = p.SendSync(context.Background(), &primitive.Message{
		Topic: topic,
		Body:  []byte(msgBody),
	})

	if err != nil {
		return err
	}

	err = p.Shutdown()

	if err != nil {
		return err
	}

	return nil
}
