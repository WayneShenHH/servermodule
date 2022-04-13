package slackalert_test

import (
	"testing"

	"github.com/WayneShenHH/servermodule/slackalert"
)

func init() {
}

func Test_slack(t *testing.T) {
	alertsInfo := slackalert.SlackAlertsInfo{
		SlackEnable:  true,
		SlackChannel: "#jenkins_alert_test",
		Env:          "local",
	}

	slackalert.Run(
		"https://hooks.slack.com/services/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",
		&alertsInfo,
		"Gamemaster",
		"asdf",
		"",
	)
	slackalert.SendStart("啟動")
}
