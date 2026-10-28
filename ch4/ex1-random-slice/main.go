package main

import (
	"fmt"
	"math/rand"
)

// Exercise 4.1:
// Write a for loop that puts 100 random numbers between 0 and 100 into an int slice.
// what's happening: we create an empty int slice, then loop 100 times appending
// a random number in [0, 100] on each iteration. rand.Intn(n) returns a value in
// [0, n), so we pass 101 to include 100.
func main() {
	var nums []int
	for i := 0; i < 100; i++ {
		nums = append(nums, rand.Intn(101))
	}
	fmt.Println(nums)
}
