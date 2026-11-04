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
	// The caller must reassign: append may return a new slice header.
	fmt.Println("Before GrowSlice:", words)
	words = GrowSlice(words, "date")
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

func GrowSlice(slice []string, s string) []string {
	// A slice header holds: pointer to array, length, capacity.
	// When passed to a function, the header is COPIED. Even if append
	// does NOT reallocate, the caller's header still has the OLD length,
	// so the caller never sees the appended element. If append DOES
	// reallocate, the local header points to a new array while the
	// caller's still points to the old array. Either way, the caller
	// must reassign: words = GrowSlice(words, "date")
	slice = append(slice, s)
	fmt.Println("Inside GrowSlice:", slice)
	return slice
}
