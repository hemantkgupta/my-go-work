package main

import (
	"fmt"
	"time"
)

func main() {
	// func time.Now() time.Time
	// Now returns the current local time.
	var now time.Time = time.Now()

	fmt.Println("The current datetime is:", now)

	utc := time.Now().UTC()
	fmt.Println("The current datetime in UTC is:", utc)

	fmt.Println("Year:", now.Year())
	fmt.Println("Month:", now.Month())
	fmt.Println("Day:", now.Day())
	fmt.Println("Hour:", now.Hour())
	fmt.Println("Minute:", now.Minute())
	fmt.Println("Second:", now.Second())
	fmt.Println("Nanosecond:", now.Nanosecond())
}
