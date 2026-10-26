package main

import "fmt"

// Exercise 2.1:
// Declare an integer variable called i with the value 20.
// Assign i to a floating-point variable named f.
// Print out i and f.
// what's happening: i is an int (20). we convert it to float64 and store in f.
// ints and floats are different types in Go, so we need an explicit conversion.
// then we print both with Printf.
func main() {
	i := 20

	f := float64(i)

	fmt.Printf("i is an int with value of %d\nf is a float64 with value of %f\n", i, f)
}
