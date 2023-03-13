package cmd

import (
	"strconv"
	"time"

	"github.com/spf13/cobra"

	"github.com/WayneShenHH/servermodule/logger"
)

var tsCmd = &cobra.Command{
	Short: "parse timestamp",
	Long:  `parse timestamp`,
	Use:   "ts",

	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug(cmd.Short)
		logger.Info("args:", args)
		if len(args) < 1 {
			return
		}

		i, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			panic(err)
		}
		t := time.UnixMilli(i)

		logger.Info("date:", t.Format(time.RFC3339))
	},
}

func init() {
	RootCmd.AddCommand(tsCmd)
}
