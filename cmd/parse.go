package cmd

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"

	"github.com/WayneShenHH/servermodule/logger"
)

var parseCmd = &cobra.Command{
	Short: "parse game-record",
	Long:  `parse game-record`,
	Use:   "parse",

	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug(cmd.Short)
		logger.Info("args:", args)
		if len(args) < 1 {
			return
		}

		sp := strings.Split(args[0], "-")

		if len(sp) < 2 {
			return
		}

		ans := sp[0]

		u, _ := strconv.ParseUint(sp[1], 16, 64)
		date := time.UnixMilli(int64(u))

		for i := 1; i < len(sp); i++ {
			u, _ := strconv.ParseUint(sp[i], 16, 64)
			ans += fmt.Sprintf("-%v", u)
		}

		logger.Info("ans:", ans, "\ndate:", date.Format(time.RFC3339))
	},
}

func init() {
	RootCmd.AddCommand(parseCmd)
}
