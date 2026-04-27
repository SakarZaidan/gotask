package task

import (
	"os"
	"path/filepath"
	"testing"
)

func TestJSONStore(t *testing.T) {
	tmpDir := t.TempDir()
	dbPath := filepath.Join(tmpDir, "tasks.json")
	store := NewJSONStore(dbPath)

	// Test Add
	t.Run("Add", func(t *testing.T) {
		tk := Task{Title: "Task 1"}
		if err := store.Add(tk); err != nil {
			t.Fatalf("Failed to add task: %v", err)
		}

		tasks, err := store.Load()
		if err != nil {
			t.Fatalf("Failed to load tasks: %v", err)
		}

		if len(tasks) != 1 {
			t.Errorf("Expected 1 task, got %d", len(tasks))
		}

		if tasks[0].ID != 1 {
			t.Errorf("Expected ID 1, got %d", tasks[0].ID)
		}

		if tasks[0].Title != "Task 1" {
			t.Errorf("Expected Title 'Task 1', got %s", tasks[0].Title)
		}
	})

	// Test Update
	t.Run("Update", func(t *testing.T) {
		tasks, _ := store.Load()
		tk := tasks[0]
		tk.Title = "Updated Task"
		tk.Status = StatusDone

		if err := store.Update(tk); err != nil {
			t.Fatalf("Failed to update task: %v", err)
		}

		tasks, _ = store.Load()
		if tasks[0].Title != "Updated Task" {
			t.Errorf("Expected Title 'Updated Task', got %s", tasks[0].Title)
		}
		if tasks[0].Status != StatusDone {
			t.Errorf("Expected StatusDone, got %v", tasks[0].Status)
		}
	})

	// Test Delete
	t.Run("Delete", func(t *testing.T) {
		if err := store.Delete(1); err != nil {
			t.Fatalf("Failed to delete task: %v", err)
		}

		tasks, _ := store.Load()
		if len(tasks) != 0 {
			t.Errorf("Expected 0 tasks, got %d", len(tasks))
		}
	})

	// Test Security (permissions)
	t.Run("Permissions", func(t *testing.T) {
		store.Add(Task{Title: "Perm Test"})
		info, err := os.Stat(dbPath)
		if err != nil {
			t.Fatalf("Failed to stat file: %v", err)
		}
		
		// In some environments, we might need to check just the last 3 bits or so
		// But 0600 is what we set.
		expectedPerm := os.FileMode(0600)
		if info.Mode().Perm() != expectedPerm {
			t.Errorf("Expected permissions %v, got %v", expectedPerm, info.Mode().Perm())
		}
	})
}
