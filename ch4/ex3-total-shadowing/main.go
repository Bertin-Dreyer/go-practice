package main

import "fmt"

// Exercise 4.3:
// Start a new program. In main, declare an int variable called total.
// Write a for loop that uses a variable named i to iterate from 0 (inclusive)
// to 10 (exclusive).
// The body of the for loop should be:
//
//	total := total + i
//	fmt.Println(total)
//
// After the for loop, print out the value of total.
//
// How it works, step by step:
//  1. We declare an outer variable called total and set it to 0.
//  2. We write a for loop with i from 0 to 9.
//  3. Inside the loop, we use total := total + i.
//     The := creates a NEW total scoped to the loop body.
//     This shadows the outer total — the outer total is NOT changed.
//  4. We print the inner total each loop.
//  5. After the loop ends, we print the outer total again.
//     It still prints 0 because the outer total was never modified.
//
// Difference:
//
//	total += i   mutates the outer total (no shadowing).
//	total := ... creates a new inner total (shadowing).
func main() {
	// Step 1: Declare outer total, starts at 0.
	var total int

	// Step 2: Loop i from 0 to 9.
	for i := 0; i < 10; i++ {
		// Step 3: Create a new total scoped to this loop body.
		// It shadows the outer total.
		total := total + i

		// Step 4: Print the inner total.
		fmt.Println(total)
	}

	// Step 5: Print the outer total — still 0.
	fmt.Println(total)
}
