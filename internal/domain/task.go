package domain

import (
	"time"

	"github.com/google/uuid"
)

type TaskStatus string

const (
	TaskStatusTodo       TaskStatus = "todo"
	TaskStatusInProgress TaskStatus = "in_progress"
	TaskStatusDone       TaskStatus = "done"
)

type Task struct {
	ID          uuid.UUID  `json:"id"`
	ProjectID   uuid.UUID  `json:"project_id"`
	AssigneeID  *uuid.UUID `json:"assignee_id,omitempty"`
	Title       string     `json:"title"`
	Description string     `json:"description,omitempty"`
	Status      TaskStatus `json:"status"`
	Priority    int        `json:"priority"`
	DueDate     *time.Time `json:"due_date,omitempty"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

// IsComplete returns true if the task is done.
//
// NOTE: Value receiver (t Task) — we don't mutate, so no pointer needed.
// In Node.js this would be a class method with `this.status === 'done'`.
func (t Task) IsComplete() bool {
	return t.Status == TaskStatusDone
}

// CanTransitionTo checks if the task can move to the given status.
// Rules: todo → in_progress, in_progress → done. No backwards transitions.
//
// NOTE: Value receiver — only reading fields.
// This is Go's version of a TypeScript discriminated union guard.
func (t Task) CanTransitionTo(status TaskStatus) bool {
	switch t.Status {
	case TaskStatusTodo:
		return status == TaskStatusInProgress
	case TaskStatusInProgress:
		return status == TaskStatusDone
	case TaskStatusDone:
		return false
	}
	return false
}

// IsValid checks that required fields are set.
// Returns (false, reason) if invalid.
//
// NOTE: Go's "if err != nil" pattern — no exceptions, just return values.
// Compare to Node.js: throw new Error('title is required').
func (t Task) IsValid() (bool, string) {
	if t.Title == "" {
		return false, "title is required"
	}
	if t.Priority < 0 || t.Priority > 5 {
		return false, "priority must be between 0 and 5"
	}
	if t.ProjectID == uuid.Nil {
		return false, "project ID is required"
	}
	return true, ""
}
