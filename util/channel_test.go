package util

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/WayneShenHH/servermodule/logger"
)

type message struct {
	Text string
}

func Test_Select(t *testing.T) {
	w := make(chan *message, 500)
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		for {
			select {
			case got := <-w:
				logger.Debug(got, len(w))
			case <-ctx.Done():
				logger.Debug("ctx.Done()")
			}
		}
	}()
	go mockWrite(w)

	time.Sleep(time.Second * 1)
	close(w)
	cancel()
	logger.Debug("close end")
	time.Sleep(time.Second * 1)
}

func Test_PrioritySelect(t *testing.T) {
	w := make(chan *message, 500)
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		defer logger.Debug("read end")
		for {
			select {
			case got := <-w:
				select {
				case <-ctx.Done():
					logger.Debug("ctx.Done() inside")
					return
				default:
					logger.Debug(got, len(w))
				}
			case <-ctx.Done():
				logger.Debug("ctx.Done() outside")
				return
			}
		}
	}()
	go mockWrite(w)

	time.Sleep(time.Second * 1)
	close(w)
	cancel()
	logger.Debug("cancel")
	time.Sleep(time.Second * 1)
}

func Test_Range(t *testing.T) {
	w := make(chan *message, 500)
	go func() {
		for got := range w {
			logger.Debug(got, len(w))
		}
		logger.Debug("read end")
	}()
	go mockWrite(w)

	time.Sleep(time.Second * 1)
	close(w)
	logger.Debug("close end")
	time.Sleep(time.Second * 1)
}

func Test_SelectTimeout(t *testing.T) {
	w := make(chan *message, 500)
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		defer logger.Debug("read end")

		for {
			select {
			case got := <-w:
				logger.Debugf("remain: %v, got: %v", len(w), got)
				select {
				case <-ctx.Done():
					logger.Debug("ctx.Done() inside")
					rangeChannel(w)
					return
				default:
				}
				time.Sleep(time.Second * 1)
			case <-ctx.Done():
				logger.Debug("ctx.Done() outside")
				rangeChannel(w)
				return
			}
		}
	}()
	go mockWrite(w)

	time.Sleep(time.Second * 1)
	cancel()
	logger.Debug("cancel")
	time.Sleep(time.Second * 2)
}
func mockWrite(w chan *message) {
	func() {
		for i := 1; i <= 10; i++ {
			w <- &message{
				Text: fmt.Sprintf("hello %v", i),
			}
		}
		logger.Debug("write 10 messages end")
	}()
}
func rangeChannel(w chan *message) {
	length := len(w)
	for i := 0; i < length; i++ {
		got := <-w
		logger.Debugf("index: %v, range %v times, remain: %v, got: %v", i, length, len(w), got)
	}
}
