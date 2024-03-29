package server

import (
	"github.com/labstack/echo/v4"
	"github.com/madsilver/task-manager/internal/infra/env"
)

type Server struct {
	manager *Manager
}

func NewServer(manager *Manager) *Server {
	return &Server{
		manager,
	}
}

func (s *Server) Start() {
	e := echo.New()
	e.Use(ValidateHeader)
	e.GET("/v1/tasks", s.manager.TaskController.FindTasks, AuthAdminOrOwner)

	e.Logger.Fatal(e.Start(":" + env.GetString("SERVER_PORT", "8000")))
}
