package websocket

import (
	"context"
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

type MockServer struct {
	ctx             context.Context
	config          *config.WebsocketConfig
	mockResponseMap map[string][]byte
}

func NewMock(ctx context.Context, cfg *config.WebsocketConfig, mockResponseMap map[string][]byte) *MockServer {
	return &MockServer{
		ctx:             ctx,
		config:          cfg,
		mockResponseMap: mockResponseMap,
	}
}

// Start server
func (hdr *MockServer) Start() error {
	l, err := net.Listen("tcp", hdr.config.Addr)
	if err != nil {
		return err
	}
	logger.Infof("listening on %v", hdr.config.Addr)

	s := &http.Server{
		Handler: echoServer{
			mockResponseMap: hdr.mockResponseMap,
		},
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
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

// echoServer is the WebSocket echo server implementation.
// It ensures the client speaks the echo subprotocol and
// only allows one message every 100ms with a 10 message burst.
type echoServer struct {
	mockResponseMap map[string][]byte
}

func (s echoServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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
func (s echoServer) handleMessage(ctx context.Context, c *websocket.Conn, l *rate.Limiter) error {
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

	//  default echo this request
	resp := req

	if s.mockResponseMap != nil {
		val, exist := s.mockResponseMap[string(req)]
		if exist {
			resp = val
		}
	}

	err = c.Write(ctx, mt, resp)
	if err != nil {
		return fmt.Errorf("failed to io.Copy: %w", err)
	}

	return nil
}
