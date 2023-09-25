package main

import (
	"fmt"
	"time"
)

func main() {
	// func time.Parse(layout string, value string) (time.Time, error)
	// Parse parses a formatted string and returns the time value it represents.
	// See the documentation for the constant called Layout to see how to represent the format.
	// The second argument must be parseable using the format string (layout) provided as the first argument.
	t, err := time.Parse("2006-02-01 15:04:05 -0700", "2016-13-03 00:00:00 +0000")
	if err != nil {
		fmt.Println("Error in parsing string to time: ", err)
	}
	// func (time.Time).Format(layout string) string
	// Format returns a textual representation of the time value
	// formatted according to the layout defined by the argument.
	// See the documentation for the constant called Layout to see how to represent the layout format.
	fmt.Println(t.Format("January 2, 2006 at 3:04:05PM MST"))
}
