package main

import "fmt"

// Declare two individual constants. Yes,
// non-ASCII letters can be used in identifiers.
const π = 3.1416
const Pi = π // <=> const Pi = 3.1416

// Declare multiple constants in a group.
const (
	No         = !Yes
	Yes        = true
	MaxDegrees = 360
	Unit       = "radian"
)

func main() {
	// Declare multiple constants in one line.
	const TwoPi, HalfPi, Unit2 = π * 2, π * 0.5, "degree"
	fmt.Println(TwoPi, HalfPi, Unit2, Yes)
}
