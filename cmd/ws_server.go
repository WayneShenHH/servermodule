package cmd

import (
	"context"
	"encoding/json"

	"github.com/spf13/cobra"

	"github.com/WayneShenHH/servermodule/config"
	"github.com/WayneShenHH/servermodule/logger"
	"github.com/WayneShenHH/servermodule/transport/websocket"
)

var serverWSCmd = &cobra.Command{
	Short: "Start WebSocket Server",
	Long:  `Start WebSocket Server`,
	Use:   "ws:server",

	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug(cmd.Short)
		handlers := map[websocket.Action]websocket.ActionHandler{
			"act1": func(request websocket.Payload) []byte {
				resp := websocket.Payload{
					Action: "act1-res",
					Data:   "ok",
				}
				bytes, err := json.Marshal(resp)
				if err != nil {
					logger.Error("act1 error:", err)
				}
				return bytes
			},
		}
		server := websocket.NewServer(context.TODO(), config.Setting.Websocket, handlers)
		logger.Fatal(server.Start())
	},
}

func init() {
	RootCmd.AddCommand(serverWSCmd)
}
