package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"unicode"
)

func main() {
	data, err := os.ReadFile("bible.txt")
	if err != nil {
		fmt.Println("Error reading bible.txt:", err)
		return
	}

	text := string(data)

	// Split on whitespace into individual words, all lowercase.
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

	var freq []wordCount

	for word, count := range counts {
		freq = append(freq, wordCount{word, count})
	}

	sort.Slice(freq, func(i, j int) bool {
		return freq[i].count > freq[j].count
	})

	// Print each word and its frequency.
	for i := 0; i < 20 && i < len(freq); i++ {
		fmt.Println(freq[i].word, freq[i].count)
	}
}

type wordCount struct {
	word  string
	count int
}
