// go-callvis -group pkg,type -focus os ./

package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Open("README.md")
	if err != nil {
		panic(err)
	}
	n, _ := io.WriteString(f, "0123456789")
	fmt.Println(n)
	defer f.Close()
}
