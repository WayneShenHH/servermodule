package cmd

import (
	"github.com/spf13/cobra"

	"github.com/WayneShenHH/servermodule/config"
	"github.com/WayneShenHH/servermodule/logger"
	"github.com/WayneShenHH/servermodule/transport/http"
)

var serverHTTPCmd = &cobra.Command{
	Short: "Start HTTP Server",
	Long:  `Start HTTP Server`,
	Use:   "http:server",

	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug(cmd.Short)

		http.Start(config.Setting)

		select {}
	},
}

func init() {
	RootCmd.AddCommand(serverHTTPCmd)
}
