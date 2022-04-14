package websocket

import (
	"context"

	"github.com/WayneShenHH/servermodule/errorcode"
	"github.com/WayneShenHH/servermodule/logger"
	"github.com/WayneShenHH/servermodule/protocol"
	"nhooyr.io/websocket"
)

var instance *clientManager

// ClientManager 管理所有連線的客戶
//
// server 啟動時同時生成管理員
//
// 1.主動通知訊息
//
// 2.廣播功能
//
// 3.替換連線的識別碼 ex: playerId
type ClientManager interface {
	Send(key string, data []byte) protocol.ErrorCode
	Broadcast(data []byte) protocol.ErrorCode
	Register(newKey, oldKey string) protocol.ErrorCode
}

type clientManager struct {
	ctx       context.Context
	clientMap map[string]*websocket.Conn
	connMap   map[*websocket.Conn]string
}

func initManager() {
	if instance == nil {
		instance = &clientManager{
			clientMap: make(map[string]*websocket.Conn),
			connMap:   make(map[*websocket.Conn]string),
		}
	}
}

func GetClientManager() ClientManager {
	return instance
}

func (c *clientManager) Register(newKey, oldKey string) protocol.ErrorCode {
	conn, exist := c.clientMap[oldKey]
	if exist && oldKey != newKey {
		delete(c.clientMap, oldKey)
	}
	c.register(newKey, conn)

	return 0
}

func (c *clientManager) Send(key string, data []byte) protocol.ErrorCode {
	err := c.clientMap[key].Write(c.ctx, websocket.MessageText, data)
	if err != nil {
		logger.Errorf("websocket/Send: key=[%v], err: %v", key, err)
		return errorcode.Transport
	}
	return 0
}

func (c *clientManager) Broadcast(data []byte) protocol.ErrorCode {
	for k := range c.clientMap {
		code := c.Send(k, data)
		if code > 0 {
			return code
		}
	}
	return 0
}

// --------------------------------------------------------------------------------------

func (c *clientManager) register(key string, conn *websocket.Conn) {
	c.clientMap[key] = conn
	c.connMap[conn] = key
	logger.Debugf("websocket/register: key=[%v]", key)
}

func (c *clientManager) unregister(conn *websocket.Conn) protocol.ErrorCode {
	if key, exist := c.connMap[conn]; exist {
		delete(c.connMap, conn)
		delete(c.clientMap, key)
		logger.Debugf("websocket/unregister: key=[%v]", key)
	}
	return 0
}
