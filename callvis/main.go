// グラフ出力用

package main

import (
	"io"
	"os"
)

func main() {
	// go-callvis -group pkg,type -focus os ./
	// go-callvis -group pkg,type -focus io ./
	f, _ := os.Open("README.md")
	_, _ = io.WriteString(f, "0123456789")
	defer f.Close()
}
