package websocket

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
	Action Action
	Data   interface{}
}

// ActionHandler define logic of each route
type ActionHandler func(request *Request) *Response
