package main

type MqServer interface {
	// 提供 client 推訊息到指定主題
	Publish(topic string, data *Flow) int

	// 提供 client 訂閱指定主題
	Subscribe(topic string, subscriber chan *Flow) int
}

// 實現 mock 版本 mq server
// 1.可以提供多個 client 註冊
// 2.依照主題廣播訊息給 client
func NewMqServer() MqServer {
	return nil
}
