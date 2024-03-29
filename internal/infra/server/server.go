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
	e.GET("/v1/tasks", s.manager.TaskController.FindTasks, AuthRole(TechRole, AdminRole))
	e.POST("/v1/tasks", s.manager.TaskController.CreateTask, AuthRole(TechRole))
	e.PATCH("/v1/tasks/:id", s.manager.TaskController.UpdateTask, AuthRole(TechRole))
	e.DELETE("/v1/tasks/:id", s.manager.TaskController.DeleteTask, AuthRole(AdminRole))

	e.Logger.Fatal(e.Start(":" + env.GetString("SERVER_PORT", "8000")))
}
