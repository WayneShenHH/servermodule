package main

type Flow struct {
	Message string
}

// Handler mq interface
type MqClient interface {
	Publish(topic string, data *Flow) int

	// 訂閱者需要傳入收訊息的通道，模組接收訊號後可以轉發給訂閱者
	Subscribe(topic string, subscriber chan *Flow) int
}

// 1.實現 mock 版本 client 模組
// 2.透過 MqServer 訂閱＆發送訊息
func NewMqClient(server MqServer) MqClient {
	return nil
}
