package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	// func context.WithTimeout(parent context.Context, timeout time.Duration) (context.Context, context.CancelFunc)
	// WithTimeout returns WithDeadline(parent, time.Now().Add(timeout)).

	// Canceling this context releases resources associated with it,
	// so code should call cancel as soon as the operations running in this [Context] complete:
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	go doSomething(ctx)
	time.Sleep(6 * time.Second)
}

func doSomething(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Timout happened.")
			return
		default:
			fmt.Println("Work in progress...")
		}
		time.Sleep(1 * time.Second)
	}
}
