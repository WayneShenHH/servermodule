package util

import (
	"time"
)

type Retry struct {
	count    int           // 次數
	interval time.Duration // 間隔

}

// 取得
func NewRetry(count int, interval time.Duration) *Retry {
	return &Retry{
		count:    count,
		interval: interval,
	}
}

// 執行retry
//
// @param func() (error, interface{})  retry func
//
// @return error 錯誤
//
// @return interface{} 回傳資料
func (r *Retry) Run(f func() (error, interface{})) (error, interface{}) {
	var err error
	var result interface{}
	maxCount := r.count
	tick := time.NewTicker(r.interval)
	defer tick.Stop()

	for range tick.C {
		err, result = f()
		if err == nil {
			return nil, result
		}
		if maxCount--; maxCount <= 0 {
			return err, nil
		}
	}
	return err, result
}

// 根據條件判斷嘗試retry
//
// @param func() (error, interface{})  retry func
//
// @param func(error) bool 判斷是否繼續retry的方法
//
// @return error 錯誤
//
// @return interface{} 回傳資料
func (r *Retry) RunAccordingToJudgment(f func() (error, interface{}), isReTry func(error) bool) (error, interface{}) {
	var err error
	var result interface{}
	maxCount := r.count
	tick := time.NewTicker(r.interval)
	defer tick.Stop()

	for range tick.C {
		err, result = f()
		if err == nil {
			return nil, result
		}
		if !isReTry(err) {
			return err, nil
		}
		if maxCount--; maxCount <= 0 {
			return err, nil
		}
	}
	return err, result
}
