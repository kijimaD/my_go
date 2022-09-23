package main

import (
	"flag"
	"fmt"
)

var i int
var s string
var b bool

func init(){
	flag.IntVar(&i, "i", 0, "数値")
	flag.StringVar(&s, "s", "default", "文字列 ")
	flag.BoolVar(&b, "b", false, "真偽値 ")
}

func main() {
	flag.Parse()
	fmt.Println(i, s, b)
}

// $ go run flag.go
// 0 default false

// $ go run flag.go -i=123 -s=hoge -b
// 123 hoge true
