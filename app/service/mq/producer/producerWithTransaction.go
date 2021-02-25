package producer

import (
	"context"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"github.com/mamachengcheng/12306/app/static"
	"gopkg.in/ini.v1"
	"sync"
)

type DemoListener struct {
	localTrans       *sync.Map
	localTransaction func() bool
}

func NewDemoListener(f func() bool) *DemoListener {
	return &DemoListener{
		localTrans:       new(sync.Map),
		localTransaction: f,
	}
}

func (dl *DemoListener) ExecuteLocalTransaction(msg *primitive.Message) primitive.LocalTransactionState {

	var state primitive.LocalTransactionState

	if dl.localTransaction() {
		state = primitive.CommitMessageState
	} else {
		state = primitive.RollbackMessageState
	}

	dl.localTrans.Store(msg.TransactionId, state)

	return state
}

func (dl *DemoListener) CheckLocalTransaction(msg *primitive.MessageExt) primitive.LocalTransactionState {
	v, existed := dl.localTrans.Load(msg.TransactionId)

	if !existed {
		return primitive.UnknowState
	}

	state := v.(primitive.LocalTransactionState)
	return state
}

func SendMsgWithTransaction(f func() bool, topic, msgBody string) error {
	cfg, _ := ini.Load(static.ConfFilePath)
	rocketmqCfg := cfg.Section("rocketmq")
	address := rocketmqCfg.Key("host").String() + ":" + rocketmqCfg.Key("port").String()

	p, _ := rocketmq.NewTransactionProducer(
		NewDemoListener(f),
		producer.WithNsResovler(primitive.NewPassthroughResolver([]string{address})),
		producer.WithRetry(1),
	)
	err := p.Start()
	if err != nil {
		return err
	}

	_, err = p.SendMessageInTransaction(context.Background(), primitive.NewMessage(topic, []byte(msgBody)))

	if err != nil {
		return err
	}

	err = p.Shutdown()

	if err != nil {
		return err
	}

	return nil
}
