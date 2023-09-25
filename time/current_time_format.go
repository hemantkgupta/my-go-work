package main

import (
	"fmt"
	"time"
)

func main() {

	now := time.Now()

	fmt.Println("Time: ", now.Format("15:04:05"))
	fmt.Println("Date:", now.Format("Jan 2, 2006"))
	fmt.Println("Timestamp:", now.Format(time.Stamp))
	fmt.Println("ANSIC:", now.Format(time.ANSIC))
	fmt.Println("UnixDate:", now.Format(time.UnixDate))
}
