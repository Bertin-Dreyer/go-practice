package main

import "fmt"

// Exercise 5.3:
// Write a function called prefixer that has an input parameter of type string
// and returns a function that has an input parameter of type string and returns a
// string. The returned function should prefix its input with the string passed into
// prefixer.
//
// How it works, step by step:
//  1. We define a function called prefixer.
//  2. prefixer takes one string parameter called input (the prefix).
//  3. prefixer returns another function.
//  4. The returned function takes one string and returns one string.
//  5. Inside the returned function, we combine input + " " + s.
//  6. Because of closures, the returned function remembers input
//     even after prefixer has finished running.
//  7. In main, we call prefixer("Hello") to get a function, then call
//     that function with different names.
//
// Why this matters:
//
//	Functions in Go are values. They can be returned from other functions.
//	A closure remembers variables from its surrounding scope.
func main() {
	// Step 7a: Create a prefixer function with prefix "Hello".
	helloPrefix := prefixer("Hello")

	// Step 7b: Call the returned function with "Bob".
	fmt.Println(helloPrefix("Bob")) // should print Hello Bob

	// Step 7c: Call the same returned function with "Maria".
	fmt.Println(helloPrefix("Maria")) // should print Hello Maria
}

// prefixer takes a string and returns a function.
// The returned function prefixes its input with the captured string.
func prefixer(input string) func(string) string {
	// Step 5: Return a function that captures input.
	return func(s string) string {
		return input + " " + s
	}
}
