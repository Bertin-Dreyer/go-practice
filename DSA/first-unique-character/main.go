package main

import (
	"fmt"
	"strings"
)

// Problem 2: First Unique Character in a String
// Given a string s, return the index of the first character that appears
// only once, or -1 if there isn't one.
//
// Example:
//
//	s = "leetcode" -> 0  ('l' is first unique)
//	s = "aabb"     -> -1 (no unique characters)
//	s = "aA"       -> -1 ('a' and 'A' are the same after ToLower)
//	s = "Aa"       -> -1 ('A' and 'a' are the same after ToLower)
func firstUniqChar(s string) int {
	freqMap := map[rune]int{}

	for _, r := range strings.ToLower(s) {
		freqMap[r]++
	}

	for idx, r := range strings.ToLower(s) {
		if freqMap[r] == 1 {
			return idx
		}
	}
	return -1
}

func main() {
	fmt.Println(firstUniqChar("leetcode")) // expect 0
	fmt.Println(firstUniqChar("aabb"))     // expect -1
	fmt.Println(firstUniqChar("aA"))       // expect -1 (a and A treated as same)
	fmt.Println(firstUniqChar("Aa"))       // expect -1 (A and a treated as same)
}
