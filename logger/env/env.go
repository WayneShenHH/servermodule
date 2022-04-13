package env

import (
	"os"
	"strconv"

	"github.com/WayneShenHH/servermodule/logger/constants"
)

// Setting env cfg
var Setting *TraceConfig

// TraceConfig trace cfg
type TraceConfig struct {
	Level     string
	Formatter string
	Code      constants.ServiceCode
}

func init() {
	Setting = new(TraceConfig)
	Setting.Level = os.Getenv("TRACE_LEVEL")
	Setting.Formatter = os.Getenv("TRACE_FORMATTER")
	envcode := os.Getenv("TRACE_CODE")
	if len(envcode) > 0 {
		i, err := strconv.Atoi(envcode)
		if err == nil {
			Setting.Code = constants.ServiceCode(i)
		}
	}
}
