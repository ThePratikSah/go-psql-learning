package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/thepratiksah/nextask/internal/domain"
)

// TaskRepository manages persistent storage for tasks.
type TaskRepository struct{}

// NewTaskRepository returns a TaskRepository.
func NewTaskRepository() *TaskRepository {
	return &TaskRepository{}
}

func (r *TaskRepository) Create(ctx context.Context, task domain.Task) (domain.Task, error) {
	return task, nil // placeholder — implemented with sqlc in Phase 4
}

func (r *TaskRepository) FindByID(ctx context.Context, id uuid.UUID) (domain.Task, error) {
	return domain.Task{}, nil // placeholder
}

func (r *TaskRepository) ListByProject(ctx context.Context, projectID uuid.UUID) ([]domain.Task, error) {
	return nil, nil // placeholder
}
