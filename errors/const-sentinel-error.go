package main

import (
	"fmt"
	"github.com/hemantkgupta/my-go-work/goerror/consterr"
)

func main() {
	const (
		ErrFoo = consterr.Sentinel("foo error")
		ErrBar = consterr.Sentinel("bar error")
	)
	fmt.Println(ErrFoo, ErrBar)
}
