package websocket

import (
	"net/http"

	"github.com/gorilla/websocket"

	"github.com/WayneShenHH/servermodule/config"
	"github.com/WayneShenHH/servermodule/logger"
)

// Start server
func Start(cfg config.Config) {
	upgrader := &websocket.Upgrader{
		//如果有 cross domain 的需求，可加入這個，不檢查 cross domain
		CheckOrigin: func(r *http.Request) bool { return true },
	}
	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		c, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			logger.Infof("upgrade: %v", err)
			return
		}
		defer func() {
			logger.Infof("disconnect !!")
			c.Close()
		}()
		for {
			mtype, msg, err := c.ReadMessage()
			if err != nil {
				logger.Errorf("read: %v", err)
				break
			}
			logger.Infof("receive: %s", msg)
			err = c.WriteMessage(mtype, msg)
			if err != nil {
				logger.Errorf("write: %v", err)
				break
			}
		}
	})
	logger.Infof("server start at :8899")
	logger.Fatal(http.ListenAndServe(":8899", nil))
}
