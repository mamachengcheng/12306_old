package producer

import (
	"context"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"github.com/mamachengcheng/12306/app/static"
	"gopkg.in/ini.v1"
)

func SendMsg(topic, msgBody string) (string, error) {

	cfg, _ := ini.Load(static.ConfFilePath)
	rocketmqCfg := cfg.Section("rocketmq")
	address := rocketmqCfg.Key("host").String() + ":" + rocketmqCfg.Key("port").String()

	p, err := rocketmq.NewProducer(
		producer.WithNsResovler(primitive.NewPassthroughResolver([]string{address})),
		producer.WithRetry(3),
	)

	if err == nil {
		err = p.Start()
	}

	var res *primitive.SendResult

	if err == nil {
		msg := &primitive.Message{
			Topic: topic,
			Body:  []byte(msgBody),
		}
		res, err = p.SendSync(context.Background(), msg)
	}

	if err == nil {
		err = p.Shutdown()
	}

	if res == nil {
		return "", err
	}
	return res.String(), err
}
