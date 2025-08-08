package handler

import (
	"github.com/madsilver/task-manager/internal/adapter/presenter"
	"github.com/madsilver/task-manager/internal/usecase"
	"net/http"
	"strconv"

	"github.com/madsilver/task-manager/internal/domain"
)

type TaskHandler interface {
	Create(c ContextHTTP) error
	GetAll(c ContextHTTP) error
	GetByID(c ContextHTTP) error
	Update(c ContextHTTP) error
	Delete(c ContextHTTP) error
}

type taskHandler struct {
	uc  usecase.TaskUseCase
	log Logger
}

func NewTaskHandler(uc usecase.TaskUseCase, log Logger) TaskHandler {
	return &taskHandler{uc, log}
}

// Create godoc
// @Summary Create task.
// @Description Create new task.
// @Tags Task
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param Task body domain.Task true "Task payload"
// @Success 201 {object} presenter.TaskResponse
// @Failure 400 {object} presenter.ErrorResponse
// @Failure 500 {object} presenter.ErrorResponse
// @Router /tasks [post]
func (h *taskHandler) Create(c ContextHTTP) error {
	var task domain.Task
	if err := c.Bind(&task); err != nil {
		h.log.Error(c, "failed to bind task", err)
		return c.JSON(http.StatusBadRequest, presenter.ErrorResponse{Error: "invalid request"})
	}

	task.OwnerID = h.UserIDFromContext(c)

	if err := h.uc.CreateTask(c, &task); err != nil {
		h.log.Error(c, "failed to create task", err)
		return c.JSON(http.StatusInternalServerError, presenter.ErrorResponse{Error: "failed to create task"})
	}

	return c.JSON(http.StatusCreated, presenter.ToTaskResponse(&task))
}

// GetAll godoc
// @Summary Find all tasks.
// @Description Find all tasks.
// @Tags Task
// @Produce json
// @Security BearerAuth
// @Success 200 {object} presenter.TaskResponse
// @Failure 404 {object} presenter.ErrorResponse
// @Router /tasks [get]
func (h *taskHandler) GetAll(c ContextHTTP) error {
	tasks, err := h.uc.GetAllTasks(c)
	if err != nil {
		return c.JSON(http.StatusNotFound, presenter.ErrorResponse{Error: "tasks not found"})
	}
	return c.JSON(http.StatusOK, presenter.ToTaskResponseList(tasks))
}

// GetByID godoc
// @Summary Find task by ID.
// @Description Find a task by the ID.
// @Tags Task
// @Produce json
// @Security BearerAuth
// @Param id path int true "Task ID"
// @Success 200 {object} presenter.TaskResponse
// @Failure 404 {object} presenter.ErrorResponse
// @Failure 500 {object} presenter.ErrorResponse
// @Router /tasks/{id} [get]
func (h *taskHandler) GetByID(c ContextHTTP) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, presenter.ErrorResponse{Error: "invalid id"})
	}

	task, err := h.uc.GetTaskByID(c, id)
	if err != nil {
		return c.JSON(http.StatusNotFound, presenter.ErrorResponse{Error: "task not found"})
	}

	if !h.IsAllowedToRead(c, task) {
		return c.JSON(http.StatusForbidden, presenter.ErrorResponse{Error: "not allowed to read this task"})
	}

	return c.JSON(http.StatusOK, presenter.ToTaskResponse(task))
}

// Update godoc
// @Summary Update task.
// @Description Update a task by ID.
// @Tags Task
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Task ID"
// @Param Task body domain.Task true "Task payload"
// @Success 200 {object} presenter.TaskResponse
// @Failure 400 {object} presenter.ErrorResponse
// @Failure 500 {object} presenter.ErrorResponse
// @Router /tasks/{id} [put]
func (h *taskHandler) Update(c ContextHTTP) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, presenter.ErrorResponse{Error: "invalid id"})
	}

	var update domain.Task
	if err := c.Bind(&update); err != nil {
		return c.JSON(http.StatusBadRequest, presenter.ErrorResponse{Error: "invalid payload"})
	}
	update.ID = id

	if err := h.uc.UpdateTask(c, &update); err != nil {
		return c.JSON(http.StatusInternalServerError, presenter.ErrorResponse{Error: err.Error()})
	}
	return c.JSON(http.StatusOK, update)
}

// Delete godoc
// @Summary Delete task.
// @Description Delete a task by ID.
// @Tags Task
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Task ID"
// @Success 204
// @Failure 400 {object} presenter.ErrorResponse
// @Failure 500 {object} presenter.ErrorResponse
// @Router /tasks/{id} [delete]
func (h *taskHandler) Delete(c ContextHTTP) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, presenter.ErrorResponse{Error: "invalid id"})
	}

	if err := h.uc.DeleteTask(c, id); err != nil {
		return c.JSON(http.StatusInternalServerError, presenter.ErrorResponse{Error: err.Error()})
	}
	return c.JSON(http.StatusNoContent, nil)
}

// UserIDFromContext extracts the user ID from the context.
func (h *taskHandler) UserIDFromContext(c ContextHTTP) int64 {
	userID := c.Value("user_id")
	if userID == nil {
		return 0
	}
	return int64(userID.(float64))
}

// IsAllowedToRead checks if the user is allowed to read the task
// c: ContextHTTP
// task: the task to check
// returns true if the user is allowed to read the task, false otherwise
func (h *taskHandler) IsAllowedToRead(c ContextHTTP, task *domain.Task) bool {
	roles := c.Value("roles")
	if roles != nil {
		for _, role := range roles.([]interface{}) {
			if role == "manager" {
				return true
			}
		}
	}
	return task.OwnerID == h.UserIDFromContext(c)
}
