package main

import (
	"fmt"
	"unicode/utf8"
)

func truncate(b []byte) []byte {
	if true {
		n := 5
		for i := 0; i < len(b); {
			n--
			if n < 0 {
				return b[:i]
			}
			wid := 1
			if b[i] >= utf8.RuneSelf {
				_, wid = utf8.DecodeRune(b[i:])
			}
			// 文字のバイト数分ループを飛ばす
			i += wid
		}
	}
	return b
}

func main() {
	test := []byte("abcdefg")
	fmt.Println(test, truncate(test)) // truncateされる
	nihon := []byte("日本語日本語")
	fmt.Println(len(nihon), len(truncate(nihon)))
}
