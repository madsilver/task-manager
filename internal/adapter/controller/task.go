package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/madsilver/task-manager/internal/adapter/presenter"
	"github.com/madsilver/task-manager/internal/entity"
	"net/http"
)

type Repository interface {
	FindAll(args any) ([]*entity.Task, error)
	FindByID(args any) (*entity.Task, error)
	Create(task *entity.Task) error
	Update(task *entity.Task) error
	Delete(id any) error
}

type TaskController struct {
	repository Repository
}

func NewTaskController(repository Repository) *TaskController {
	return &TaskController{
		repository,
	}
}

func (c *TaskController) FindTasks(ctx echo.Context) error {
	var arg any
	if ctx.Get("role") == "technician" {
		arg = ctx.Get("user").(uint64)
	}
	tasks, err := c.repository.FindAll(arg)
	if err != nil {
		log.Error(err.Error())
		return ctx.JSON(http.StatusInternalServerError, presenter.InternalErrorResponse())
	}
	return ctx.JSON(http.StatusOK, presenter.PopulateTasks(tasks))
}

func (c *TaskController) FindTaskByID(ctx echo.Context) error {
	param := ctx.Param("id")
	task, err := c.repository.FindByID(param)
	if err != nil {
		log.Error(err.Error())
		return ctx.JSON(http.StatusInternalServerError, presenter.InternalErrorResponse())
	}
	return ctx.JSON(http.StatusOK, presenter.PopulateTask(task))
}

func (c *TaskController) CreateTask(ctx echo.Context) error {
	body := &presenter.TaskCreate{}
	if err := ctx.Bind(body); err != nil {
		log.Info(err.Error())
		return ctx.JSON(http.StatusBadRequest, presenter.NewErrorResponse("bad request", err.Error()))
	}
	task := &entity.Task{
		UserID:  ctx.Get("user").(uint64),
		Summary: body.Summary,
	}
	err := c.repository.Create(task)
	if err != nil {
		log.Error(err.Error())
		return ctx.JSON(http.StatusInternalServerError, presenter.InternalErrorResponse())
	}
	return ctx.JSON(http.StatusCreated, presenter.PopulateTask(task))
}

func (c *TaskController) UpdateTask(ctx echo.Context) error {
	body := &presenter.TaskCreate{}
	if err := ctx.Bind(body); err != nil {
		log.Info(err.Error())
		return ctx.JSON(http.StatusBadRequest, presenter.NewErrorResponse("bad request", err.Error()))
	}
	param := ctx.Param("id")
	task, err := c.repository.FindByID(param)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, presenter.NewErrorResponse("task not found", ""))
	}

	if task.UserID != ctx.Get("user").(uint64) {
		return ctx.JSON(http.StatusForbidden, presenter.NewErrorResponse("operation not allowed", ""))
	}

	task.Summary = body.Summary
	err = c.repository.Update(task)
	if err != nil {
		log.Error(err.Error())
		return ctx.JSON(http.StatusInternalServerError, presenter.InternalErrorResponse())
	}
	return ctx.JSON(http.StatusOK, presenter.PopulateTask(task))
}

func (c *TaskController) DeleteTask(ctx echo.Context) error {
	param := ctx.Param("id")
	_, err := c.repository.FindByID(param)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, presenter.NewErrorResponse("task not found", ""))
	}

	err = c.repository.Delete(param)
	if err != nil {
		log.Error(err.Error())
		return ctx.JSON(http.StatusInternalServerError, presenter.InternalErrorResponse())
	}
	return ctx.JSON(http.StatusOK, nil)
}
