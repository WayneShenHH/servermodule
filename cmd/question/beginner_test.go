package question_test

import (
	"fmt"
	"testing"
)

/*
完成 Cat 實做，符合 Animal 界面
*/
func Test_Implement(t *testing.T) {
	cat := NewCat("Jack")
	cat.Move()  // 印出 "Jack : cat walk"
	cat.Speak() // 印出 "Jack : miew"
	cat.Rename("Mike")
	cat.Move()  // 印出 "Mike : cat walk"
	cat.Speak() // 印出 "Mike : miew"
}

type Animal interface {
	Move()
	Speak()
	Rename(name string)
}

func NewCat(name string) Animal {
	return &cat{
		name: name,
	}
}

type cat struct {
	name string
}

func (c *cat) Move() {
	fmt.Println(c.name, ": cat walk")
}

func (c *cat) Speak() {
	fmt.Println(c.name, ": miew")
}

func (c *cat) Rename(name string) {
	c.name = name
}
