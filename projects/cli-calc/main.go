package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

// opFunc matches all math operations: two ints in, int and error out.
type opFunc func(int, int) (int, error)

func main() {
	if len(os.Args) != 4 {
		fmt.Println("usage: go run main.go <num1> <op> <num2>")
		os.Exit(1)
	}

	a, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	op := os.Args[2]

	b, err := strconv.Atoi(os.Args[3])
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	var opMap = map[string]opFunc{
		"+": add,
		"-": sub,
		"*": mult,
		"/": div,
	}

	opFunc, ok := opMap[op]
	if !ok {
		fmt.Println("unsupported operator:", op)
		os.Exit(1)
	}

	result, err := opFunc(a, b)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(result)
}

func add(i, j int) (int, error) {
	return i + j, nil
}

func sub(i, j int) (int, error) {
	return i - j, nil
}

func mult(i, j int) (int, error) {
	return i * j, nil
}

func div(i, j int) (int, error) {
	if j == 0 {
		return 0, errors.New("division by zero")
	}
	return i / j, nil
}
