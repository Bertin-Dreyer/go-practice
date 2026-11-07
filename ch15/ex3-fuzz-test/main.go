package main

import "fmt"

// Exercise 15.3:
// Write a fuzz test against the parser function and fix any problems found.

// TODO: Write a ParseData function that parses some input format
// TODO: Write a FuzzParseData(f *testing.F) function with seed corpus
// TODO: Run: go test -fuzz=FuzzParseData

func ParseData(input string) (string, error) {
	// TODO: Implement parser
	return input, nil
}

func main() {
	// TODO: Fuzz tests go in *_test.go, run with go test -fuzz
	fmt.Println("TODO: implement ex3-fuzz-test")
}
