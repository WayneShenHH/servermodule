package util

import (
	"errors"
	"testing"
	"time"

	"github.com/WayneShenHH/servermodule/logger"
)

const (
	RetryCount    = 5
	RetryInterval = time.Millisecond * 300
)

func Test_NewRetry(t *testing.T) {
	retry := NewRetry(RetryCount, RetryInterval)
	logger.Infof("NewRetry : %v", retry)
}

func Test_RetryRun(t *testing.T) {
	retry := NewRetry(RetryCount, RetryInterval)
	err, result := retry.Run(func() (error, interface{}) {
		return nil, "test"
	})
	if err != nil {
		logger.Warnf("Test_RetryRun error: %v", err)
		return
	}
	logger.Infof("Test_RetryRun result : %v", result)
}

func Test_RetryError(t *testing.T) {
	count := 0
	retry := NewRetry(RetryCount, RetryInterval)
	err, result := retry.Run(func() (error, interface{}) {
		count++
		logger.Infof("retry count: %v", count)
		return errors.New("Retry Error test"), nil
	})
	if err != nil {
		logger.Warnf("Test_RetryError error: %v", err)
		return
	}
	logger.Infof("Test_RetryError result : %v", result)
}

func Test_RetryRunAccordingToJudgment(t *testing.T) {
	count := 0
	retry := NewRetry(RetryCount, RetryInterval)
	err, result := retry.RunAccordingToJudgment(func() (error, interface{}) {
		count++
		logger.Infof("retry count: %v", count)
		return errors.New("test"), nil
	}, func(e error) bool {
		return !(e.Error() == "test")
	})
	if err != nil {
		logger.Warnf("Test_RetryRunAccordingToJudgment error: %v", err)
		return
	}
	logger.Infof("Test_RetryRunAccordingToJudgment result : %v", result)
}
