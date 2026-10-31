package main

import (
	"fmt"
	"os"
)

// Exercise 5.2:
// Write a function called fileLen that has an input parameter of type string and
// returns an int and an error. The function takes in a filename and returns the
// number of bytes in the file. If there is an error reading the file, return the error.
// Use defer to make sure the file is closed properly.
func main() {
	length, err := fileLen("test.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("File length:", length, "bytes")
}

func fileLen(input string) (i int, err error) {
	f, err := os.Open(input)
	if err != nil {
		return 0, err
	}
	defer f.Close()
	info, err := f.Stat()
	if err != nil {
		return 0, err
	}
	return int(info.Size()), nil
}
