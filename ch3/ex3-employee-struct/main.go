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
//
// How it works, step by step:
//  1. We define the Employee struct with three fields.
//  2. We create first using positional literal: Employee{"John", "Doe", 0}.
//     Field names are not written — order matters.
//  3. We create second using named literal: Employee{firstName: ..., ...}.
//     Field names are written — order does not matter.
//  4. We create third using var declaration: var third Employee, then set
//     each field with dot notation (third.firstName = ...).
//  5. We print all three structs.
func main() {
	// Step 2: First instance — positional literal, no field names.
	first := Employee{"John", "Doe", 0}

	// Step 3: Second instance — named literal, field names included.
	second := Employee{
		firstName: "Bob",
		lastName:  "Dylan",
		id:        1,
	}

	// Step 4: Third instance — var declaration, then dot notation.
	third := Employee{}
	third.firstName = "Jack"
	third.lastName = "Black"
	third.id = 2

	// Step 5: Print all three structs.
	fmt.Println(first, "\n", second, "\n", third)
}

// Employee holds one person's info.
//
//	firstName - person's first name
//	lastName  - person's last name
//	id        - person's ID number
type Employee struct {
	firstName string
	lastName  string
	id        int
}
