package main

import "fmt"

// Exercise 12.1:
// Create a function that launches three goroutines that communicate using a
// channel. The first two goroutines each write 10 numbers to the channel.
// The third goroutine reads all the numbers from the channel and prints them.
// The function should exit when all values have been printed.
// Make sure that none of the goroutines leak.

// TODO: Write a writer function that sends 10 numbers to a channel
// TODO: Write a reader function that reads from the channel and prints values
// TODO: In main, create channel, launch 2 writers + 1 reader, close channel when done, wait for all

func main() {
	// TODO: Implement three-goroutine pattern
	fmt.Println("TODO: implement ex1-three-goroutines")
}
