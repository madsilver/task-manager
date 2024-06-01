package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/madsilver/task-manager/internal/adapter/core"
	"github.com/madsilver/task-manager/internal/adapter/presenter"
	"net/http"
)

const (
	TaskNotFoundMessage = "task not found"
	NotAllowedMessage   = "operation not allowed"
	BadRequestMessage   = "bad request"
)

type TaskController struct {
	service core.TaskService
}

func NewTaskController(service core.TaskService) *TaskController {
	return &TaskController{
		service,
	}
}

// FindTasks godoc
// @Summary Find tasks.
// @Description Find all tasks owned by a user. If the user is a manager, it returns all tasks of all users.
// @Tags Tasks
// @Accept  json
// @Produce  json
// @Param x-user-id header int64 true "User ID"
// @Param x-role header string true "Role"
// @Success 200 {object} []presenter.Task
// @Failure 403 {object} presenter.ErrorResponse
// @Failure 500 {object} presenter.ErrorResponse
// @Router /tasks [get]
func (c *TaskController) FindTasks(ctx echo.Context) error {
	tasks, err := c.service.FindTasks(ctx)
	if err != nil {
		log.Error(err.Error())
		return ctx.JSON(http.StatusInternalServerError, presenter.InternalErrorResponse())
	}
	return ctx.JSON(http.StatusOK, presenter.PopulateTasks(tasks))
}

// FindTaskByID godoc
// @Summary Find task.
// @Description Find task by ID.
// @Tags Tasks
// @Accept  json
// @Produce  json
// @Param x-user-id header int64 true "User ID"
// @Param x-role header string true "Role"
// @Param id path int64 true "Task ID"
// @Success 200 {object} presenter.Task
// @Failure 403 {object} presenter.ErrorResponse
// @Failure 404 {object} presenter.ErrorResponse
// @Router /tasks/{id} [get]
func (c *TaskController) FindTaskByID(ctx echo.Context) error {
	task, err := c.service.FindTaskByID(ctx)
	if err != nil {
		log.Error(err.Error())
		return ctx.JSON(http.StatusNotFound, presenter.NewErrorResponse(TaskNotFoundMessage, ""))
	}
	return ctx.JSON(http.StatusOK, presenter.PopulateTask(task))
}

// CreateTask godoc
// @Summary Create task.
// @Description Create a new task.
// @Tags Tasks
// @Accept  json
// @Produce  json
// @Param x-user-id header int64 true "User ID"
// @Param x-role header string true "Role"
// @Param Task body presenter.TaskCreate true " "
// @Success 201 {object} presenter.Task
// @Failure 403 {object} presenter.ErrorResponse
// @Failure 404 {object} presenter.ErrorResponse
// @Failure 500 {object} presenter.ErrorResponse
// @Router /tasks [post]
func (c *TaskController) CreateTask(ctx echo.Context) error {
	body := &presenter.TaskCreate{}
	_ = ctx.Bind(body)
	if err := ctx.Validate(body); err != nil {
		return ctx.JSON(http.StatusBadRequest, presenter.NewErrorResponse(BadRequestMessage, err.Error()))
	}
	task, err := c.service.CreateTask(ctx, body)
	if err != nil {
		log.Error(err.Error())
		return ctx.JSON(http.StatusInternalServerError, presenter.InternalErrorResponse())
	}
	return ctx.JSON(http.StatusCreated, presenter.PopulateTask(task))
}

// UpdateTask godoc
// @Summary Update task.
// @Description Update task by ID.
// @Tags Tasks
// @Accept  json
// @Produce  json
// @Param x-user-id header int64 true "User ID"
// @Param x-role header string true "Role"
// @Param Task body presenter.TaskCreate true " "
// @Param id path int64 true "Task ID"
// @Success 200 {object} presenter.Task
// @Failure 403 {object} presenter.ErrorResponse
// @Failure 404 {object} presenter.ErrorResponse
// @Failure 500 {object} presenter.ErrorResponse
// @Router /tasks/{id} [patch]
func (c *TaskController) UpdateTask(ctx echo.Context) error {
	body := &presenter.TaskCreate{}
	_ = ctx.Bind(body)
	if err := ctx.Validate(body); err != nil {
		return ctx.JSON(http.StatusBadRequest, presenter.NewErrorResponse(BadRequestMessage, err.Error()))
	}
	task, err := c.service.FindTaskByID(ctx)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, presenter.NewErrorResponse(TaskNotFoundMessage, ""))
	}

	if task.UserID != ctx.Get("user").(uint64) {
		return ctx.JSON(http.StatusForbidden, presenter.NewErrorResponse(NotAllowedMessage, ""))
	}

	task.Summary = body.Summary
	err = c.service.UpdateTask(task)
	if err != nil {
		log.Error(err.Error())
		return ctx.JSON(http.StatusInternalServerError, presenter.InternalErrorResponse())
	}
	return ctx.JSON(http.StatusOK, presenter.PopulateTask(task))
}

// DeleteTask godoc
// @Summary Delete task.
// @Description Delete task by ID.
// @Tags Tasks
// @Accept  json
// @Produce  json
// @Param x-user-id header int64 true "User ID"
// @Param x-role header string true "Role"
// @Param id path int64 true "Task ID"
// @Success 200
// @Failure 403 {object} presenter.ErrorResponse
// @Failure 404 {object} presenter.ErrorResponse
// @Failure 500 {object} presenter.ErrorResponse
// @Router /tasks/{id} [delete]
func (c *TaskController) DeleteTask(ctx echo.Context) error {
	_, err := c.service.FindTaskByID(ctx)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, presenter.NewErrorResponse(TaskNotFoundMessage, ""))
	}

	err = c.service.DeleteTask(ctx)
	if err != nil {
		log.Error(err.Error())
		return ctx.JSON(http.StatusInternalServerError, presenter.InternalErrorResponse())
	}
	return ctx.JSON(http.StatusOK, nil)
}

// CloseTask godoc
// @Summary Close task.
// @Description Close task by ID.
// @Tags Tasks
// @Accept  json
// @Produce  json
// @Param x-user-id header int64 true "User ID"
// @Param x-role header string true "Role"
// @Param id path int64 true "Task ID"
// @Success 200 {object} presenter.Task
// @Failure 403 {object} presenter.ErrorResponse
// @Failure 404 {object} presenter.ErrorResponse
// @Failure 500 {object} presenter.ErrorResponse
// @Router /tasks/{id}/close [patch]
func (c *TaskController) CloseTask(ctx echo.Context) error {
	task, err := c.service.FindTaskByID(ctx)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, presenter.NewErrorResponse(TaskNotFoundMessage, ""))
	}

	if task.UserID != ctx.Get("user").(uint64) {
		return ctx.JSON(http.StatusForbidden, presenter.NewErrorResponse(NotAllowedMessage, ""))
	}

	err = c.service.CloseTask(task)
	if err != nil {
		log.Error(err.Error())
		return ctx.JSON(http.StatusInternalServerError, presenter.InternalErrorResponse())
	}

	return ctx.JSON(http.StatusOK, presenter.PopulateTask(task))
}
