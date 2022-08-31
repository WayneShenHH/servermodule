package question_test

import (
	"fmt"
	"testing"
)

/*
完成 Cat 實做，符合 Animal 界面
*/
func Test_Implement(t *testing.T) {
	cat := NewCat()
	cat.Move()  // 印出 "cat walk"
	cat.Speak() // 印出 "miew"
}

type Animal interface {
	Move()
	Speak()
}

func NewCat() Animal {
	return &cat{}
}

type cat struct {
	name string
}

func (c *cat) Move() {
	fmt.Println("cat walk")
}
func (c *cat) Speak() {
	fmt.Println("miew")
}
