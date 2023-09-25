package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	t1 := now.Add(time.Hour * 27)
	fmt.Println(t1.Format(time.UnixDate))

	t2 := now.AddDate(2, 10, 11)
	fmt.Println(t2.Format(time.UnixDate))

	t3 := now.Add(-time.Hour * 6)
	fmt.Println(t3.Format(time.UnixDate))
}
