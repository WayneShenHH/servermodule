package util

import (
	"errors"
	"net/http"
	"time"

	"github.com/WayneShenHH/servermodule/logger"
)

// PingServer check health
func PingServer(url string) error {
	for i := 0; i < 10; i++ {
		// Ping the server by sending a GET request to `/health`.
		logger.Infof("pingServer %v count:%d", url, i)
		var res *http.Response
		res, err := http.Get(url) //nolint
		if err != nil {
			logger.Errorf("pingServer failed, %v", err)
			// Sleep for a second to continue the next ping.
			logger.Infof("waiting for the router, retry in 1 second.")
			time.Sleep(time.Second)
			continue
		}
		defer func() {
			res.Body.Close()
		}()
		if res.StatusCode == 200 {
			logger.Infof("ping server StatusCode: %v", res.StatusCode)
			return nil
		}

		logger.Errorf("ping return code %d", res.StatusCode)
		logger.Infof("waiting for the router, retry in 1 second.")
		time.Sleep(time.Second)
	}
	return errors.New("cannot connect to the router")
}
