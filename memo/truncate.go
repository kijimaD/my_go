package main

import (
	"fmt"
)

func truncateString(s string, b bool) string {
	if b {
		n := 5
		for i := range s {
			n--
			if n < 0 {
				return s[:i]
			}
		}
	}
	return s
}

func main() {
	fmt.Println("aaa"[:2])
	fmt.Println(truncateString("aaaiiiuuu", true))
	fmt.Println(truncateString("aaa", true))
	fmt.Println(truncateString("", true))
}
