package main

import "time"

// TODO: implement the CLI commands: add, list, done, delete
// TODO: manage a slice of Task in memory
// TODO: persist to todos.json after Ch13

func main() {
	tasks := []Task{
		{ID: 1, Title: "Buy milk", Description: "", Done: false, CreatedAt: time.Now()},
		{ID: 2, Title: "Walk the dog", Description: "", Done: true, CreatedAt: time.Now()},
		{ID: 3, Title: "Call mom", Description: "", Done: false, CreatedAt: time.Now()},
	}
	_ = tasks
}

