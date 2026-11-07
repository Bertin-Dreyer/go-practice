package main

import (
	"fmt"
	_ "embed"
)

// Exercise 11.1:
// Embed UDHR text files into package-level variables. Program takes a language
// name as command-line parameter and prints the UDHR in that language.

// TODO: Use //go:embed directives to embed english_rights.txt, spanish_rights.txt, french_rights.txt
// TODO: Create a map[string]string of language -> text
// TODO: In main, read os.Args[1] for language, look up and print the text

func main() {
	// TODO: Implement language lookup from embedded files
	fmt.Println("TODO: implement ex1-embed-udhr")
}
