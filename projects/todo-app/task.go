package main

import (
	"fmt"
	"slices"
	"time"
)

// Task represents a single to-do item.
type Task struct {
	ID          int
	Title       string
	Description string
	Done        bool
	CreatedAt   time.Time
}

type TaskList struct {
	Tasks []Task
}

// TODO: add methods on Task, e.g. String()

func (t *Task) Complete(){
	t.Done = true
}

func (tl *TaskList) Delete(id int) error {
	for i, t := range tl.Tasks {
		if t.ID == id {
			tl.Tasks = slices.Delete(tl.Tasks, i, i+1)
			return nil
		}
	}
	return fmt.Errorf("task %d not found", id)
}
