package main

import "fmt"

const orderTopic = "orderCreated"

func main() {
	fmt.Println("start")
	srv := NewMqServer()

	msgCh1 := make(chan *Flow, 10)
	msgCh2 := make(chan *Flow, 10)

	cli1 := NewMqClient(srv)
	cli2 := NewMqClient(srv)
	cli3 := NewMqClient(srv)

	// 客戶訂閱主題，收到 server 訊息後送進通道
	cli1.Subscribe(orderTopic, msgCh1)
	cli2.Subscribe(orderTopic, msgCh2)

	// 客戶對指定主題送出訊息
	cli3.Publish(orderTopic, &Flow{
		Message: "Hello",
	})
}
