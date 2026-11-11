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
//
// How it works, step by step:
//  1. We call os.Open(filename) to open the file.
//     This returns a *os.File and an error.
//  2. If there is an error opening the file, we return 0 and the error.
//  3. We use defer f.Close() to guarantee the file is closed when the
//     function returns, even if an error happens later.
//  4. We call f.Stat() to get file information.
//     This returns a FileInfo and an error.
//  5. If Stat fails, we return 0 and the error.
//  6. We read the file size with info.Size() and convert it to int.
//  7. We return the size and nil (no error).
func main() {
	// Step 1: Call fileLen with a filename.
	length, err := fileLen("test.txt")

	// Step 2: Check if there was an error.
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Step 3: Print the file size in bytes.
	fmt.Println("File length:", length, "bytes")
}

// fileLen opens a file and returns its size in bytes.
// input: path to the file.
// returns: file size in bytes, or an error if something goes wrong.
func fileLen(input string) (i int, err error) {
	// Step 1: Open the file.
	f, err := os.Open(input)
	if err != nil {
		return 0, err
	}

	// Step 2: Guarantee the file is closed when this function returns.
	defer f.Close()

	// Step 3: Get file information.
	info, err := f.Stat()
	if err != nil {
		return 0, err
	}

	// Step 4: Return the file size.
	return int(info.Size()), nil
}
