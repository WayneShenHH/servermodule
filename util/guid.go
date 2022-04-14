package util

import "github.com/rs/xid"

// Genereate global unique id for tracing any request
func GenerateGuid() string {
	return xid.New().String()
}
