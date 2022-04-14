package websocket

import "github.com/WayneShenHH/servermodule/protocol"

type Action string

type Payload struct {
	Action Action      `json:"action"`
	Data   interface{} `json:"data"`
}

type Request struct {
	Payload   Payload `json:"payload"`
	ClientKey string  `json:"clientKey"`
}

type Response struct {
	Code   protocol.ErrorCode `json:"code"`
	Action Action             `json:"action"`
	Data   interface{}        `json:"data"`
}

// ActionHandler define logic of each route
type ActionHandler func(request *Request) *Response
