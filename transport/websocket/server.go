package websocket

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	"golang.org/x/time/rate"
	"nhooyr.io/websocket"

	"github.com/WayneShenHH/servermodule/config"
	"github.com/WayneShenHH/servermodule/logger"
)

type Action string

type Payload struct {
	Action Action
	Data   interface{}
}

type Server struct {
	ctx             context.Context
	config          *config.WebsocketConfig
	routeHandlerMap map[Action]ActionHandler
}

// ActionHandler define logic of each route
type ActionHandler func(request Payload) []byte

func NewServer(ctx context.Context, cfg *config.WebsocketConfig, routeHandlerMap map[Action]ActionHandler) *Server {
	return &Server{
		ctx:             ctx,
		config:          cfg,
		routeHandlerMap: routeHandlerMap,
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
		Handler:      websocketServer{},
		ReadTimeout:  hdr.config.ReadTimeout,
		WriteTimeout: hdr.config.WriteTimeout,
	}

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
	routeHandlerMap map[Action]ActionHandler
}

func (s websocketServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c, err := websocket.Accept(w, r, nil)
	if err != nil {
		logger.Errorf("%v", err)
		return
	}
	defer c.Close(websocket.StatusInternalError, "the sky is falling")

	l := rate.NewLimiter(rate.Every(time.Millisecond*100), 10)
	for {
		err = s.handleMessage(r.Context(), c, l)
		if websocket.CloseStatus(err) == websocket.StatusNormalClosure {
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
func (s websocketServer) handleMessage(ctx context.Context, c *websocket.Conn, l *rate.Limiter) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	err := l.Wait(ctx)
	if err != nil {
		return err
	}

	mt, req, err := c.Read(ctx)
	if err != nil {
		return err
	}

	// process input here
	payload := Payload{}
	err = json.Unmarshal(req, &payload)
	if err != nil {
		return err
	}

	hdr, exist := s.routeHandlerMap[payload.Action]
	if !exist {
		return fmt.Errorf("Action: %v undefined", payload.Action)
	}

	resp := hdr(payload)

	err = c.Write(ctx, mt, resp)
	if err != nil {
		return fmt.Errorf("failed to io.Copy: %w", err)
	}

	return nil
}
