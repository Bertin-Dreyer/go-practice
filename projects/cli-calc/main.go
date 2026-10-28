package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

// opFunc is the signature for a binary integer operation.
type opFunc func(int, int) int

// add returns the sum of a and b.
func add(a, b int) int {
	return a + b
}

// sub returns the difference of a and b.
func sub(a, b int) int {
	return a - b
}

// mul returns the product of a and b.
func mul(a, b int) int {
	return a * b
}

// div returns the quotient of a and b, or an error if b is zero.
func div(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

// opMap maps operator symbols to their operation functions.
var opMap = map[string]opFunc{
	"+": add,
	"-": sub,
	"*": mul,
}

func main() {
	if len(os.Args) != 4 {
		fmt.Println("usage: go run main.go <num1> <op> <num2>")
		os.Exit(1)
	}

	a, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("invalid first number:", err)
		os.Exit(1)
	}

	op := os.Args[2]

	b, err := strconv.Atoi(os.Args[3])
	if err != nil {
		fmt.Println("invalid second number:", err)
		os.Exit(1)
	}

	var result int

	// Look up the operation. Addition, subtraction, and multiplication
	// return just an int. Division returns (int, error) so we handle it
	// separately.
	if op == "/" {
		result, err = div(a, b)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	} else if opFunc, ok := opMap[op]; ok {
		result = opFunc(a, b)
	} else {
		fmt.Println("unsupported operator:", op)
		os.Exit(1)
	}

	fmt.Println(result)
}
