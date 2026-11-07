package main

import "fmt"

// Exercise 12.3:
// Write a function that builds a map[int]float64 where keys are 0..99999 and
// values are square roots. Use sync.OnceValue to cache the map.
// Use the cached value to look up square roots for every 1,000th number.

// TODO: Write buildSqrtMap() map[int]float64 that fills 0..99999 with math.Sqrt
// TODO: Use sync.OnceValue(buildSqrtMap) to create a cached getter
// TODO: Look up every 1,000th number and print results

func main() {
	// TODO: Implement sqrt map caching with sync.OnceValue
	fmt.Println("TODO: implement ex3-sqrt-cache")
}
