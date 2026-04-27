package task

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// JSONStore is a "Concrete Implementation" of the Store interface.
// It uses a local file to keep our data even after the program closes.
type JSONStore struct {
	filePath string // The path to the .json file (e.g., /home/user/tasks.json)
}

// NewJSONStore is a "Constructor" function. It sets up a new store.
func NewJSONStore(filePath string) *JSONStore {
	return &JSONStore{filePath: filePath}
}

// Load reads the JSON file and turns it back into Go code (Slices of Tasks).
func (s *JSONStore) Load() ([]Task, error) {
	// 1. Check if the file exists.
	if _, err := os.Stat(s.filePath); os.IsNotExist(err) {
		// If it doesn't exist, just return an empty list of tasks.
		return []Task{}, nil
	}

	// 2. Read the raw bytes from the file.
	data, err := os.ReadFile(s.filePath)
	if err != nil {
		// fmt.Errorf with %w "wraps" the error, adding context so we know WHERE it failed.
		return nil, fmt.Errorf("loading tasks: reading file: %w", err)
	}

	// 3. "Unmarshal" (decode) the JSON bytes into a Go Slice ([]Task).
	var tasks []Task
	if err := json.Unmarshal(data, &tasks); err != nil {
		return nil, fmt.Errorf("loading tasks: unmarshaling json: %w", err)
	}

	return tasks, nil
}

// Save writes our tasks to the file. We use an "Atomic Write" pattern here.
func (s *JSONStore) Save(tasks []Task) error {
	// 1. Convert Go code into JSON bytes ("Marshal").
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return fmt.Errorf("saving tasks: marshaling json: %w", err)
	}

	// 2. ATOMIC WRITE PATTERN:
	// Instead of writing directly to "tasks.json", we write to a temporary file first.
	// Why? If the computer crashes or loses power while writing, "tasks.json" 
	// might become corrupted and unreadable. By writing to a temp file and then 
	// RENAMING it, the operation becomes "Atomic" (it either succeeds completely 
	// or doesn't happen at all).
	
	dir := filepath.Dir(s.filePath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("saving tasks: creating directory: %w", err)
	}

	// Create a temp file
	tmpFile, err := os.CreateTemp(dir, "tasks.*.json.tmp")
	if err != nil {
		return fmt.Errorf("saving tasks: creating temp file: %w", err)
	}
	defer os.Remove(tmpFile.Name()) // Clean up if we fail

	// Write the data to temp file
	if _, err := tmpFile.Write(data); err != nil {
		tmpFile.Close()
		return fmt.Errorf("saving tasks: writing to temp file: %w", err)
	}

	// Ensure bytes are actually written to the physical disk (not just cache)
	if err := tmpFile.Sync(); err != nil {
		tmpFile.Close()
		return fmt.Errorf("saving tasks: syncing temp file: %w", err)
	}

	if err := tmpFile.Close(); err != nil {
		return fmt.Errorf("saving tasks: closing temp file: %w", err)
	}

	// Final Step: Overwrite the real file with our finished temp file
	if err := os.Rename(tmpFile.Name(), s.filePath); err != nil {
		return fmt.Errorf("saving tasks: renaming temp file: %w", err)
	}

	// Set permissions: 0600 means only the owner can read/write this file.
	// This is a basic security practice for sensitive data.
	if err := os.Chmod(s.filePath, 0600); err != nil {
		return fmt.Errorf("saving tasks: setting file permissions: %w", err)
	}

	return nil
}

// Add appends a new task to our list.
func (s *JSONStore) Add(task Task) error {
	tasks, err := s.Load()
	if err != nil {
		return err
	}

	// ID Generation: Find the highest current ID and add 1.
	maxID := 0
	for _, t := range tasks {
		if t.ID > maxID {
			maxID = t.ID
		}
	}
	task.ID = maxID + 1

	tasks = append(tasks, task)
	return s.Save(tasks)
}

// Update replaces an existing task.
func (s *JSONStore) Update(task Task) error {
	tasks, err := s.Load()
	if err != nil {
		return err
	}

	found := false
	for i, t := range tasks {
		if t.ID == task.ID {
			tasks[i] = task // Replace the old task with the new one
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("task with ID %d not found", task.ID)
	}

	return s.Save(tasks)
}

// Delete removes a task by ID.
func (s *JSONStore) Delete(id int) error {
	tasks, err := s.Load()
	if err != nil {
		return err
	}

	// To delete in Go, we create a new list and COPY everything EXCEPT the one we delete.
	newTasks := make([]Task, 0, len(tasks))
	found := false
	for _, t := range tasks {
		if t.ID == id {
			found = true
			continue // Skip this one (delete)
		}
		newTasks = append(newTasks, t)
	}

	if !found {
		return fmt.Errorf("task with ID %d not found", id)
	}

	return s.Save(newTasks)
}
