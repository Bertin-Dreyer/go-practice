package main

import "fmt"

// Exercise 3.1:
// Define a slice of strings called greetings with these values:
// "Hello", "Hola", "नमस्कार", "こんにちは", "Привіт".
// Create a subslice containing the first two values.
// Create a second subslice with the second, third, and fourth values.
// Create a third subslice with the fourth and fifth values.
// Print out all four slices.
//
// How it works, step by step:
//  1. We create a slice called greetings with five string values.
//  2. We create sub1 using greetings[:2] — first two values.
//  3. We create sub2 using greetings[1:4] — second, third, fourth values.
//  4. We create sub3 using greetings[3:5] — fourth and fifth values.
//  5. We print all four slices.
//     Subslices share the same underlying array as greetings.
func main() {
	// Step 1: Define the slice with five greetings.
	var greetings = []string{"Hello", "Hola", "नमस्कार", "こんにちは", "Привіт"}

	// Step 2: First two values (indices 0 and 1).
	sub1 := greetings[:2]

	// Step 3: Second, third, and fourth values (indices 1, 2, 3).
	sub2 := greetings[1:4]

	// Step 4: Fourth and fifth values (indices 3 and 4).
	sub3 := greetings[3:5]

	// Step 5: Print all four slices.
	fmt.Println(greetings, "\n", sub1, "\n", sub2, "\n", sub3)
}
