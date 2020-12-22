package slackAlert_test

import (
	"testing"

	"github.com/WayneShenHH/servermodule/slackAlert"
)

func init() {
}

func Test_slack(t *testing.T) {
	alertsInfo := slackAlert.SlackAlertsInfo{
		SlackEnable:  true,
		SlackChannel: "#jenkins_alert_test",
		Env:          "local",
	}

	slackAlert.Run(
		"https://hooks.slack.com/services/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",
		&alertsInfo,
		"Gamemaster",
		"asdf",
		"",
	)
	slackAlert.SendStart("啟動")
}
