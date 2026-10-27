package main

import "fmt"

// Exercise 3.3:
// Define a struct called Employee with three fields:
// firstName (string), lastName (string), id (int).
// Create three instances of this struct:
//   - first: struct literal style without field names
//   - second: struct literal style with field names
//   - third: var declaration, populate fields using dot notation
//
// Print out all three structs.
// what's happening: structs let us group related data into a single type.
// there are two literal styles (positional vs named) and var declarations.
// dot notation lets us set fields after declaration.
type Employee struct {
	firstName string
	lastName  string
	id        int
}

func main() {
	first := Employee{"John", "Doe", 0}
	second := Employee{
		firstName: "Bob",
		lastName:  "Dylan",
		id:        1,
	}
	third := Employee{}
	third.firstName = "Jack"
	third.lastName = "Black"
	third.id = 2

	fmt.Println(first, "\n", second, "\n", third)
}
