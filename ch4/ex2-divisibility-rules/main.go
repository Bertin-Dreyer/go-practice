package main

import (
	"fmt"
	"math/rand"
)

// Exercise 4.2:
// Loop over the slice you created in exercise 1. For each value:
//
//	a. If divisible by 2, print "Two!"
//	b. If divisible by 3, print "Three!"
//	c. If divisible by both 2 and 3, print "Six!" (don't print anything else)
//	d. Otherwise, print "Never mind"
//
// How it works, step by step:
//  1. We create a slice of 100 random numbers (same as exercise 1).
//  2. We loop over each number in the slice.
//  3. For each number, we check divisibility with the % operator.
//  4. Order matters: we check the combined case (divisible by both 2 and 3)
//     BEFORE the individual cases so we only print one message per value.
//  5. We use continue to skip the rest of the loop body after printing.
func main() {
	// Step 1: Create the slice of 100 random numbers.
	var nums []int
	for i := 0; i < 100; i++ {
		nums = append(nums, rand.Intn(101))
	}

	// Step 2: Loop over each number.
	for _, num := range nums {
		// Step 3: Check combined case first (divisible by both 2 and 3).
		if num%2 == 0 && num%3 == 0 {
			fmt.Println("Six!")
			continue // skip to next number
		}

		// Step 4: Check divisible by 2 only.
		if num%2 == 0 {
			fmt.Println("Two!")
			continue // skip to next number
		}

		// Step 5: Check divisible by 3 only.
		if num%3 == 0 {
			fmt.Println("Three!")
			continue // skip to next number
		}

		// Step 6: Otherwise — divisible by neither.
		fmt.Println("Never mind")
	}
}
