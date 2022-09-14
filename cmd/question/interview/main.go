package main

import (
	"time"

	"github.com/WayneShenHH/servermodule/cmd/question/interview/pkg/client"
	"github.com/WayneShenHH/servermodule/cmd/question/interview/pkg/server"
)

const orderTopic = "orderCreated"

func main() {
	server := server.NewMsgBroker()

	client1 := client.NewClient(server)
	client2 := client.NewClient(server)
	client3 := client.NewClient(server)

	// 客戶訂閱主題，收到 server 訊息後送進通道
	client1.Subscribe(orderTopic)
	client2.Subscribe(orderTopic)

	// 客戶對指定主題送出訊息
	client3.Publish(orderTopic, "Hello")

	time.Sleep(5 * time.Second)
}
