package util_test

import (
	"fmt"
	"testing"

	"github.com/WayneShenHH/servermodule/util"
)

func Test_Sort(t *testing.T) {
	a := []int{2, 1, 3}
	util.GenericSort(a, func(a, b int) bool { return a < b })
	fmt.Println(a)

	b := []string{"2", "1", "3"}
	util.GenericSort(b, func(a, b string) bool { return a < b })
	fmt.Println(b)
}
