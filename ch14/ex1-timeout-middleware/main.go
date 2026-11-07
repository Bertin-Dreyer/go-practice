package main

import "fmt"

// Exercise 14.1:
// Create a middleware-generating function that creates a context with a timeout.
// The function should have one parameter (milliseconds allowed).
// It should return a func(http.Handler) http.Handler.

// TODO: Write TimeoutMiddleware(ms int) func(http.Handler) http.Handler
// TODO: Inside, use context.WithTimeout with the given milliseconds
// TODO: Call next.ServeHTTP with the timeout context via r.WithContext(ctx)
// TODO: Remember to call cancel() after the request

func main() {
	// TODO: Implement timeout middleware
	fmt.Println("TODO: implement ex1-timeout-middleware")
}
