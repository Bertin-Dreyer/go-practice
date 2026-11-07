package main

import "fmt"

// Exercise 9.1:
// Create a sentinel error to represent an invalid ID. In main, use errors.Is to
// check for the sentinel error, and print a message when it is found.

// TODO: Define ErrInvalidID sentinel error variable
// TODO: Write LookupUser(id int) (string, error) that returns ErrInvalidID for bad IDs

func main() {
	// TODO: Call LookupUser with valid and invalid IDs
	// TODO: Use errors.Is to check for ErrInvalidID
	fmt.Println("TODO: implement ex1-sentinel-error")
}
