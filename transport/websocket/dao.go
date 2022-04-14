package websocket

import "github.com/WayneShenHH/servermodule/protocol"

type Action string

type Payload struct {
	Action Action
	Data   interface{}
}

type Request struct {
	Payload   Payload
	ClientKey string
}

type Response struct {
	Code   protocol.ErrorCode
	Action Action
	Data   interface{}
}

// ActionHandler define logic of each route
type ActionHandler func(request *Request) *Response
