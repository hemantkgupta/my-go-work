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
		case <-time.After(time.Second):
			fmt.Println("tick", i)
		}
	}
}
