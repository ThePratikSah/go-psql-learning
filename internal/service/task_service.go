package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/thepratiksah/nextask/internal/domain"
)

// TaskRepository defines the DB interface for tasks.
type TaskRepository interface {
	Create(ctx context.Context, params domain.Task) (domain.Task, error)
	FindByID(ctx context.Context, id uuid.UUID) (domain.Task, error)
	ListByProject(ctx context.Context, projectID uuid.UUID) ([]domain.Task, error)
}

// TaskService handles task business logic.
type TaskService struct {
	repo TaskRepository
}

// NewTaskService creates a new TaskService.
func NewTaskService(repo TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) FindByID(ctx context.Context, id uuid.UUID) (*domain.Task, error) {
	task, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return &task, nil
}

// Create validates input and persists a new task.
//
// NOTE: This is Go's version of "throw new Error()" in Node.js.
// Instead of exceptions, we return errors and wrap them with %w.
// The caller can then use errors.Is() or errors.As() to check specific errors.
func (s *TaskService) Create(ctx context.Context, title string, projectID uuid.UUID) (*domain.Task, error) {
	if title == "" {
		return nil, errors.New("createTask: title is required")
	}
	if projectID == uuid.Nil {
		return nil, errors.New("createTask: project ID is required")
	}

	task := domain.Task{
		ID:        uuid.New(),
		ProjectID: projectID,
		Title:     title,
		Status:    domain.TaskStatusTodo,
		Priority:  0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	result, err := s.repo.Create(ctx, task)
	if err != nil {
		return nil, fmt.Errorf("createTask: persist task: %w", err)
	}
	return &result, nil
}
