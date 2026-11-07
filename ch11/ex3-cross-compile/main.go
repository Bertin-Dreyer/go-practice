package main

import "fmt"

// Exercise 11.3:
// Cross-compile your program for ARM64 on Windows. If you are using an ARM64
// Windows computer, cross-compile for AMD64 on Linux.
//
// Build commands:
//   GOOS=linux GOARCH=amd64 go build .
//   GOOS=windows GOARCH=arm64 go build .

func main() {
	fmt.Println("Cross-compilation example")
	fmt.Println("Build with: GOOS=linux GOARCH=amd64 go build .")
	fmt.Println("Or:         GOOS=windows GOARCH=arm64 go build .")
}
