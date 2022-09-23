// 変数を定義しなくてもフラグをポインタで受け取ることができる

package main

import (
	"flag"
	"fmt"
)

func main() {
	var i = flag.Int("i", 0, "数値")
	var s = flag.String("s", "", "文字列 ")
	var b = flag.Bool("b", false, "真偽値")
	flag.Parse()

	fmt.Println(*i, *s, *b)
}

// $ go run flag_pointer.go -i=123 -s=hoge -b
// 123 hoge true
