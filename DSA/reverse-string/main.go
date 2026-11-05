package main

import "fmt"

// ReverseString reverses the input slice of characters in place.
//
// Write a function that reverses a string. The input string is given as an
// array of characters s.
//
// You must do this by modifying the input array in-place with O(1) extra memory.
//
// Example 1:
//
// Input: s = ["h","e","l","l","o"]
// Output: ["o","l","l","e","h"]
//
// Example 2:
//
// Input: s = ["H","a","n","n","a","h"]
// Output: ["h","a","n","n","a","H"]
func reverseString(s []string) {
	// TODO: implement reverseString
}

func main() {
	s := []string{"h", "e", "l", "l", "o"}
	reverseString(s)
	fmt.Println(s) // expect [o l l e h]
}
