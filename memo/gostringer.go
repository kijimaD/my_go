package main

import (
	"fmt"
)

type test struct {
	Name string
}

func (t test) GoString() string {
	return "this is GoString()"
}

func (t test) String() string {
	return "this is String()"
}

func main() {
	t := test{
		Name: "aaa",
	}

	fmt.Printf("%#v\n", t)
	// this is GoString()

	fmt.Printf("%+f\n", t)
	// {%!f(string=aaa)}

	fmt.Printf("%#f\n", t)
	// {%!f(string=aaa)}

	fmt.Println(t)
	// this is String()
}
