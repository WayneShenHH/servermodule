package nats

import (
	"github.com/WayneShenHH/servermodule/flow"
	"github.com/WayneShenHH/servermodule/protocol"
	"github.com/WayneShenHH/servermodule/protocol/errorcode"
)

func (hdr *handler) Publish(topic string, f *flow.DataFlow) protocol.ErrorCode {
	err := hdr.connection.Publish(topic, f.Data.([]byte))
	if err != nil {
		hdr.publishBackup.Store(topic, f.Data)
	}
	return errorcode.Success
}

func (hdr *handler) Subscribe(topic string, subscriber chan *flow.DataFlow) protocol.ErrorCode {
	err := hdr.register(topic, func(data []byte, sequenceId protocol.SequenceId) {
		f := &flow.DataFlow{
			SequenceId: sequenceId,
			Data:       data,
		}
		subscriber <- f
	})
	if err != nil {
		return errorcode.Transport
	}
	return errorcode.Success

}

func (hdr *handler) SubscribeQueueGroup(topic, qgroup string, subscriber chan *flow.DataFlow) protocol.ErrorCode {
	err := hdr.registerQueueGroup(topic, qgroup, func(data []byte) {
		f := &flow.DataFlow{
			Data: data,
		}
		subscriber <- f
	})
	if err != nil {
		return errorcode.Transport
	}
	return errorcode.Success
}

func (hdr *handler) SubscribeBySequence(topic string, sequenceId protocol.SequenceId, subscriber chan *flow.DataFlow) protocol.ErrorCode {
	err := hdr.registerBySequence(topic, sequenceId, func(data []byte, sequenceId protocol.SequenceId) {
		f := &flow.DataFlow{
			SequenceId: sequenceId,
			Data:       data,
		}
		subscriber <- f
	})
	if err != nil {
		return errorcode.Transport
	}
	return errorcode.Success
}

func (hdr *handler) Unsubscribe(topic string) protocol.ErrorCode {
	hdr.unsubscribeByKey(topic)
	return errorcode.Success
}
