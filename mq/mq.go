package mq

import (
	"github.com/WayneShenHH/servermodule/flow"
	"github.com/WayneShenHH/servermodule/protocol"
)

// Handler mq interface
type Handler interface {
	Start()
	Publish(topic string, data *flow.DataFlow) protocol.ErrorCode
	Subscribe(topic string, subscriber chan *flow.DataFlow) protocol.ErrorCode
	SubscribeQueueGroup(topic, qgroup string, subscriber chan *flow.DataFlow) protocol.ErrorCode
	SubscribeBySequence(topic string, sequence protocol.SequenceId, subscriber chan *flow.DataFlow) protocol.ErrorCode
	Unsubscribe(topic string) protocol.ErrorCode
}
