package main

import "fmt"

// Exercise 3.2:
// Define a string variable called message with the value "Hi 👨 and 👩".
// Print the fourth rune in it as a character, not a number.
//
// How it works, step by step:
// 1. Strings in Go are read-only slices of bytes.
// 2. Indexing into a string gives bytes, not runes (characters).
// 3. Some characters, like emojis, take multiple bytes.
// 4. We convert the string to []rune so each rune is one element.
// 5. We index runes[3] to get the fourth rune.
// 6. We print it with %c to show the character, not the number.
func main() {
	// Step 1: Define the message with male and female emojis.
	message := "Hi 👨 and 👩"

	// Step 2: Convert string to slice of runes.
	runes := []rune(message)

	// Step 3: Get the fourth rune (index 3).
	fmt.Printf("%c\n", runes[3])
}
