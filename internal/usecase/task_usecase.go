package usecase

import (
	"context"
	"github.com/madsilver/task-manager/internal/domain"
	"time"
)

type TaskUseCase interface {
	CreateTask(ctx context.Context, task *domain.Task) error
	GetAllTasks(ctx context.Context) ([]*domain.Task, error)
	GetTaskByID(ctx context.Context, id int64) (*domain.Task, error)
	UpdateTask(ctx context.Context, task *domain.Task) error
	DeleteTask(ctx context.Context, id int64) error
}

const (
	CreateOperation = "CREATE"
	UpdateOperation = "UPDATE"
	DeleteOperation = "DELETE"
	ReadOperation   = "READ"
)

type taskUseCase struct {
	taskRepository  domain.TaskRepository
	auditRepository domain.TaskAuditRepository
	notifier        domain.TaskNotify
}

func NewTaskUseCase(taskRepository domain.TaskRepository, auditRepository domain.TaskAuditRepository, notifier domain.TaskNotify) TaskUseCase {
	uc := &taskUseCase{taskRepository, auditRepository, notifier}
	go uc.notifier.Subscribe(context.Background())
	return uc
}

func (u *taskUseCase) CreateTask(ctx context.Context, task *domain.Task) error {
	err := u.taskRepository.Create(ctx, task)
	if err == nil {
		go u.notifier.Send(ctx, task)
		go u.AuditTask(ctx, task.ID, CreateOperation)
	}
	return err
}

func (u *taskUseCase) GetAllTasks(ctx context.Context) ([]*domain.Task, error) {
	go u.AuditTask(ctx, 0, ReadOperation)
	return u.taskRepository.GetAll(ctx)
}

func (u *taskUseCase) GetTaskByID(ctx context.Context, id int64) (*domain.Task, error) {
	go u.AuditTask(ctx, id, ReadOperation)
	return u.taskRepository.GetByID(ctx, id)
}

func (u *taskUseCase) UpdateTask(ctx context.Context, task *domain.Task) error {
	err := u.taskRepository.Update(ctx, task)
	if err == nil {
		go u.AuditTask(ctx, task.ID, UpdateOperation)
	}
	return err
}

func (u *taskUseCase) DeleteTask(ctx context.Context, id int64) error {
	err := u.taskRepository.Delete(ctx, id)
	if err == nil {
		go u.AuditTask(ctx, 0, DeleteOperation)
	}
	return err
}

func (u *taskUseCase) AuditTask(ctx context.Context, taskID int64, operation string) {
	var performedBy int64
	userID := ctx.Value("user_id")
	if userID != nil {
		performedBy = int64(userID.(float64))
	}

	u.auditRepository.Create(ctx, &domain.TaskAudit{
		TaskID:      taskID,
		Operation:   operation,
		PerformedBy: &performedBy,
		PerformedAt: time.Now(),
	})
}
