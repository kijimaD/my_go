package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func Scale2(v Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	Scale(&v, 10)
	fmt.Println(v) // ポインタによって、vの値が変わっている
	fmt.Println(Abs(v))

	v2 := Vertex{3, 4}
	Scale2(v, 10)
	fmt.Println(v2) // vの値は変わらない
	fmt.Println(Abs(v2))
}
