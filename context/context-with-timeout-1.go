package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	//d := time.Now().Add()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	defer cancel()
	select {
	case <-ctx.Done():
		fmt.Println("The context is done. And error is: ", ctx.Err())
	/*
		case <-time.After(4 * time.Second):
			fmt.Println("Work done within time.")
	*/
	case <-time.After(6 * time.Second):
		fmt.Println("Work is not done within time.")
	}
}
