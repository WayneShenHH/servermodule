package cmd

import (
	"time"

	"github.com/spf13/cobra"

	"github.com/WayneShenHH/servermodule/logger"
)

var dateToTimestampCmd = &cobra.Command{
	Short: "transfer date to timestamp",
	Long:  `transfer date to timestamp`,
	Use:   "date",

	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug(cmd.Short)
		logger.Info("args:", args)
		if len(args) < 1 {
			return
		}

		d, err := time.Parse("2006-01-02", args[0])
		if err != nil {
			panic(err)
		}

		logger.Infof("date: %v, ts: %v", d.Format(time.RFC3339), d.UnixMilli())
	},
}

func init() {
	RootCmd.AddCommand(dateToTimestampCmd)
}
