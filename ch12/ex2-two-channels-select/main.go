package main

import "fmt"

// Exercise 12.2:
// Create a function that launches two goroutines. Each goroutine writes 10
// numbers to its own channel. Use a for-select loop to read from both channels,
// printing out the number and the goroutine that wrote the value.
// Make sure that your function exits after all values are read and that none
// of your goroutines leak.

// TODO: Create two channels
// TODO: Launch two goroutines, each writing 10 numbers to its channel
// TODO: Use for-select to read from both until both are closed
// TODO: Use sync.WaitGroup to wait for writers, then close channels

func main() {
	// TODO: Implement two-channel select pattern
	fmt.Println("TODO: implement ex2-two-channels-select")
}
