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
func main() {
	result, err := add(2, 1)
	fmt.Println("addition: ", result)
	result, err = sub(2, 1)
	fmt.Println("subtraction: ", result)
	result, err = mult(2, 1)
	fmt.Println("multiplication: ", result)
	result, err = div(2, 1)
	fmt.Println("divide: ", result)
	result, err = div(0, 2)
	fmt.Println("divide: ", result)
	result, err = div(2, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("divide: ", result)
	}
}

func add(i, j int) (output int, err error) {
	return i + j, nil
}

func sub(i, j int) (output int, err error) {
	return i - j, nil
}

func mult(i, j int) (output int, err error) {
	return i * j, nil
}
func div(i, j int) (output int, err error) {
	if j == 0 {
		return 0, errors.New("division by zero")
	} else {
		return i / j, nil
	}
}
