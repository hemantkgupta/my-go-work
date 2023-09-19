package main

func main() {
	const a = -1.23
	// The type of b is deduced as float64.
	// var b = a
	// error: constant 1.23 truncated to integer.
	// var x = int32(a)
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
