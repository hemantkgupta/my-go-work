package main

import (
	"context"
	"fmt"
)

func main() {

	/*
		type Context interface {
			// Deadline returns the time when work done on behalf of this context should be canceled.
			// Deadline returns ok==false when no deadline is set.
			Deadline() (deadline time.Time, ok bool)

			// Done returns a channel that's closed when work done on behalf of this context should be canceled.
			// The close of the Done channel may happen asynchronously, after the cancel function returns.

			// WithCancel arranges for Done to be closed when cancel is called;
			// WithDeadline arranges for Done to be closed when the deadline expires;
			// WithTimeout arranges for Done to be closed when the timeout elapses.
			Done() <-chan struct{}

			// If Done is not yet closed, Err returns nil.
			// If Done is closed, Err returns a non-nil error
			Err() error
			Value(key any) any
		}
	*/

	ctx := context.Background()
	fmt.Println("ctx.Err() : ", ctx.Err())
	fmt.Println("ctx.Done() : ", ctx.Done())
	fmt.Println("ctx.Value(\"key\") : ", ctx.Value("key"))
	time, ok := ctx.Deadline()
	fmt.Println("ctx.Deadline() : ", time, ok)

}
