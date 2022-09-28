package main

import (
	"strings"
	"fmt"
)

func main() {
	str := "Helloooooooooooooooooooaaaaaaaaaaaaaaaaaaaaaaaaa"
	rd := strings.NewReader(str)

	row := make([]byte, 10)
	rd.Read(row)

	fmt.Println(string(row)) // Helloooooo 10文字分
}
