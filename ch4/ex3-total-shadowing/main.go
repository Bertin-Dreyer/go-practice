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
	var total int // outer total, stays 0

	for i := 0 ; i < 10 ; i++ {
		// total := total + i  declares a NEW total scoped to this loop body.
		// It shadows the outer total; the outer total is never modified.
		// total += i would instead mutate the outer total in place (no shadowing).
		total := total + i
		fmt.Println(total)
	}
	fmt.Println(total) // prints 0: the outer total, unchanged
}
