// Package logger 系統共用 stdout logger
//nolint:unused // 先保留 logger 介面方法
package logger

import (
	json "github.com/json-iterator/go"
	"github.com/sirupsen/logrus"
)

// WithStruct 根據傳入物件轉換為 log entry
func WithStruct(i interface{}) *logrus.Entry {
	fields := getFields(i)
	return WithField(fields)
}

func getFields(i interface{}) logrus.Fields {
	b, _ := json.Marshal(i)
	fields := logrus.Fields{}
	json.Unmarshal(b, &fields)
	return fields
}
