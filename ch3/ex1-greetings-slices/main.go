package main

import "fmt"

// Exercise 3.1:
// Define a slice of strings called greetings with these values:
// "Hello", "Hola", "नमस्कार", "こんにちは", "Привіт".
// Create a subslice containing the first two values.
// Create a second subslice with the second, third, and fourth values.
// Create a third subslice with the fourth and fifth values.
// Print out all four slices.
// what's happening: we define a slice of strings, then create several subslices
// using slice expressions (a[low:high]). subslices share the underlying array.
func main() {
	var greetings = []string{"Hello", "Hola", "नमस्कार", "こんにちは", "Привіт"}

	sub1 := greetings[:2]
	sub2 := greetings[1:4]
	sub3 := greetings[3:5]

	fmt.Println(greetings, "\n", sub1, "\n", sub2, "\n", sub3)
}
