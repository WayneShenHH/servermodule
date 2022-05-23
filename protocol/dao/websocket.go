package dao

import "github.com/WayneShenHH/servermodule/protocol"

type Payload struct {
	TraceId protocol.TraceId `json:"traceId"`
	Data    interface{}      `json:"data"`
}

type Request struct {
	Payload   Payload `json:"payload"`
	ClientKey string  `json:"clientKey"`
}

type Response struct {
	Code protocol.ErrorCode `json:"code"`
	Payload
}

// ActionHandler define logic of each route
type ActionHandler func(request *Request) (protocol.EventCode, *Response)

// PayloadHandler parse data of request payload, binding specific data type
type PayloadHandler func(raw *Payload)
