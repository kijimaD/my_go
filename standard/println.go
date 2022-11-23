// go-callvis -group pkg,type -focus fmt ./println.go

package main

import (
	"fmt"
)

func main() {
	fmt.Println("a", "b", "c")
}
