package main

import (
	"fmt"
	"math/rand"
)

// Exercise 4.1:
// Write a for loop that puts 100 random numbers between 0 and 100 into an int slice.
//
// How it works, step by step:
//  1. We create an empty int slice called nums.
//  2. We loop 100 times with i from 0 to 99.
//  3. On each loop, we call rand.Intn(101) to get a random number from 0 to 100.
//  4. We append that number to the nums slice.
//  5. After the loop, we print the slice.
//     rand.Intn(n) returns a value in [0, n), so we pass 101 to include 100.
func main() {
	// Step 1: Create an empty int slice.
	var nums []int

	// Step 2: Loop 100 times.
	for i := 0; i < 100; i++ {
		// Step 3: Get a random number from 0 to 100.
		// Step 4: Append it to the slice.
		nums = append(nums, rand.Intn(101))
	}

	// Step 5: Print the slice.
	fmt.Println(nums)
}
