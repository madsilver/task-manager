package controller

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/madsilver/task-manager/internal/adapter/presenter"
	"github.com/madsilver/task-manager/internal/entity"
	"net/http"
	"time"
)

const (
	TaskNotFoundMessage = "task not found"
	NotAllowedMessage   = "operation not allowed"
	BadRequestMessage   = "bad request"
)

type TaskController struct {
	repository Repository
	broker     Broker
}

func NewTaskController(repository Repository, broker Broker) *TaskController {
	return &TaskController{
		repository,
		broker,
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
	param := ctx.Param("id")
	task, err := c.repository.FindByID(param)
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
	param := ctx.Param("id")
	task, err := c.repository.FindByID(param)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, presenter.NewErrorResponse(TaskNotFoundMessage, ""))
	}

	if task.UserID != ctx.Get("user").(uint64) {
		return ctx.JSON(http.StatusForbidden, presenter.NewErrorResponse(NotAllowedMessage, ""))
	}

	task.Summary = body.Summary
	err = c.repository.Update(task)
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
	param := ctx.Param("id")
	_, err := c.repository.FindByID(param)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, presenter.NewErrorResponse(TaskNotFoundMessage, ""))
	}

	err = c.repository.Delete(param)
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
	param := ctx.Param("id")
	task, err := c.repository.FindByID(param)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, presenter.NewErrorResponse(TaskNotFoundMessage, ""))
	}

	if task.UserID != ctx.Get("user").(uint64) {
		return ctx.JSON(http.StatusForbidden, presenter.NewErrorResponse(NotAllowedMessage, ""))
	}

	now := time.Now().Format("2006-01-02 15:04:05")
	task.Date = &now
	err = c.repository.Update(task)
	if err != nil {
		log.Error(err.Error())
		return ctx.JSON(http.StatusInternalServerError, presenter.InternalErrorResponse())
	}

	go func() {
		_ = c.Notify(task)
	}()

	return ctx.JSON(http.StatusOK, presenter.PopulateTask(task))
}

func (c *TaskController) Notify(task *entity.Task) (err error) {
	message := fmt.Sprintf("The tech %d performed the task %s on date %s", task.UserID, task.Summary, *task.Date)
	err = c.broker.Publish([]byte(message))
	if err != nil {
		log.Error(err.Error())
	}
	return
}
