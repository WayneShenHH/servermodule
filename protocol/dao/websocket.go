package dao

import "github.com/WayneShenHH/servermodule/protocol"

type Action string

type Payload struct {
	TraceId protocol.TraceId `json:"traceId"`
	Action  Action           `json:"action"`
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
type ActionHandler func(request *Request) *Response

type PayloadHandler func(raw *Payload)
