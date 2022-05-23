package cmd

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/WayneShenHH/servermodule/config"
	"github.com/WayneShenHH/servermodule/logger"
	"github.com/WayneShenHH/servermodule/protocol"
	"github.com/WayneShenHH/servermodule/protocol/dao"
	"github.com/WayneShenHH/servermodule/protocol/eventcode"
	"github.com/WayneShenHH/servermodule/transport/websocket"
)

var serverWSCmd = &cobra.Command{
	Short: "Start WebSocket Server",
	Long:  `Start WebSocket Server`,
	Use:   "ws:server",

	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug(cmd.Short)
		handlers := map[protocol.EventCode]dao.ActionHandler{
			0: func(request *dao.Request) (protocol.EventCode, *dao.Response) {
				resp := &dao.Response{
					Payload: dao.Payload{
						TraceId: request.Payload.TraceId,
						Data:    "ok",
					},
				}

				return 0, resp
			},
			eventcode.Websocket_RegisterService_Request: func(request *dao.Request) (protocol.EventCode, *dao.Response) {
				id := request.Payload.Data.(string)
				code := websocket.GetClientManager().Register(id, request.ClientKey)

				resp := &dao.Response{
					Code: code,
					Payload: dao.Payload{
						TraceId: request.Payload.TraceId,
						Data:    id,
					},
				}

				return eventcode.Websocket_RegisterService_Response, resp
			},
		}

		parser := map[protocol.EventCode]dao.PayloadHandler{
			0: func(raw *dao.Payload) {},
			eventcode.Websocket_RegisterService_Request: func(raw *dao.Payload) {
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
