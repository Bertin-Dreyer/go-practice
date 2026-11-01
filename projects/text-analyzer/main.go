package main

import (
	"flag"
	"fmt"
	"os"
	"sort"
	"strings"
	"unicode"
)

type wordCount struct {
	word  string
	count int
}

func main() {
	// Step 1: Define command-line flags with sensible defaults.
	// flag.String returns a *string (pointer to string).
	// flag.Int returns a *int (pointer to int).
	filename := flag.String("file", "bible.txt", "input text file to analyze")
	topN := flag.Int("top", 20, "number of top words to print")

	// Step 2: Parse the flags from os.Args.
	// This reads things like -file sample.txt -top 10 from the command line.
	flag.Parse()

	// Step 3: Read the file using the filename from the flag.
	// Dereference the pointer with *filename to get the actual string value.
	data, err := os.ReadFile(*filename)
	if err != nil {
		fmt.Println("Error reading", *filename, ":", err)
		return
	}

	// Step 4: Convert bytes to string.
	text := string(data)

	// Step 5: Count each cleaned, lowercased word.
	counts := countWords(text)

	var freq []wordCount

	for word, count := range counts {
		freq = append(freq, wordCount{word, count})
	}

	sort.Slice(freq, func(i, j int) bool {
		return freq[i].count > freq[j].count
	})

	// Step 6: Use the topN flag instead of hardcoding 20.
	// Cap it at len(freq) so we don't go out of bounds if there are fewer unique words.
	limit := *topN
	if limit > len(freq) {
		limit = len(freq)
	}

	// Print each word and its frequency.
	for i := 0; i < limit && i < len(freq); i++ {
		fmt.Println(freq[i].word, freq[i].count)
	}
}

// cleanWord trims non-letter, non-digit characters from both ends of a word.
func cleanWord(word string) string {
	return strings.TrimFunc(word, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsDigit(r)
	})
}

// countWords returns how many times each cleaned, lowercased word appears.
func countWords(text string) map[string]int {
	counts := make(map[string]int)
	words := strings.Fields(strings.ToLower(text))

	for _, w := range words {
		w = cleanWord(w)
		if w == "" {
			continue
		}
		counts[w]++
	}
	return counts
}
