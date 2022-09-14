package server

import (
	"github.com/WayneShenHH/servermodule/cmd/question/interview/pkg/flow"
)

type IMessageBroker interface {
	// 提供 Client 推送訊息到指定主題
	// @param topic 指定主題
	// @param data 訊息封包
	//
	// @return int status code
	Publish(topic string, data *flow.Flow) (status int)

	// 提供 client 訂閱指定主題
	// @param topic 指定主題
	//
	// @return chan *Flow 指定主題的消息隊列
	Subscribe(topic string) (subscriber chan *flow.Flow)
}

func NewMsgBroker() IMessageBroker {
	// Put your code here

	return &MsgBroker{}
}

type MsgBroker struct {
	// Put your code here
}

// 提供 Client 推送訊息到指定主題
// @param topic 主題
// @param data 訊息封包
//
// @return int status code
func (s *MsgBroker) Publish(topic string, data *flow.Flow) (status int) {
	// Put your code here

	return 0
}

// 提供 client 訂閱指定主題
// @param topic 指定主題
//
// @return chan *Flow 指定主題的消息隊列
func (s *MsgBroker) Subscribe(topic string) chan *flow.Flow {
	// Put your code here

	return nil
}
