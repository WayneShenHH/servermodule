package client

import (
	"github.com/WayneShenHH/servermodule/cmd/question/interview/pkg/server"
)

// Handler mq interface
type IClient interface {
	// 提供 User 推送訊息到指定主題的功能
	// @param topic 指定主題
	// @param msg 訊息內容
	//
	// @return int status code
	Publish(topic string, msg string) (status int)

	// 提供 User 訂閱指定主題, 並自動開始接收該主題內的所有訊息
	// @param topic 指定主題
	//
	// @return int status code
	Subscribe(topic string) (status int)
}

// 指定一台 Message Broker, 建立一個新的 Client 實例
// @param server 後續使用的 message broker
func NewClient(server server.IMessageBroker) IClient {
	// Put your code here

	return &Client{}
}

type Client struct {
	// Put your code here
}

// 提供 User 推送訊息到指定主題的功能
// @param topic 指定主題
// @param msg 訊息內容
//
// @return int status code
func (c *Client) Publish(topic string, msg string) int {
	// Put your code here

	return 0
}

// 提供 User 訂閱指定主題, 並自動開始接收該主題內的所有訊息
// @param topic 指定主題
//
// @return int status code
func (c *Client) Subscribe(topic string) int {
	// Put your code here

	return 0
}
