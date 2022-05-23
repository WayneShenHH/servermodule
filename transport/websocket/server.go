package websocket

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	"nhooyr.io/websocket"

	"github.com/WayneShenHH/servermodule/config"
	"github.com/WayneShenHH/servermodule/logger"
	"github.com/WayneShenHH/servermodule/protocol"
	"github.com/WayneShenHH/servermodule/protocol/dao"
	"github.com/WayneShenHH/servermodule/util"
)

type Server struct {
	ctx             context.Context
	config          *config.WebsocketConfig
	routeHandlerMap map[protocol.EventCode]dao.ActionHandler
	payloadParser   map[protocol.EventCode]dao.PayloadHandler
}

func NewServer(ctx context.Context, cfg *config.WebsocketConfig, routeHandlerMap map[protocol.EventCode]dao.ActionHandler, payloadParser map[protocol.EventCode]dao.PayloadHandler) *Server {
	return &Server{
		ctx:             ctx,
		config:          cfg,
		routeHandlerMap: routeHandlerMap,
		payloadParser:   payloadParser,
	}
}

// Start server
func (hdr *Server) Start() error {
	l, err := net.Listen("tcp", hdr.config.Addr)
	if err != nil {
		return err
	}
	logger.Infof("listening on %v", hdr.config.Addr)

	s := &http.Server{
		Handler: websocketServer{
			routeHandlerMap: hdr.routeHandlerMap,
			payloadParser:   hdr.payloadParser,
		},
		ReadTimeout:  hdr.config.ReadTimeout,
		WriteTimeout: hdr.config.WriteTimeout,
	}

	initManager()

	errc := make(chan error, 1)
	go func() {
		errc <- s.Serve(l)
	}()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt)
	select {
	case err := <-errc:
		logger.Errorf("failed to serve: %v", err)
	case f := <-sigs:
		logger.Infof("terminating: %v", f)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	return s.Shutdown(ctx)
}

type websocketServer struct {
	routeHandlerMap map[protocol.EventCode]dao.ActionHandler
	payloadParser   map[protocol.EventCode]dao.PayloadHandler
}

func (s websocketServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c, err := websocket.Accept(w, r, &websocket.AcceptOptions{
		OriginPatterns: []string{"*"},
	})
	if err != nil {
		logger.Errorf("%v", err)
		return
	}

	// setup new conn
	instance.register(util.GenerateGuid(), c)

	defer instance.unregister(c)
	defer c.Close(websocket.StatusInternalError, "the sky is falling")

	for {
		err = s.handleMessage(r.Context(), c)
		if websocket.CloseStatus(err) == websocket.StatusNormalClosure {
			return
		}
		if websocket.CloseStatus(err) == websocket.StatusNoStatusRcvd {
			return
		}
		if err != nil {
			logger.Errorf("failed to echo with %v: %v", r.RemoteAddr, err)
			return
		}
	}
}

// echo reads from the WebSocket connection and then writes
// the received message back to it.
// The entire function has 10s to complete.
func (s websocketServer) handleMessage(ctx context.Context, c *websocket.Conn) error {
	mt, req, err := c.Read(ctx)
	if err != nil {
		return fmt.Errorf("websocket/handleMessage/Read error: %v", err)
	}

	// process input here
	payload := dao.Payload{}

	dec, code := util.Decode(req)
	if code > 0 {
		return fmt.Errorf("websocket/handleMessage/Decode code: %v", code)
	}

	if parser, exist := s.payloadParser[dec.EventCode]; exist {
		parser(&payload)
	}

	err = util.Unmarshal(dec.Data, &payload)
	if err != nil {
		return fmt.Errorf("websocket/handleMessage/Unmarshal error: %v", err)
	}

	logger.Debug1f("websocket/handleMessage received message:\n %v", string(req))

	hdr, exist := s.routeHandlerMap[dec.EventCode]
	if !exist {
		return fmt.Errorf("EventCode: [%v] undefined", dec.EventCode)
	}

	key, exist := instance.connMap[c]
	if !exist {
		return fmt.Errorf("client connection not found")
	}

	evtcode, resp := hdr(&dao.Request{
		ClientKey: key,
		Payload:   payload,
	})

	bytes, code := util.Encode(&util.EncodeData{
		EventCode: evtcode,
		Payload:   resp,
	})

	if code > 0 {
		return fmt.Errorf("websocket/handleMessage/Encode code: %v", code)
	}

	err = c.Write(ctx, mt, bytes)
	if err != nil {
		return fmt.Errorf("websocket/handleMessage/Write: %w", err)
	}

	logger.Debug1f("websocket/handleMessage response message:\n %v", string(bytes))
	return nil
}
