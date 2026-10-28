package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	text := " The quick brown fox jumps over the lazy seal in Cape Town!"

	// Split on whitespace into individual words, all lowercase.
	// words = ["the", "quick", "brown", "fox", "jumps", "over", "the", "lazy",
	//          "seal", "in", "cape", "town!"]
	words := strings.Fields(strings.ToLower(text))

	counts := make(map[string]int)

	for _, word := range words {
		// Strip non-letter characters (e.g. "town!" -> "town").
		// strings.Builder avoids O(n^2) string concatenation.
		var b strings.Builder
		for _, letter := range word { // letter is a rune
			if unicode.IsLetter(letter) {
				b.WriteRune(letter)
			}
		}
		cleaned := b.String()

		// Count the cleaned word. map[string]int zero value is 0,
		// so counts[cleaned]++ starts at 1 for new words.
		counts[cleaned]++
	}

	// Print each word and its frequency.
	for word, count := range counts {
		fmt.Println(word, count)
	}
}
