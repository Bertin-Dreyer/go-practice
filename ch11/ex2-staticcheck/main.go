package main

import "fmt"

// Exercise 11.2:
// Use go install to install staticcheck. Run it against your program and fix
// any problems it finds.

// TODO: This file has intentional issues for staticcheck to find:
//   - unused variable
//   - unnecessary conversion
//   - redundant if-else branches
// Fix them after running staticcheck ./...

func main() {
	x := 10
	y := 20
	unused := "hello"

	if x > 5 {
		fmt.Println("x is big")
	} else if x > 5 {
		fmt.Println("x is big (duplicate)")
	} else {
		fmt.Println("x is small")
	}

	fmt.Println("unused:", unused)
	fmt.Println("x + y =", x+y)
}
