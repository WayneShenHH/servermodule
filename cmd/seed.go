package cmd

import (
	"github.com/spf13/cobra"

	"github.com/WayneShenHH/servermodule/config"
	"github.com/WayneShenHH/servermodule/db/arangodb"
	"github.com/WayneShenHH/servermodule/gracefulshutdown"
	"github.com/WayneShenHH/servermodule/logger"
)

var seedCmd = &cobra.Command{
	Short: "run arangodb seed",
	Long:  `run arangodb seed`,
	Use:   "db:seed",

	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug(cmd.Short)

		gracefulshutdown.Start()
		db := arangodb.NewArango(config.Setting.Database)
		code := db.Seed()
		if code > 0 {
			logger.Fatalf("seed failed: %v", code)
		}

		logger.Info("db seed done")
	},
}

func init() {
	RootCmd.AddCommand(seedCmd)
}
