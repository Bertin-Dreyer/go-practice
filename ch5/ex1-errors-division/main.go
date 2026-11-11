package main

import (
	"errors"
	"fmt"
)

// Exercise 5.1:
// The simple calculator program doesn't handle one error case: division by zero.
// Change the function signature for the math operations to return both an int and
// an error. In the div function, if the divisor is 0, return errors.New("division
// by zero") for the error. In all other cases, return nil. Adjust the main function
// to check for this error.
//
// How it works, step by step:
//  1. We write four math functions: add, sub, mult, div.
//  2. Each function returns two values: (int, error).
//  3. The int is the result. The error is nil if nothing went wrong.
//  4. In div, we check if the divisor is 0 before dividing.
//     If it is 0, we return 0 and an error.
//     Otherwise we return the result and nil.
//  5. In main, we call each function and capture both return values.
//  6. Before printing a result, we check if err != nil.
//     If there is an error, we print the error instead of the result.
func main() {
	// Step 5a: Call add(2, 1), capture result and error.
	result, err := add(2, 1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("addition:", result)
	}

	// Step 5b: Call sub(2, 1).
	result, err = sub(2, 1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("subtraction:", result)
	}

	// Step 5c: Call mult(2, 1).
	result, err = mult(2, 1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("multiplication:", result)
	}

	// Step 5d: Call div(2, 1) — should work.
	result, err = div(2, 1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("divide:", result)
	}

	// Step 5e: Call div(0, 2) — should work.
	result, err = div(0, 2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("divide:", result)
	}

	// Step 5f: Call div(2, 0) — should fail with division by zero.
	result, err = div(2, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("divide:", result)
	}
}

// add returns the sum of i and j, with no error.
func add(i, j int) (output int, err error) {
	return i + j, nil
}

// sub returns the difference of i and j, with no error.
func sub(i, j int) (output int, err error) {
	return i - j, nil
}

// mult returns the product of i and j, with no error.
func mult(i, j int) (output int, err error) {
	return i * j, nil
}

// div returns the quotient of i divided by j.
// If j is 0, it returns an error instead.
func div(i, j int) (output int, err error) {
	if j == 0 {
		return 0, errors.New("division by zero")
	}
	return i / j, nil
}
