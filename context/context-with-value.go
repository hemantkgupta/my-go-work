package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "Name", "Hemant")
	fmt.Println(ctx.Value("Name"))
}
