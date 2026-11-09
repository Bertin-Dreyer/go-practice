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
	Tasks  []Task
	nextID int
}

// TODO: add methods on Task, e.g. String()

func (t *Task) Complete() {
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

func (tl *TaskList) Add(title string) Task {
	// TODO: bump nextID, build a Task with that ID, the title and time.Now(),
	// append it to tl.Tasks, and return it
	tl.nextID++ // 0 -> 1 on the first call
	t := Task{ID: tl.nextID, Title: title, CreatedAt: time.Now()}
	tl.Tasks = append(tl.Tasks, t)
	return t
}

func (tl *TaskList) Complete(id int) error {
	// TODO: find the task by ID and call Complete() on tl.Tasks[i]
	// (use the index, not the loop copy), or return a "not found" error

	for i, t := range tl.Tasks {
		if t.ID == id {
			tl.Tasks[i].Complete()
			return nil
		}
	}
	return fmt.Errorf("task %d not found", id)
}

func (tl *TaskList) List() {
	for _, t := range tl.Tasks {
		mark := "[ ]"
		if t.Done {
			mark = "[x]"
		}
		fmt.Println(mark, t.ID, t.Title)
	}
}
