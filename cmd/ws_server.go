package cmd

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/WayneShenHH/servermodule/config"
	"github.com/WayneShenHH/servermodule/logger"
	"github.com/WayneShenHH/servermodule/protocol/dao"
	"github.com/WayneShenHH/servermodule/transport/websocket"
)

var serverWSCmd = &cobra.Command{
	Short: "Start WebSocket Server",
	Long:  `Start WebSocket Server`,
	Use:   "ws:server",

	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug(cmd.Short)
		handlers := map[dao.Action]dao.ActionHandler{
			"act1": func(request *dao.Request) *dao.Response {
				resp := &dao.Response{
					Payload: dao.Payload{
						TraceId: request.Payload.TraceId,
						Action:  "act1-res",
						Data:    "ok",
					},
				}

				return resp
			},
			"register": func(request *dao.Request) *dao.Response {
				id := request.Payload.Data.(string)
				code := websocket.GetClientManager().Register(id, request.ClientKey)

				resp := &dao.Response{
					Code: code,
					Payload: dao.Payload{
						TraceId: request.Payload.TraceId,
						Action:  "register-res",
						Data:    id,
					},
				}

				return resp
			},
		}

		parser := map[dao.Action]dao.PayloadHandler{
			"act1": func(raw *dao.Payload) {},
			"register": func(raw *dao.Payload) {
				var data string
				raw.Data = &data
			},
		}

		server := websocket.NewServer(context.TODO(), config.Setting.Websocket, handlers, parser)
		err := server.Start()
		if err != nil {
			logger.Fatalf("WebSocket Server start failed: %v", err)
		}
	},
}

func init() {
	RootCmd.AddCommand(serverWSCmd)
}
