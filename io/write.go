package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("write.txt")
	str := "write this file by Golang!"
	data := []byte(str)
	count, err := f.Write(data) // []byteスライスに格納されている内容が、ファイルにそのまま書き込まれる
	// countにはメソッド実行の結果、ファイルに何byte書き込まれたか
	if err != nil {
		fmt.Println(err)
		fmt.Println("fail to write file")
	}

	fmt.Printf("write %d bytes\n", count)
}
