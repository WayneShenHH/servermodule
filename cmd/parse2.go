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

var parse2Cmd = &cobra.Command{
	Short: "parse2 game-record",
	Long:  `parse2 game-record`,
	Use:   "parse2",

	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug(cmd.Short)
		logger.Info("args:", args)
		if len(args) < 1 {
			return
		}

		parse2(args[0])

		// logger.Info("ans:", strings.ToUpper(ans), "\ndate:", date.Format(time.RFC3339))
	},
}

func parse2(src string) []string {
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

		u, _ := strconv.ParseUint(sp[1], 16, 64)
		date := time.UnixMilli(int64(u)).UTC()

		for i := 1; i < len(sp); i++ {
			u, _ := strconv.ParseUint(sp[i], 16, 64)
			ans += fmt.Sprintf("-%v", u)
		}

		logger.Info("list:", list[idx], "ans:", strings.ToUpper(ans), "\ndate:", date.Format(time.RFC3339))
		anslist = append(anslist, strings.ToUpper(ans))
	}

	return anslist
}

func init() {
	RootCmd.AddCommand(parse2Cmd)
}
