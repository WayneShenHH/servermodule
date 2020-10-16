package http

import (
	"fmt"
	"net/http"
	"time"

	"github.com/WayneShenHH/servermodule/config"
	"github.com/WayneShenHH/servermodule/logger"
	"github.com/WayneShenHH/servermodule/util"
)

// Start server
func Start(cfg config.Config) {
	if cfg.HTTP == nil {
		logger.Fatal("HTTP config is not setting")
	}

	http.HandleFunc(cfg.HTTP.PingURL, health)
	go func() {
		if err := http.ListenAndServe(cfg.HTTP.Addr, nil); err != nil {
			logger.Errorf("http.ListenAndServe faild: %v", err)
		}
	}()

	started := make(chan bool)
	go func() {
		time.Sleep(time.Second)
		if err := util.PingServer(fmt.Sprintf("http://%v%v", cfg.HTTP.Addr, cfg.HTTP.PingURL)); err != nil {
			logger.Fatalf("The router has no response, or it might took too long to start up.")
		}
		started <- true
	}()
	go func() {
		<-started
		logger.Infof("http server deployed successfully on %v.", cfg.HTTP.Addr)
	}()
}

func health(w http.ResponseWriter, req *http.Request) {
	if _, err := w.Write([]byte("ok")); err != nil {
		logger.Errorf("health Write failed: %v", err)
	}
}
