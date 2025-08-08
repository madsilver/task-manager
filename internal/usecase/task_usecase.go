package usecase

import (
	"context"
	"github.com/madsilver/task-manager/internal/domain"
)

type TaskUseCase interface {
	CreateTask(ctx context.Context, task *domain.Task) error
	GetAllTasks(ctx context.Context) ([]*domain.Task, error)
	GetTaskByID(ctx context.Context, id int64) (*domain.Task, error)
	UpdateTask(ctx context.Context, task *domain.Task) error
	DeleteTask(ctx context.Context, id int64) error
}

type taskUseCase struct {
	repository domain.TaskRepository
	notifier   domain.TaskNotify
}

func NewTaskUseCase(repository domain.TaskRepository, notifier domain.TaskNotify) TaskUseCase {
	uc := &taskUseCase{repository, notifier}
	go uc.notifier.Subscribe(context.Background())
	return uc
}

func (u *taskUseCase) CreateTask(ctx context.Context, task *domain.Task) error {
	err := u.repository.Create(ctx, task)
	if err == nil {
		go u.notifier.Send(ctx, task)
	}
	return err
}

func (u *taskUseCase) GetAllTasks(ctx context.Context) ([]*domain.Task, error) {
	return u.repository.GetAll(ctx)
}

func (u *taskUseCase) GetTaskByID(ctx context.Context, id int64) (*domain.Task, error) {
	return u.repository.GetByID(ctx, id)
}

func (u *taskUseCase) UpdateTask(ctx context.Context, task *domain.Task) error {
	return u.repository.Update(ctx, task)
}

func (u *taskUseCase) DeleteTask(ctx context.Context, id int64) error {
	return u.repository.Delete(ctx, id)
}
