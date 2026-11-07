package main

import "fmt"

// Add returns the sum of a and b.
//
// See https://www.mathsisfun.com/numbers/addition.html for more about addition.
func Add(a, b int) int {
	return a + b
}

func main() {
	fmt.Println("Add(2, 3) =", Add(2, 3))
}
