package main

import "fmt"

// main is the CLI entry point for the todo app.
//
// Current behavior: runs a hardcoded command ("list").
// TODO: replace with os.Args / flag parsing so the user can pass
//       add, list, done, delete on the command line.
func main() {
	var list TaskList

	// Seed some tasks so the app isn't empty on first run.
	list.Add("Buy milk")
	list.Add("Walk the dog")
	list.Add("Call mom")

	command := "list"

	switch command {
	case "list":
		list.List()
	case "add":
		t := list.Add("Pay rent")
		fmt.Println("added", t.ID)
	case "done":
		if err := list.Complete(2); err != nil {
			fmt.Println(err)
		}
		list.List()
	case "delete":
		if err := list.Delete(1); err != nil {
			fmt.Println(err)
		}
		list.List()
	default:
		fmt.Println("unknown command")
	}
}
