package websocket

import (
	"context"

	"github.com/WayneShenHH/servermodule/config"
	"github.com/WayneShenHH/servermodule/flow"
	"nhooyr.io/websocket"
)

type Client struct {
	ctx                 context.Context
	conn                *websocket.Conn
	config              *config.WebsocketConfig
	writeFlow           chan *flow.DataFlow
	status              uint8 // 標記此模組 運作中，關閉中，結束
	reader              chan []byte
	remainResponseCount int // 紀錄剩下多少個對 server 的請求還沒有回應，未來要改為 map 搭配 traceId
}

// New 產生 websocket 實體，並管理連線
// @param ctx 執行緒追蹤信號
// @param config 配置設定，須注意連線資訊是 Url (包含protocal, ex: "ws://" or "wss://")
// @param writeFlow 預先建立一個通道讓 Handler 可以收到要傳送出去的訊息
// @return *Handler
func NewClient(ctx context.Context, config *config.WebsocketConfig, writeFlow chan *flow.DataFlow) *Client {
	// c := connect(config)

	hdr := &Client{
		ctx:    ctx,
		config: config,
		// conn:      c,
		writeFlow: writeFlow,
		reader:    make(chan []byte),
	}

	return hdr
}

// TODO: 連線，重連，關閉，註冊，PING