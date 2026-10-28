package main

import "time"

// Task represents a single to-do item.
type Task struct {
	ID          int
	Title       string
	Description string
	Done        bool
	CreatedAt   time.Time
}

// TODO: add methods on Task, e.g. MarkDone(), String()
