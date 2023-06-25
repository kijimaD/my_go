package main

import (
	"unsafe"
)

func main() {
	// 1
	// fmt.Print("Hello world!")

	// 2
	// os.Stdout.Write([]byte("Hello world!"))

	// 3
	// os.Stdout.Write(
	// 	[]byte{72, 101, 108, 108, 111, 32, 119, 111, 113, 108, 100, 33})

	// 4
	// syscall.Write(
	// 	1,
	// 	[]byte{72, 101, 108, 108, 111, 32, 119, 111, 113, 108, 100, 33}
	// )

	// 5
	// fd := 1
	// p := []byte{72, 101, 108, 108, 111, 32, 119, 111, 113, 108, 100, 33}
	// _p0 := unsafe.Pointer(&p[0])
	// syscall.Syscall(1, uintptr(fd), uintptr(_p0), uintptr(len(p)))

	// 6
	// go build -o print
	fd := 1
	p := []byte{72, 101, 108, 108, 111, 32, 119, 111, 113, 108, 100, 33}
	_p0 := unsafe.Pointer(&p[0])
	myPrint(1, uintptr(fd), uintptr(_p0), uintptr(len(p)))
}

func myPrint(trap, a1, a2, a3 uintptr)
