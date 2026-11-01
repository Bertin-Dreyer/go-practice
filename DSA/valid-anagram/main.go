package main

import (
	"fmt"
	"strings"
)

// Problem 1: Valid Anagram
// Given two strings s and t, return true if t is an anagram of s,
// meaning it has the same letters with the same counts.
//
// Example:
//
//	s = "listen", t = "silent"  -> true
//	s = "rat",   t = "car"      -> false
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	anaMap := map[rune]int{}

	for _, letter := range strings.ToLower(s) {
		anaMap[letter]++
	}

	for _, letter := range strings.ToLower(t) {
		anaMap[letter]--
	}
	for _, count := range anaMap {
		if count != 0 {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(isAnagram("listen", "silent")) // expect true
	fmt.Println(isAnagram("rat", "car"))       // expect false
}
