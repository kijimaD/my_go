package main

import (
	"github.com/fatih/color"
	"fmt"
)

func main() {
	color.Green("Green")
	color.Red("Red")
	color.Yellow("Yellow")

	// ラッパー関数
	Greenf("this is green")
	Redf("this is red")
	Yellowf("this is yellow")
}

func Greenf(t string, args ...interface{}){
	color.Green(fmt.Sprintf(t, args...))
}

func Redf(t string, args ...interface{}){
	color.Red(fmt.Sprintf(t, args...))
}

func Yellowf(t string, args ...interface{}){
	color.Yellow(fmt.Sprintf(t, args...))
}
