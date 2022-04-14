package nats_test

import (
	"os"
	"testing"
	"time"

	"github.com/WayneShenHH/servermodule/config"
	"github.com/WayneShenHH/servermodule/flow"
	"github.com/WayneShenHH/servermodule/gracefulshutdown"
	"github.com/WayneShenHH/servermodule/logger"
	"github.com/WayneShenHH/servermodule/mq/nats"
)

var cfg = &config.NatsConfig{
	Addr:              "localhost:4222,localhost:4223,localhost:4224",
	Username:          "jim",
	Password:          "password",
	ClusterID:         "test-nats-streaming-cluster",
	ReconnInterval:    time.Second,
	ConnectTimeOut:    time.Second * 5,
	StanPingsInterval: 1,
	StanPingsMaxOut:   2,
}

func Test_Subscribe(t *testing.T) {
	logger.Init("debug", "console", 0)
	signalChan := gracefulshutdown.Start()
	w := make(chan *flow.DataFlow, 10)

	topic := "topic_1"

	mqhdr := nats.New(cfg)

	mqhdr.Subscribe(topic, w)
	defer mqhdr.Unsubscribe(topic)

	go func() {
		for f := range w {
			logger.Debug("got:", string(f.Data.([]byte)))
		}
	}()

	// 開始接收MQ廣播訊息
	mqhdr.Start()

	mqhdr.Publish(topic, &flow.DataFlow{
		Data: []byte("hello"),
	})

	mqhdr.Publish(topic, &flow.DataFlow{
		Data: []byte("world"),
	})

	time.Sleep(time.Second * 5)
	signalChan <- os.Interrupt
	time.Sleep(time.Second * 1)
}
