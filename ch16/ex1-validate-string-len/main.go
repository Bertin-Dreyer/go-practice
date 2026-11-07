package main

import "fmt"

// Exercise 16.1:
// Use reflection to create a string-length validator for struct fields.
// Write ValidateStringLength(struct) error.
// Returns error if any string field has a "minStrlen" struct tag and value
// length is less than the tag value. Non-string fields and string fields without
// the tag are ignored. Use errors.Join for all invalid fields.
// Validate that a struct was passed in. Return nil if all fields valid.

// TODO: Write ValidateStringLength(v any) error
// TODO: Use reflect.ValueOf and reflect.TypeOf to inspect the struct
// TODO: Iterate fields, check Kind == String and Tag.Get("minStrlen")
// TODO: Compare len(value) against parsed tag int
// TODO: Use errors.Join to combine errors

func main() {
	fmt.Println("TODO: implement ex1-validate-string-len")
}
