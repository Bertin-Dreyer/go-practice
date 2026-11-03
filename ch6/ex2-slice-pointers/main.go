package main

import "fmt"

// Exercise 6.2:
// Write two functions. The UpdateSlice function takes in a []string and a
// string. It sets the last position in the passed-in slice to the passed-in string.
// The GrowSlice function also takes in a []string and a string. It appends the
// string onto the slice. Call these functions from main. Print out the slice
// before each function is called and after each function is called.
func main() {
	// 1. Create our starting slice
	words := []string{"apple", "banana", "carrot"}

	// 2. Test UpdateSlice - changes an EXISTING element
	fmt.Println("Before UpdateSlice:", words)
	UpdateSlice(words, "cherry")
	fmt.Println("After UpdateSlice:", words)

	fmt.Println("-------------------")

	// 3. Test GrowSlice - appends a NEW element
	fmt.Println("Before GrowSlice:", words)
	GrowSlice(words, "date")
	fmt.Println("After GrowSlice:", words)
}

func UpdateSlice(slice []string, s string) {
	// A slice header holds: pointer to array, length, capacity.
	// When passed to a function, the header is COPIED, but the pointer still
	// points to the SAME underlying array. So mutating an index is visible
	// to the caller.
	slice[len(slice)-1] = s
	fmt.Println("Inside UpdateSlice:", slice)
}

func GrowSlice(slice []string, s string) {
	// append may reallocate if capacity is full. When it reallocates, the
	// local slice header now points to a NEW array. The caller's header still
	// points to the OLD array, so the caller does NOT see the new element.
	// That's why you must do: words = GrowSlice(words, "date")
	slice = append(slice, s)
	fmt.Println("Inside GrowSlice:", slice)
}
