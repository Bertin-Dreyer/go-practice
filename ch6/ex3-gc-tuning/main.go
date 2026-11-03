// Exercise 6.3:
// Write a program that builds a []Person with 10,000,000 entries. See how long
// it takes to run. Change the value of GOGC and see how that affects the time.
package main

import (
	"fmt"
	"runtime"
	"time"
)

// Person matches the struct from Exercise 6.1.
type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	// Start the stopwatch so we can measure how long the build takes.
	start := time.Now()

	// Slice that will hold 10 million Person entries.
	// Without make(..., cap), append reallocates and copies many times as
	// the slice grows. Each reallocation triggers GC pressure.
	var people []Person

	// Build the slice one entry at a time.
	for i := 0; i < 10_000_000; i++ {
		people = append(people, Person{
			FirstName: "John",
			LastName:  "Doe",
			Age:       30,
		})
	}

	// Stop the stopwatch and report elapsed time.
	duration := time.Since(start)
	fmt.Printf("Time taken: %v\n", duration)

	// Ask the runtime for GC stats so you can see how many collections ran.
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
	fmt.Printf("GC runs: %d\n", memStats.NumGC)

	// TO TEST GOGC (run from terminal):
	//   GOGC=100 go run main.go   (default: GC when heap grows 100%)
	//   GOGC=500  go run main.go  (less frequent GC, usually faster)
	//   GOGC=10   go run main.go  (more frequent GC, usually slower)
	// GOGC = percent of heap growth allowed before triggering a GC cycle.
	// Try it with and without pre-allocating capacity:
	//   people := make([]Person, 0, 10_000_000)
}

// EXPECTED OUTPUT (format, times vary by machine):
//   Time taken: 1.xxxxxxxs
//   GC runs: N
//
// Actual results (this machine, no pre-allocation):
//   GOGC=100 (default): Time taken: 1.629412s   GC runs: 15
//   GOGC=500:           Time taken: 1.1985908s  GC runs: 4   <-- fastest, fewest GCs
//   GOGC=10:            Time taken: 1.8114521s  GC runs: 30  <-- slowest, most GCs
//
// Takeaway: Higher GOGC = GC runs less often = usually faster, but uses more
//           memory between collections. Lower GOGC = GC runs more often =
//           usually slower, but keeps memory lower.
}
