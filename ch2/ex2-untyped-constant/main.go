package main

import "fmt"

// Exercise 2.2:
// Declare a constant called value that can be assigned to both an integer
// and a floating-point variable.
// Assign it to an integer called i and a floating-point variable called f.
// Print out i and f.
// what's happening: value is an untyped constant (20). Go figures out its type
// based on context — int when assigned to i, float64 when assigned to f.
// this works because untyped constants don't have a fixed type until used.
func main() {
	const value = 20

	i := int(value)
	f := float64(value)

	fmt.Printf("i is an int with value of %d\nf is a float64 with value of %f\n", i, f)
}
