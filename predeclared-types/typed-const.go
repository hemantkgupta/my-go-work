package main

import "fmt"

func main() {

	const X float32 = 3.14

	const (
		A, B int64   = -3, 5
		Y    float32 = 2.718
	)

	fmt.Println(X, A, B, Y)

	const XX = float32(3.14)

	const (
		AA, BB = int64(-3), int64(5)
		YY     = float32(2.718)
	)
	fmt.Println(XX, AA, BB, YY)

	const b = uint8(255) + uint8(1) //  (constant 256 of type uint8) overflows uint8
	const c = int8(-128) / int8(-1) // (constant 128 of type int8) overflows int8

	const MaxUint_a = uint(^0) // cannot convert ^0 (untyped int constant -1) to type uint

	const MaxUint_b uint = ^0 //  cannot use ^0 (untyped int constant -1) as uint value in constant declaration (overflows)
}
