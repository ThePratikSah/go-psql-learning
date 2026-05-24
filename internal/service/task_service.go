package service

import (
	"context"

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
