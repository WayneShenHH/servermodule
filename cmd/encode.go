package cmd

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"

	"github.com/WayneShenHH/servermodule/logger"
)

var encodeCmd = &cobra.Command{
	Short: "encode game-record",
	Long:  `encode game-record`,
	Use:   "encode",

	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug(cmd.Short)
		logger.Info("args:", args)
		if len(args) < 1 {
			return
		}

		encode(args[0])

		// logger.Info("ans:", strings.ToUpper(ans), "\ndate:", date.Format(time.RFC3339))
	},
}

func encode(src string) []string {
	list := []string{}
	err := json.Unmarshal([]byte(src), &list)
	if err != nil {
		panic(err)
	}
	anslist := []string{}
	for idx := range list {
		sp := strings.Split(list[idx], "-")

		if len(sp) < 2 {
			continue
		}

		ans := sp[0]

		u, _ := strconv.ParseUint(sp[1], 10, 64)
		date := time.UnixMilli(int64(u)).UTC()

		for i := 1; i < len(sp); i++ {
			u, _ := strconv.ParseInt(sp[i], 10, 64)
			h := strconv.FormatInt(u, 16)
			ans += fmt.Sprintf("-%v", h)
		}

		logger.Info("list:", list[idx], "ans:", strings.ToUpper(ans), "\ndate:", date.Format(time.RFC3339))
		anslist = append(anslist, strings.ToUpper(ans))
	}

	return anslist
}

func init() {
	RootCmd.AddCommand(encodeCmd)
}
