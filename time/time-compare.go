package main

import (
	"fmt"
	"time"
)

func main() {
	t1 := time.Now()

	// func (time.Time).Add(d time.Duration) time.Time
	// Add returns the time t+d.
	t2 := t1.Add(time.Hour * 24)

	// func (time.Time).Before(u time.Time) bool
	// Before reports whether the time instant t is before u.
	fmt.Println(t1.Before(t2))

	// func (time.Time).Sub(u time.Time) time.Duration
	// Sub returns the duration t-u.
	// If the result exceeds the maximum (or minimum) value that can be stored in a Duration,
	// the maximum (or minimum) duration will be returned.
	// To compute t-d for a duration d, use t.Add(-d).
	fmt.Println(t1.Sub(t2))
}
