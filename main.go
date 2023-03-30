package main

import (
	"fmt"
	"math/cmplx"
)

var c, python, java bool = true, false, true

var (
	tobe   bool       = false
	maxint uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func add(a, b int) int {
	return a + b
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {

	fmt.Println(add(10, 20))

	var a, b string
	a, b = swap("hello", "world")
	fmt.Println(a, b)

	fmt.Println(split(17))

	var i int = 10
	fmt.Println(i, c, python, java)

	j := "Ashish"
	fmt.Println(j)

	fmt.Printf("Type: %T Value: %v\n", tobe, tobe)
	fmt.Printf("Type: %T Value: %v\n", maxint, maxint)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}

// bool
// string
// int int8 int16 int32 int64
// uint uint8 uint16 uint32 uint64
// byte // alias for uint8
// rune // alias for int32
// represents a unicode code point

// float32 float64
// complex64 complex128
