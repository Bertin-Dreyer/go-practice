package main

import "fmt"

// Exercise 6.1:
// Create a struct named Person with three fields: FirstName and LastName of
// type string and Age of type int. Write a function called MakePerson that
// takes in firstName, lastName, and age and returns a Person. Write a second
// function MakePersonPointer that takes in firstName, lastName, and age and
// returns a *Person. Call both from main.
// Person holds one person's info.
//   FirstName - first name
//   LastName  - last name
//   Age       - age in years
type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	// MakePerson returns a Person VALUE (a copy).
	p1 := MakePerson("Bertin", "Dreyer", 28)
	fmt.Println("MakePerson struct:", p1.FirstName, p1.LastName, p1.Age)

	// MakePersonPointer returns a *Person POINTER (address of the struct).
	// Go auto-dereferences when you access fields via p2.FirstName.
	p2 := MakePersonPointer("Bertin", "Dreyer", 28)
	fmt.Println("MakePersonPointer:", p2.FirstName, p2.LastName, p2.Age)

	// KEY DIFFERENCE:
	//   Value  = the whole struct is copied when passed/returned
	//   Pointer = only the address (8 bytes) is copied; the struct lives once
}

func MakePerson(firstname, lastname string, age int) Person {
	return Person{
		FirstName: firstname,
		LastName:  lastname,
		Age:       age,
	}
}

func MakePersonPointer(firstname, lastname string, age int) *Person {
	return &Person{
		FirstName: firstname,
		LastName:  lastname,
		Age:       age,
	}
}
