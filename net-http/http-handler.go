package main

import (
	"fmt"
	"net/http"
)

type CounterHandler struct {
}

func (ct CounterHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Counter")
}

func main() {
	th := &CounterHandler{}
	// A ResponseWriter interface is used by an HTTP handler to construct an HTTP response.
	/*
		type ResponseWriter interface {
			Header() http.Header
			Write([] byte)(int, error)
			WriteHeader(statusCode int)
		}

		type Handler interface {
			ServeHTTP(ResponseWriter, *Request)
		}
	*/

	// func http.Handle(pattern string, handler http.Handler)
	// Handle registers the handler for the given pattern in the DefaultServeMux.
	// The documentation for ServeMux explains how patterns are matched.
	http.Handle("/count", th)

	// func http.ListenAndServe(addr string, handler http.Handler) error
	// ListenAndServe listens on the TCP network address addr
	// and then calls Serve with handler to handle requests on incoming connections.
	// Accepted connections are configured to enable TCP keep-alives.
	// The handler is typically nil, in which case the DefaultServeMux is used.

	http.ListenAndServe(":8080", nil)
}
