package main

import "fmt"

// Exercise 5.3:
// Write a function called prefixer that has an input parameter of type string
// and returns a function that has an input parameter of type string and returns a
// string. The returned function should prefix its input with the string passed into
// prefixer. Use the following main function to test prefixer:
//
//	helloPrefix := prefixer("Hello")
//	fmt.Println(helloPrefix("Bob")) // should print Hello Bob
//	fmt.Println(helloPrefix("Maria")) // should print Hello Maria
func main() {
	helloPrefix := prefixer("Hello")
	fmt.Println(helloPrefix("Bob"))   // should print Hello Bob
	fmt.Println(helloPrefix("Maria")) // should print Hello Maria
}

func prefixer(input string) func(inFunc string) (output string) {
	return func(s string) string {
		return input + " " + s
	}
}
