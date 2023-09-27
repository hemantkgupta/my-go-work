package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	done := ctx.Done()

	for i := 0; ; i++ {
		select {
		case <-done:
			fmt.Println("Context is done.")
			return
		default:
			fmt.Println("Work in progress...")
		}
		time.Sleep(1 * time.Second)
	}
}
