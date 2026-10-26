package main

import "fmt"

// Exercise 2.3:
// Declare three variables: b of type byte, smallI of type int32, bigI of type uint64.
// Assign each variable the maximum legal value for its type.
// Add 1 to each variable.
// Print out their values.
// what's happening: we set each to its max, then add 1. they overflow:
// byte (uint8) max 255 + 1 = 0
// int32 max 2147483647 + 1 = -2147483648 (wraps to negative min)
// uint64 max 18446744073709551615 + 1 = 0
// this is integer overflow — fixed-width types wrap around when they exceed their range.
func main() {
	var b byte = 255
	var smallI int32 = 2147483647
	var bigI uint64 = 18446744073709551615

	b += 1
	smallI += 1
	bigI += 1

	fmt.Printf("b is %d\nsmallI is %d\nbigI is %d\n", b, smallI, bigI)
}
