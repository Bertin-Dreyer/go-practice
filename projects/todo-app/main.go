package main

import (
	"fmt"
	"time"
)

// TODO: implement the CLI commands: add, list, done, delete
// TODO: manage a slice of Task in memory
// TODO: persist to todos.json after Ch13

func main() {
	tasks := []Task{
		{ID: 1, Title: "Buy milk", Description: "", Done: false, CreatedAt: time.Now()},
		{ID: 2, Title: "Walk the dog", Description: "", Done: true, CreatedAt: time.Now()},
		{ID: 3, Title: "Call mom", Description: "", Done: false, CreatedAt: time.Now()},
	}

	command := "list"

	switch command {
	case "list":
		for _, t := range tasks {
			if t.Done {
				fmt.Println(t.Title, "[X]")
				continue
			}
			fmt.Println(t.Title, "[ ]")
		}
	case "add":
		fmt.Println("Adding...")
	default:
		fmt.Println("Unknown Command")
	}
}
