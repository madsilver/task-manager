package server

import "github.com/madsilver/task-manager/internal/adapter/controller"

type Manager struct {
	TaskController *controller.TaskController
}
