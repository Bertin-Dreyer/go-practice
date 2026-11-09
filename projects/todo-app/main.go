package main

import "fmt"

func main() {
	var list TaskList
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
