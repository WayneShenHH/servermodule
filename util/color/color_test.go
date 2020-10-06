package color

import (
	"fmt"
	"testing"
)

func Test_Color(t *testing.T) {
	// colors := []int{31, 32, 33, 34, 35, 36, 91, 92, 93, 94, 95, 96}
	for j := 0; j < 256; j++ {
		fmt.Printf("[%3d] \x1b[%dm%s\x1b[0m\n", j, j, "Hello World!")
	}
}
func Test_Color2(t *testing.T) {
	for j := 0; j < 256; j++ {
		fmt.Printf("[%3d] \033[38;5;%dm%s\033[39;49m\n", j, j, "Hello World!")
	}
}
