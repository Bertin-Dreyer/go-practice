package main

import "fmt"

// Exercise 8.2:
// Define a generic interface called Printable that matches a type that implements
// fmt.Stringer and has an underlying type of int or float64. Define types that
// meet this interface. Write a function that takes in a Printable and prints its
// value to the screen using fmt.Println.

// TODO: Define Printable[T any] interface with fmt.Stringer + ~int | ~float64
// TODO: Define MyInt type with String() method
// TODO: Define MyFloat type with String() method
// TODO: Write PrintValue[T Printable[T]](p T) that prints p.String()

func main() {
	// TODO: Create MyInt and MyFloat values, call PrintValue on each
	fmt.Println("TODO: implement ex2-printable-interface")
}
