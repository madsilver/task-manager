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
	UserIDFromContext(c context.Context) int64
	IsAllowedToRead(c context.Context, task *domain.Task) bool
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

// CreateTask creates a new task and sends a notification asynchronously.
// ctx: the context for the operation
// task: the task to be created
// returns an error if the task could not be created
func (u *taskUseCase) CreateTask(ctx context.Context, task *domain.Task) error {
	task.OwnerID = u.UserIDFromContext(ctx)
	err := u.taskRepository.Create(ctx, task)
	if err == nil {
		go u.notifier.Send(ctx, task)
		go u.AuditTask(ctx, task.ID, CreateOperation)
	}
	return err
}

// GetAllTasks retrieves all tasks and logs the read operation asynchronously.
// ctx: the context for the operation
// returns a slice of tasks and an error if any
func (u *taskUseCase) GetAllTasks(ctx context.Context) ([]*domain.Task, error) {
	go u.AuditTask(ctx, 0, ReadOperation)
	return u.taskRepository.GetAll(ctx)
}

// GetTaskByID retrieves a task by its ID and logs the read operation asynchronously.
// ctx: the context for the operation
// id: the ID of the task to retrieve
// returns the task and an error if any
func (u *taskUseCase) GetTaskByID(ctx context.Context, id int64) (*domain.Task, error) {
	go u.AuditTask(ctx, id, ReadOperation)
	return u.taskRepository.GetByID(ctx, id)
}

// UpdateTask updates an existing task and logs the update operation asynchronously.
// ctx: the context for the operation
// task: the task to be updated
// returns an error if the task could not be updated
func (u *taskUseCase) UpdateTask(ctx context.Context, task *domain.Task) error {
	err := u.taskRepository.Update(ctx, task)
	if err == nil {
		go u.AuditTask(ctx, task.ID, UpdateOperation)
	}
	return err
}

// DeleteTask deletes a task by its ID and logs the delete operation asynchronously.
// ctx: the context for the operation
// id: the ID of the task to delete
// returns an error if the task could not be deleted
func (u *taskUseCase) DeleteTask(ctx context.Context, id int64) error {
	err := u.taskRepository.Delete(ctx, id)
	if err == nil {
		go u.AuditTask(ctx, 0, DeleteOperation)
	}
	return err
}

// AuditTask logs the operation performed on a task.
// ctx: the context for the operation
// taskID: the ID of the task on which the operation was performed
// operation: the type of operation performed (CREATE, UPDATE, DELETE, READ)
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

// UserIDFromContext extracts the user ID from the context.
// c: the context containing user ID
// returns the user ID as int64, or 0 if not found
func (u *taskUseCase) UserIDFromContext(c context.Context) int64 {
	userID := c.Value("user_id")
	if userID == nil {
		return 0
	}
	return int64(userID.(float64))
}

// IsAllowedToRead checks if the user is allowed to read the task
// c: the context containing user roles and ID
// task: the task to check
// returns true if the user is allowed to read the task, false otherwise
func (u *taskUseCase) IsAllowedToRead(c context.Context, task *domain.Task) bool {
	roles := c.Value("roles")
	if roles != nil {
		for _, role := range roles.([]interface{}) {
			if role == "manager" {
				return true
			}
		}
	}
	return task.OwnerID == u.UserIDFromContext(c)
}
