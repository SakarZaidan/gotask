/*
Package task handles the core "business logic" of our application.
In backend engineering, we separate how data is STORED from how data behaves.
*/
package task

import (
	"fmt"
	"time"
)

// TaskStatus is a custom type based on an integer (int).
// In Go, creating a custom type makes our code "type-safe" (we can't accidentally 
// use a random number where a status is expected).
type TaskStatus int

// "const" defines values that never change.
// "iota" is a special Go keyword used in const blocks to automatically 
// increment numbers (0, 1, 2...). This is called an "Enum" in other languages.
const (
	StatusPending    TaskStatus = iota // 0
	StatusInProgress                   // 1
	StatusDone                         // 2
)

// This is a "Method". It's a function attached to the TaskStatus type.
// It allows us to turn the number (0, 1, 2) into a human-readable string.
func (s TaskStatus) String() string {
	switch s {
	case StatusPending:
		return "pending"
	case StatusInProgress:
		return "in-progress"
	case StatusDone:
		return "done"
	default:
		return "unknown"
	}
}

// "type Task struct" defines a custom data structure (like a template).
// It groups different pieces of information together.
type Task struct {
	// The strings inside `json:"..."` are called "Struct Tags".
	// They tell the computer: "When you turn this into a JSON file, 
	// call this field 'id' instead of 'ID'".
	ID          int        `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description,omitempty"` // omitempty: hide if empty
	Status      TaskStatus `json:"status"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	// The * means this is a "Pointer". It allows the value to be "nil" (empty).
	// Not every task has a due date, so we use a pointer.
	DueDate     *time.Time `json:"due_date,omitempty"`
}

// Complete is a method with a "Pointer Receiver" (t *Task).
// The * means we are modifying the ACTUAL task in memory, 
// not a copy of it.
func (t *Task) Complete() {
	t.Status = StatusDone
	t.UpdatedAt = time.Now()
}

// String defines how a Task looks when printed.
func (t *Task) String() string {
	return fmt.Sprintf("[%d] %s (%s)", t.ID, t.Title, t.Status)
}
