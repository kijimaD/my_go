package main

import (
	"fmt"
	"os"
)

func main() {
	// 読み込み権限onlyで開く
	f, err := os.Open("text.txt")

	// ファイル内容の読み込み
	data := make([]byte, 1024)
	count, err := f.Read(data)
	if err != nil {
		fmt.Println(err)
		fmt.Println("fail to read file")
	}

	fmt.Printf("read %d bytes:\n", count)
	fmt.Println(data) // 1024バイト分で、未定義のバイトもある
	fmt.Println("raw byte: ", data[:count]) // countにはopenで何バイト読み込んだか入っている。1024バイトのうち、その長さ分だけ表示すればよい
	fmt.Println(string(data[:count])) // 文字列にする
}
