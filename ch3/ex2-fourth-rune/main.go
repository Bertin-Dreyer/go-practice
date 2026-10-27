package main

import "fmt"

// Exercise 3.2:
// Define a string variable called message with the value "Hi and ".
// Print the fourth rune in it as a character, not a number.
// what's happening: strings in Go are read-only slices of bytes. indexing into
// a string gives you bytes, but runes can be multiple bytes. to get runes we
// convert to []rune first, then index to get the fourth rune and print it as
// a character using %c.
func main() {
}
