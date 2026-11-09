package main

import (
	"fmt"
	"slices"
	"time"
)

// ---------------------------------------------------------------------------
// Task — a single to-do item
// ---------------------------------------------------------------------------

// Task holds everything about one to-do item.
//   ID          - unique number, auto-assigned when added
//   Title       - the short description of what to do
//   Description - optional longer notes (can be empty)
//   Done        - true once the task is finished
//   CreatedAt   - when the task was added
type Task struct {
	ID          int
	Title       string
	Description string
	Done        bool
	CreatedAt   time.Time
}

// Complete marks this task as done.
// It uses a pointer receiver (*Task) so the change sticks to the original
// task inside the TaskList slice.
func (t *Task) Complete() {
	t.Done = true
}

// ---------------------------------------------------------------------------
// TaskList — a collection of tasks
// ---------------------------------------------------------------------------

// TaskList holds the slice of all tasks plus a counter for the next ID.
type TaskList struct {
	Tasks  []Task
	nextID int
}

// Add creates a new task with the given title, assigns it the next ID,
// timestamps it, appends it to the slice, and returns the new Task.
//
// Returns:
//   The newly created Task (with its ID, Title, CreatedAt filled in).
func (tl *TaskList) Add(title string) Task {
	tl.nextID++ // 0 -> 1 on the first call, so IDs start at 1
	t := Task{
		ID:         tl.nextID,
		Title:      title,
		CreatedAt:  time.Now(),
	}
	tl.Tasks = append(tl.Tasks, t)
	return t
}

// Complete finds the task with the matching ID and marks it done.
//
// It loops over tl.Tasks by index so it can mutate tl.Tasks[i] directly
// (using the loop copy `t` would NOT change the original slice element).
//
// Returns nil on success, or an error if no task with that ID exists.
func (tl *TaskList) Complete(id int) error {
	for i, t := range tl.Tasks {
		if t.ID == id {
			tl.Tasks[i].Complete() // mutate the slice element by index
			return nil
		}
	}
	return fmt.Errorf("task %d not found", id)
}

// Delete removes the task with the matching ID from the slice.
//
// Uses slices.Delete to remove the element at index i.
//
// Returns nil on success, or an error if no task with that ID exists.
func (tl *TaskList) Delete(id int) error {
	for i, t := range tl.Tasks {
		if t.ID == id {
			tl.Tasks = slices.Delete(tl.Tasks, i, i+1)
			return nil
		}
	}
	return fmt.Errorf("task %d not found", id)
}

// List prints every task with a done/pending marker, its ID, and its title.
//   [ ] 1 Buy milk      (not done)
//   [x] 2 Walk the dog  (done)
func (tl *TaskList) List() {
	for _, t := range tl.Tasks {
		mark := "[ ]"
		if t.Done {
			mark = "[x]"
		}
		fmt.Println(mark, t.ID, t.Title)
	}
}
