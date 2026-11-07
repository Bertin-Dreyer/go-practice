package main

import "fmt"

// Exercise 10.3:
// Change Add to make it generic. Import golang.org/x/exp/constraints.
// Combine Integer and Float types into an interface called Number.
// Rewrite Add to take two parameters of type Number and return Number.
// Version your module again (backward-breaking change = v2.0.0).

// TODO: Import "golang.org/x/exp/constraints"
// TODO: Define Number interface combining constraints.Integer | constraints.Float
// TODO: Write AddGeneric[T Number](a, b T) T

func main() {
	// TODO: Test AddGeneric with int, float64, int64
	fmt.Println("TODO: implement ex3-generic-add-v2")
}
