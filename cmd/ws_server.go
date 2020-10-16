package cmd

import (
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

		started := make(chan bool)
		tested := make(chan bool)
		go websocket.Start(config.Setting)

		go func() {
			<-started
			logger.Debug("deploy success")
			// tested <- true
		}()
		<-tested
	},
}

func init() {
	RootCmd.AddCommand(serverWSCmd)
}
