package adapter

import (
	"github.com/madsilver/task-manager/internal/adapter/handler"
	"github.com/madsilver/task-manager/internal/adapter/notify"
	"github.com/madsilver/task-manager/internal/adapter/repository"
	"github.com/madsilver/task-manager/internal/usecase"
	"gorm.io/gorm"
)

type Manager struct {
	handler.TaskHandler
	handler.UserHandler
}

// NewManager initializes the Manager with the necessary dependencies.
// conn: the database connection
// log: the logger for logging operations
// Returns a pointer to the Manager instance.
func NewManager(conn *gorm.DB, broker notify.Broker, log handler.Logger) *Manager {
	taskUseCase := usecase.NewTaskUseCase(
		repository.NewTaskRepository(conn),
		notify.NewTaskNotify(broker, log),
	)

	userUseCase := usecase.NewUserUsecase(
		repository.NewUserRepository(conn),
	)

	return &Manager{
		TaskHandler: handler.NewTaskHandler(taskUseCase, log),
		UserHandler: handler.NewUserHandler(userUseCase, log),
	}
}
