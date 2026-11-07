package main

import (
	"fmt"
	"sync"
)

// Exercise 15.2:
// Use the race detector to find a concurrency problem and fix it.
// Run with: go test -race

// TODO: Write code with a data race (e.g., shared counter without mutex)
// TODO: Fix it with sync.Mutex or sync/atomic

var counter int

func increment() {
	// TODO: Add synchronization
	counter++
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			increment()
		}()
	}
	wg.Wait()
	fmt.Println("Counter:", counter) // should be 1000
}
