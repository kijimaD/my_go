package main

import "fmt"

func main() {
	var i interface{} = "hello"

	s := i.(string)
	// この文は、インターフェースの値iが具体的な型stringを保持し、基になるTの値を変数sに代入することを主張する
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64) // アサーションに失敗する
	fmt.Println(f, ok)

	// f = i.(float64) // panic
	// fmt.Println(f)
}
