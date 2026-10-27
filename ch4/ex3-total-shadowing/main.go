package main

import "fmt"

// Exercise 4.3:
// Start a new program. In main, declare an int variable called total.
// Write a for loop that uses a variable named i to iterate from 0 (inclusive)
// to 10 (exclusive).
// The body of the for loop should be:
//   total := total + i
//   fmt.Println(total)
// After the for loop, print out the value of total.
// what's happening: this demonstrates variable shadowing. inside the loop,
// `total := total + i` declares a new `total` scoped to the loop body using
// :=. the outer `total` is never actually modified. after the loop we print
// the outer total which remains at its initial value.
func main() {
}
