package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	data, err := os.ReadFile("sample.txt")
	if err != nil {
		fmt.Println("Error reading sample.txt:", err)
		return
	}

	text := string(data)

	// Split on whitespace into individual words, all lowercase.
	// words = ["the", "quick", "brown", "fox", "jumps", "over", "the", "lazy",
	//          "seal", "in", "cape", "town!"]
	words := strings.Fields(strings.ToLower(text))

	counts := make(map[string]int)

	for _, word := range words {
		// Trim leading/trailing non-letter/non-digit chars (e.g. "town!" -> "town",
		// "state-of-the-art" stays intact because hyphens are in the middle).
		word = strings.TrimFunc(word, func(r rune) bool {
			return !unicode.IsLetter(r) && !unicode.IsDigit(r)
		})
		if word == "" {
			continue
		}

		// Count the cleaned word. map[string]int zero value is 0,
		// so counts[word]++ starts at 1 for new words.
		counts[word]++
	}

	// Print each word and its frequency.
	for word, count := range counts {
		fmt.Println(word, count)
	}
}
