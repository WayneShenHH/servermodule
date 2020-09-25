package benchmark

import (
	"fmt"
	"testing"

	"github.com/WayneShenHH/servermodule/config"
	"github.com/WayneShenHH/servermodule/logger/zaphdr"
)

func Benchmark_DevError(b *testing.B) {
	zaphdr.New(&config.LoggerConfig{
		StdLevel: config.Debug,
	})
	msgStr := "string"
	msgStruct := map[string]interface{}{
		"int": 1,
		"str": "string",
	}
	err := fmt.Errorf("error")
	for i := 0; i < b.N; i++ {
		zaphdr.Error(msgStr, err, msgStruct, 503)
	}
}

func Benchmark_DevWarn(b *testing.B) {
	zaphdr.New(&config.LoggerConfig{
		StdLevel: config.Debug,
	})
	msgStr := "string"
	msgStruct := map[string]interface{}{
		"int": 1,
		"str": "string",
	}
	err := fmt.Errorf("error")
	for i := 0; i < b.N; i++ {
		zaphdr.Warn(msgStr, msgStruct, err, 503)
	}
}
func Benchmark_DevInfo(b *testing.B) {
	zaphdr.New(&config.LoggerConfig{
		StdLevel: config.Debug,
	})
	msgStr := "string"
	msgStruct := map[string]interface{}{
		"int": 1,
		"str": "string",
	}
	err := fmt.Errorf("error")
	for i := 0; i < b.N; i++ {
		zaphdr.Info(msgStr, msgStruct, err, 503)
	}
}
func Benchmark_DevDebug(b *testing.B) {
	zaphdr.New(&config.LoggerConfig{
		StdLevel: config.Debug,
	})
	msgStr := "string"
	msgStruct := map[string]interface{}{
		"int": 1,
		"str": "string",
	}
	err := fmt.Errorf("error")
	for i := 0; i < b.N; i++ {
		zaphdr.Debug(msgStr, msgStruct, err, 503)
	}
}

func Benchmark_ProdError(b *testing.B) {
	zaphdr.New(&config.LoggerConfig{
		StdLevel: config.Info,
	})
	msgStr := "string"
	msgStruct := map[string]interface{}{
		"int": 1,
		"str": "string",
	}
	err := fmt.Errorf("error")
	for i := 0; i < b.N; i++ {
		zaphdr.Error(msgStr, err, msgStruct, 503)
	}
}

func Benchmark_ProdWarn(b *testing.B) {
	zaphdr.New(&config.LoggerConfig{
		StdLevel: config.Info,
	})
	msgStr := "string"
	msgStruct := map[string]interface{}{
		"int": 1,
		"str": "string",
	}
	err := fmt.Errorf("error")
	for i := 0; i < b.N; i++ {
		zaphdr.Warn(msgStr, msgStruct, err, 503)
	}
}
func Benchmark_ProdInfo(b *testing.B) {
	zaphdr.New(&config.LoggerConfig{
		StdLevel: config.Info,
	})
	msgStr := "string"
	msgStruct := map[string]interface{}{
		"int": 1,
		"str": "string",
	}
	err := fmt.Errorf("error")
	for i := 0; i < b.N; i++ {
		zaphdr.Info(msgStr, msgStruct, err, 503)
	}
}
func Benchmark_ProdDebug(b *testing.B) {
	zaphdr.New(&config.LoggerConfig{
		StdLevel: config.Info,
	})
	msgStr := "string"
	msgStruct := map[string]interface{}{
		"int": 1,
		"str": "string",
	}
	err := fmt.Errorf("error")
	for i := 0; i < b.N; i++ {
		zaphdr.Debug(msgStr, msgStruct, err, 503)
	}
}
