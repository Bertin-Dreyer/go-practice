package main

import "fmt"

// Exercise 4.2:
// Loop over the slice you created in exercise 1. For each value:
//   a. If divisible by 2, print "Two!"
//   b. If divisible by 3, print "Three!"
//   c. If divisible by both 2 and 3, print "Six!" (don't print anything else)
//   d. Otherwise, print "Never mind"
// what's happening: we iterate over the slice with a for-range loop. for each
// value we check divisibility with the % operator. order matters: check the
// combined case (divisible by 6) before the individual cases so we only print
// one message per value.
func main() {
}
