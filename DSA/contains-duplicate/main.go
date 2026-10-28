package main

import "fmt"

// ContainsDuplicate: given an integer array nums, return true if any value
// appears at least twice in the array, and return false if every element is
// distinct.
//
// Example:
//
//	nums = [1, 2, 3, 1] -> true
//	nums = [1, 2, 3, 4] -> false
func containsDuplicate(nums []int) bool {
	// TODO: implement
	seen := make(map[int]bool)

	for _, num := range nums {
		if seen[num] {
			return true
		}
		seen[num] = true
	}
	return false
}

func main() {
	fmt.Println(containsDuplicate([]int{1, 2, 3, 1})) // expect true
	fmt.Println(containsDuplicate([]int{1, 2, 3, 4})) // expect false
}
