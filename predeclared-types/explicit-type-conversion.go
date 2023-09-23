package main

import "fmt"

func main() {

	var x int = 10
	var y float64 = 30.2

	// Explicit conversion is needed
	var sum1 float64 = float64(x) + y

	// Explicit conversion is needed
	var sum2 int = x + int(y)
	fmt.Println(sum1)
	fmt.Println(sum2)

	const a = -1.23
	// The type of b is deduced as float64.
	var b = a
	fmt.Printf("The type of b is: %T\n", b)

	// error: constant 1.23 truncated to integer.
	// var x = int64(a)
	// fmt.Printf("The value of x is %d and type is %T\n", x, x)

	// error: cannot assign float64 to int32.
	// var y int32 = b
	// okay: z == -1, and the type of z is int32.
	//       The fraction part of b is discarded.
	// var z = int32(b)

	const k int16 = 255
	// The type of n is deduced as int16.
	// var n = k
	// error: constant 256 overflows uint8.
	// var f = uint8(k + 1)
	// error: cannot assign int16 to uint8.
	// var g uint8 = n + 1
	// okay: h == 0, and the type of h is uint8.
	//       n+1 overflows uint8 and is truncated.
	// var h = uint8(n + 1)
}
