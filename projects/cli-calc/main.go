package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

// opFunc is the signature for any binary integer operation:
// two ints in, an int result and an error out.
type opFunc func(int, int) (int, error)

func main() {
	// Step 1: Check that the user passed exactly 3 arguments:
	// <num1> <operator> <num2>. os.Args[0] is the program name.
	if len(os.Args) != 4 {
		fmt.Fprintln(os.Stderr, "usage: go run . <num1> <op> <num2>")
		os.Exit(1)
	}

	// Step 2: Parse the first number from os.Args[1].
	a, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}

	// Step 3: Grab the operator string (+, -, *, /) from os.Args[2].
	op := os.Args[2]

	// Step 4: Parse the second number from os.Args[3].
	b, err := strconv.Atoi(os.Args[3])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}

	// Step 5: Map operator strings to their math functions.
	// The map value type is opFunc (func(int, int) (int, error)).
	var opMap = map[string]opFunc{
		"+": add,
		"-": sub,
		"*": mult,
		"/": div,
	}

	// Step 6: Look up the operator in the map.
	// comma-ok idiom: fn is the function, ok is true if found.
	fn, ok := opMap[op]
	if !ok {
		fmt.Fprintln(os.Stderr, "unsupported operator:", op)
		os.Exit(1)
	}

	// Step 7: Call the selected function with the two numbers.
	result, err := fn(a, b)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Step 8: Print the result to stdout.
	fmt.Println(result)
}

// add returns the sum of i and j.
func add(i, j int) (int, error) {
	return i + j, nil
}

// sub returns the difference of i and j.
func sub(i, j int) (int, error) {
	return i - j, nil
}

// mult returns the product of i and j.
func mult(i, j int) (int, error) {
	return i * j, nil
}

// div returns the quotient of i divided by j, or an error if j is zero.
func div(i, j int) (int, error) {
	if j == 0 {
		return 0, errors.New("division by zero")
	}
	return i / j, nil
}
