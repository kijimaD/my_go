package main

import (
	"fmt"
	"math/rand"
	"math"
	"runtime"
)

func add(x int, y int) int {
	return x + y
}

// 関数の型が同じである場合は省略できる
func add2(x, y int) int {
	return x + y
}

// 複数の戻り値
func swap(x, y string) (string, string) {
	return y, x
}

// 戻り血となる変数に名前をつける
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	// return (x, y int) と同じ
	return
}

func var1() {
	// 変数に初期値を与えずに宣言すると、ゼロ値が与えられる
	var i int
	var c, python, java bool
	fmt.Println(i, c, python, java)
}

// 型変換
func conv() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
}

func typeinterface() {
	v := 42
	fmt.Printf("v is of type %T\n", v)
}

func for_example() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

func switch_example() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}

// deferへ渡した関数の実行を、呼び出し元の関数の終わりまで遅延させる
func defer_example() {
	defer fmt.Println("world")

	fmt.Println("hello")

	// hello
	// world
	// になる。
}

func defer_stack() {
	fmt.Println("counting")

	for i :=0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")

	// 9
	// 8
	// 7
	// .
	// .
}

func pointers() {
	i, j := 42, 2701

	p := &i
	fmt.Println(*p) // ポインタpを通してiから値を読み出す
	*p = 21 // ポインタpを通してiへ値を代入する
	fmt.Println(i)

	p = &j
	*p = *p / 37 // ポインタpを通してjに値を代入する
	fmt.Println(j)
}

type Vertex struct {
	X int
	Y int
}

func struct_example() {
	fmt.Println(Vertex{1, 2})
}

func struct_access() {
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)
}

func structs() {
	var (
		v1 = Vertex{1, 2}
		v2 = Vertex{X: 1}
		v3 = Vertex{}
		p = &Vertex{1, 2}
	)
	fmt.Println(v1, p, v2, v3)
}

func array_example() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}

func slice_example() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s)
}

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))

	fmt.Println(add(42, 13))

	fmt.Println(add2(42, 13))

	a, b := swap("hello", "world")
	fmt.Println(a, b)

	fmt.Println(split(17))

	var1()

	conv()

	typeinterface()

	for_example()

	switch_example()

	defer_example()

	defer_stack()

	pointers()

	struct_example()

	struct_access()

	structs()

	array_example()

	slice_example()
}
