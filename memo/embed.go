package main

import (
	"fmt"
)

type flyer interface {
	fly() string
}
type bird struct{}

func (b *bird) fly() string {
	return "bird fly!"
}

type god struct {
	*bird
}

func main() {
	aBird := &bird{}
	aGod := &god{
		bird: aBird,
	}

	if _, ok := interface{}(aGod).(flyer); ok {
		fmt.Println("godはflyインターフェースを実装している")
	}
	fmt.Println(aGod.fly())
}
