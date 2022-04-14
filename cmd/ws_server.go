package cmd

import (
	"context"

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
			"act1": func(request *websocket.Request) *websocket.Response {
				resp := &websocket.Response{
					Action: "act1-res",
					Data:   "ok",
				}

				return resp
			},
			"register": func(request *websocket.Request) *websocket.Response {
				id := request.Payload.Data.(string)
				websocket.GetClientManager().Register(id, request.ClientKey)

				resp := &websocket.Response{
					Action: "register-res",
					Data:   id,
				}

				return resp
			},
		}

		server := websocket.NewServer(context.TODO(), config.Setting.Websocket, handlers)
		logger.Fatal(server.Start())
	},
}

func init() {
	RootCmd.AddCommand(serverWSCmd)
}
