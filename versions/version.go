// Package versions 應用程式版本號
package versions

import (
	"fmt"
)

const name = "go-mod"

// // MajorVersion 主版號
// const MajorVersion = 0

// // MinorVersion 次版號
// const MinorVersion = 1

// // PatchVersion 修訂號
// const PatchVersion = 0

// Version 版本
var version string

// SetVersion set ver.
func SetVersion(ver string) {
	version = ver
}

// Format 輸出版本編號
func Format() string {
	versionFormat := "v0.0.0"

	// 如果有值傳入
	if version != "" {
		versionFormat = version
	}
	return versionFormat
}

// Full 完整編號包含名稱
func Full() string {
	return fmt.Sprintf(`%v:%v`, Name(), Format())
}

// Name 專案名稱
func Name() string {
	return name
}
