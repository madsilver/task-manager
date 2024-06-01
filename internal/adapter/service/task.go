package service

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/madsilver/task-manager/internal/adapter/core"
	"github.com/madsilver/task-manager/internal/adapter/presenter"
	"github.com/madsilver/task-manager/internal/entity"
	"time"
)

type taskService struct {
	repository core.TaskRepository
	broker     core.Broker
}

func NewTaskService(repository core.TaskRepository, broker core.Broker) core.TaskService {
	return &taskService{
		repository,
		broker,
	}
}

func (s *taskService) FindTasks(ctx echo.Context) ([]*entity.Task, error) {
	var arg any
	if ctx.Get("role") == "technician" {
		arg = ctx.Get("user").(uint64)
	}
	return s.repository.FindAll(arg)
}

func (s *taskService) FindTaskByID(ctx echo.Context) (*entity.Task, error) {
	param := ctx.Param("id")
	return s.repository.FindByID(param)
}

func (s *taskService) CreateTask(ctx echo.Context, body *presenter.TaskCreate) (*entity.Task, error) {
	task := &entity.Task{
		UserID:  ctx.Get("user").(uint64),
		Summary: body.Summary,
	}
	return task, s.repository.Create(task)
}

func (s *taskService) UpdateTask(task *entity.Task) error {
	return s.repository.Update(task)
}

func (s *taskService) DeleteTask(ctx echo.Context) error {
	param := ctx.Param("id")
	return s.repository.Delete(param)
}

func (s *taskService) CloseTask(task *entity.Task) error {
	now := time.Now().Format("2006-01-02 15:04:05")
	task.Date = &now
	err := s.repository.Update(task)
	if err != nil {
		return err
	}

	go func() {
		_ = s.Notify(task)
	}()

	return nil
}

func (s *taskService) Notify(task *entity.Task) (err error) {
	message := fmt.Sprintf("The tech %d performed the task on date %s", task.UserID, *task.Date)
	err = s.broker.Publish(context.Background(), []byte(message))
	if err != nil {
		log.Error(err.Error())
	}
	return
}
