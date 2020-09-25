package util

import (
	"fmt"
)

// Color const
type Color int

const printColor = "\033[1;%dm%s\033[0m"

// color list
const (
	DebugColor   Color = 34
	InfoColor    Color = 36
	WarningColor Color = 33
	ErrorColor   Color = 31
	FatalColor   Color = 35
)

// ColorString paint color on input string, use for std output only
func ColorString(color Color, content string) string {
	return fmt.Sprintf(printColor, color, content)
}
