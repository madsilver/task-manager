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
