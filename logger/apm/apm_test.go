package apm_test

import (
	"testing"
	"time"

	"github.com/WayneShenHH/servermodule/logger/apm"
)

func Test_ApmInit(t *testing.T) {
	apm.Init("GS", "local", "http://127.0.0.1:8200")
	tx := apm.StartTransaction("E1_G1_T1_R1", "G1")
	defer tx.End()

	span(tx, "apm1")
	span(tx, "apm2")

	// apm-agent send logs to server after a while
	select {}
}

func span(tx *apm.Apm, functionName string) {
	sp := tx.StartSpan(functionName, "G1", nil)
	defer sp.End()
	time.Sleep(time.Second)
	tx.Debug(functionName, "done")
}
