package main

import "fmt"

func main() {
	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

// 空のインターフェースは、未知の型の値を扱うコードで使用される。例えば、fmt.Printはinterface{}型の任意の数の引数を受け取る。
