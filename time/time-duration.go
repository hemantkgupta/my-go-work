package main

import (
	"fmt"
	"time"
)

func main() {
	// A Duration represents the elapsed time between two instants as an int64 nanosecond count.
	// The representation limits the largest representable duration to approximately 290 years.
	var t time.Duration = time.Duration(time.Hour * 24)
	fmt.Println(t)

}
