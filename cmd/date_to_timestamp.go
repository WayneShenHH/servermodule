package cmd

import (
	"fmt"
	"strings"
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

		sp := strings.Split(args[0], " ")
		if len(sp) < 2 {
			args[0] = fmt.Sprintf("%v 00:00:00", args[0])
		}

		d, err := time.Parse(time.DateTime, args[0])
		if err != nil {
			panic(err)
		}

		logger.Infof("date: %v, ts: %v", d.Format(time.DateTime), d.UnixMilli())
	},
}

func init() {
	RootCmd.AddCommand(dateToTimestampCmd)
}
