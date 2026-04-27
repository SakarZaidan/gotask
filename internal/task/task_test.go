package task

import (
	"testing"
)

func TestTaskComplete(t *testing.T) {
	task := Task{
		ID:     1,
		Title:  "Test Task",
		Status: StatusPending,
	}

	task.Complete()

	if task.Status != StatusDone {
		t.Errorf("Expected status to be StatusDone, got %v", task.Status)
	}

	if task.UpdatedAt.IsZero() {
		t.Error("Expected UpdatedAt to be set")
	}
}

func TestTaskStatusString(t *testing.T) {
	tests := []struct {
		status TaskStatus
		want   string
	}{
		{StatusPending, "pending"},
		{StatusInProgress, "in-progress"},
		{StatusDone, "done"},
		{TaskStatus(99), "unknown"},
	}

	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if got := tt.status.String(); got != tt.want {
				t.Errorf("TaskStatus.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
