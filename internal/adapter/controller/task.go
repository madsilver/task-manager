package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/madsilver/task-manager/internal/adapter/presenter"
	"github.com/madsilver/task-manager/internal/entity"
	"net/http"
)

type Repository interface {
	FindAllTasks(args any) ([]*entity.Task, error)
	CreateTask(task *entity.Task) error
	UpdateTask(task *entity.Task) error
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
	userParam := ctx.QueryParam("user")
	tasks, err := c.repository.FindAllTasks(userParam)
	if err != nil {
		log.Error(err.Error())
		return ctx.JSON(http.StatusInternalServerError, presenter.InternalErrorResponse())
	}
	return ctx.JSON(http.StatusOK, presenter.PopulateTasks(tasks))
}

func (c *TaskController) CreateTask(ctx echo.Context) error {
	body := &presenter.TaskCreate{}
	if err := ctx.Bind(body); err != nil {
		log.Info(err.Error())
		return ctx.JSON(http.StatusBadRequest, presenter.NewErrorResponse("bad request", err.Error()))
	}
	task, _ := entity.NewTask(ctx.Get("user"), body.Summary)
	err := c.repository.CreateTask(task)
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
	task, _ := entity.NewTask(ctx.Get("user"), body.Summary)
	err := c.repository.UpdateTask(task)
	if err != nil {
		log.Error(err.Error())
		return ctx.JSON(http.StatusInternalServerError, presenter.InternalErrorResponse())
	}
	return ctx.JSON(http.StatusOK, presenter.PopulateTask(task))
}
