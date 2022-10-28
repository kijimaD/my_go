package main

import (
	"encoding/json"
	"fmt"
)

type ConfidentialCustomer struct {
	CustomerID int64
	CreditCard CreditCard
}

type CreditCard string

func (c CreditCard) String() string {
	return "xxxx-xxxx-xxxx-xxxx"
}

func (c CreditCard) GoString() string {
	return "xxxx-xxxx-xxxx-xxxx"
}

func main() {
	c := ConfidentialCustomer{
		CustomerID: 1,
		CreditCard: "1111-1111-1111-1111",
	}

	fmt.Println(c)
	fmt.Printf("%v\n", c)
	fmt.Printf("%+v\n", c)
	fmt.Printf("%#v\n", c)

	// マスキングされない。プログラム本体で利用する関数の挙動に
	bytes, _ := json.Marshal(c)
	fmt.Println("JSON: ", string(bytes))
}
