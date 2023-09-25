package main

import (
	"fmt"
	"time"
)

func main() {
	// func time.Date(year int, month time.Month, day int, hour int, min int, sec int, nsec int,
	//	loc *time.Location) time.Time
	//
	// Date returns the Time corresponding to
	// yyyy-mm-dd hh:mm:ss + nsec nanoseconds in the appropriate zone for that time in the given location.
	t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Printf("Go launch date: %s\n", t.Local())
}
