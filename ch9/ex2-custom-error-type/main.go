package main

import "fmt"

// Exercise 9.2:
// Define a custom error type to represent an empty field error. This error should
// include the name of the empty Employee field. In main, use errors.As to check
// for this error. Print out a message that includes the field name.

// TODO: Define Employee struct with ID, FirstName, LastName, Email
// TODO: Define EmptyFieldError struct with Field string
// TODO: Implement Error() string on EmptyFieldError
// TODO: Write ValidateEmployee(e Employee) error returning EmptyFieldError for empty fields

func main() {
	// TODO: Create an Employee with an empty field
	// TODO: Call ValidateEmployee
	// TODO: Use errors.As to extract EmptyFieldError and print field name
	fmt.Println("TODO: implement ex2-custom-error-type")
}
