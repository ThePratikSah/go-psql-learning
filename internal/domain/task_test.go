package domain

import (
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestTask_IsComplete(t *testing.T) {
	tests := []struct {
		name   string
		status TaskStatus
		want   bool
	}{
		{"done task is complete", TaskStatusDone, true},
		{"todo task is not complete", TaskStatusTodo, false},
		{"in progress task is not complete", TaskStatusInProgress, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			task := Task{Status: tt.status}
			if got := task.IsComplete(); got != tt.want {
				t.Errorf("IsComplete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTask_CanTransitionTo(t *testing.T) {
	tests := []struct {
		name      string
		current   TaskStatus
		target    TaskStatus
		want      bool
	}{
		{"todo to in_progress", TaskStatusTodo, TaskStatusInProgress, true},
		{"todo to done (skip)", TaskStatusTodo, TaskStatusDone, false},
		{"in_progress to done", TaskStatusInProgress, TaskStatusDone, true},
		{"in_progress to todo (backwards)", TaskStatusInProgress, TaskStatusTodo, false},
		{"done stays done", TaskStatusDone, TaskStatusDone, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			task := Task{Status: tt.current}
			if got := task.CanTransitionTo(tt.target); got != tt.want {
				t.Errorf("CanTransitionTo(%v) = %v, want %v", tt.target, got, tt.want)
			}
		})
	}
}

func TestTask_IsValid(t *testing.T) {
	tests := []struct {
		name    string
		task    Task
		wantOK  bool
		wantMsg string
	}{
		{
			name:    "valid task",
			task:    Task{Title: "Fix bug", ProjectID: uuid.New(), Priority: 3},
			wantOK:  true,
		},
		{
			name:    "missing title",
			task:    Task{ProjectID: uuid.New(), Priority: 3},
			wantOK:  false,
			wantMsg: "title is required",
		},
		{
			name:    "missing project ID",
			task:    Task{Title: "Fix bug", Priority: 3},
			wantOK:  false,
			wantMsg: "project ID is required",
		},
		{
			name:    "negative priority",
			task:    Task{Title: "Fix bug", ProjectID: uuid.New(), Priority: -1},
			wantOK:  false,
			wantMsg: "priority must be between 0 and 5",
		},
		{
			name:    "priority too high",
			task:    Task{Title: "Fix bug", ProjectID: uuid.New(), Priority: 10},
			wantOK:  false,
			wantMsg: "priority must be between 0 and 5",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ok, msg := tt.task.IsValid()
			if ok != tt.wantOK {
				t.Errorf("IsValid() = %v, want %v", ok, tt.wantOK)
			}
			if !ok && msg != tt.wantMsg {
				t.Errorf("IsValid() msg = %q, want %q", msg, tt.wantMsg)
			}
		})
	}
}

// task_test.go uses time.Now() — suppress unused import warning
var _ = time.Now
