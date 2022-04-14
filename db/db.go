package db

import "github.com/WayneShenHH/servermodule/protocol"

// NoSQL database interface
type NoSQL interface {
	UpdateDocument(collection, key string, data interface{}) protocol.ErrorCode
	NestedUpdate(collection, version string, params map[string]string) protocol.ErrorCode
	NestedSelectSummary(collection string, params []string) string
}
