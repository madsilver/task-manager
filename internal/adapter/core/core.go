package core

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/madsilver/task-manager/internal/adapter/presenter"
	"github.com/madsilver/task-manager/internal/entity"
)

type TaskService interface {
	FindTasks(ctx echo.Context) ([]*entity.Task, error)
	FindTaskByID(ctx echo.Context) (*entity.Task, error)
	CreateTask(ctx echo.Context, body *presenter.TaskCreate) (*entity.Task, error)
	UpdateTask(task *entity.Task) error
	DeleteTask(ctx echo.Context) error
	CloseTask(task *entity.Task) error
}

type NotifyService interface {
	Listen()
}

type Broker interface {
	Publish(ctx context.Context, data []byte) error
	Consume()
}

type TaskRepository interface {
	FindAll(args any) ([]*entity.Task, error)
	FindByID(args any) (*entity.Task, error)
	Create(task *entity.Task) error
	Update(task *entity.Task) error
	Delete(id any) error
}

type DB interface {
	Query(query string, args any, fn func(scan func(dest ...any) error) error) error
	QueryRow(query string, args any, fn func(scan func(dest ...any) error) error) error
	Save(query string, args ...any) (any, error)
	Update(query string, args ...any) error
	Delete(query string, args any) error
}
