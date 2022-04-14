package flow

import "github.com/WayneShenHH/servermodule/protocol"

type DataFlow struct {
	TraceId    protocol.TraceId
	SequenceId protocol.SequenceId // MQ序列號
	Data       interface{}
}
