package main

import (
	"fmt"
)

// Exercise 16.3:
// Copy the C code at ch16/tree/main/sample_code/mini_calc into your own module
// and use cgo to call it from a Go program.

// TODO: Create a mini_calc directory with a C file (e.g., calc.c) containing
//   int add(int a, int b) { return a + b; }
//   int multiply(int a, int b) { return a * b; }
// TODO: Write Go code using //cgo directives to link and call the C functions
// TODO: Import "C" and call C.add, C.multiply

// #cgo LDFLAGS: -L./mini_calc -lmini_calc
// #include <mini_calc.h>
// import "C"

func main() {
	// TODO: Call C.add(2, 3) and C.multiply(4, 5), print results
	fmt.Println("TODO: implement ex3-cgo-mini-calc")
}
