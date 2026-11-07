package main

import "fmt"

// Exercise 9.3:
// Rather than returning the first error found, return back a single error that
// contains all errors discovered during validation. Update the code in main to
// properly report multiple errors.

// TODO: Write ValidateForm(username, email, age string) error
// TODO: Collect all validation errors into a slice
// TODO: Use errors.Join to return a single combined error

func main() {
	// TODO: Call ValidateForm with all-empty fields
	// TODO: Print the combined error
	// TODO: Call ValidateForm with all-valid fields
	fmt.Println("TODO: implement ex3-multiple-errors")
}
