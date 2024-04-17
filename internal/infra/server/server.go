package server

import (
	"context"
	"errors"
	"github.com/labstack/echo/v4"
	_ "github.com/madsilver/task-manager/docs"
	"github.com/madsilver/task-manager/internal/infra/env"
	"github.com/madsilver/task-manager/internal/infra/server/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type Server struct {
	manager *Manager
}

func NewServer(manager *Manager) *Server {
	return &Server{
		manager,
	}
}

// @title           Task Manager
// @version         1.0
// @description     API Developer Practical Exercise

// @contact.name   Rodrigo Prata
// @contact.email  rbpsilver@gmail.com

// @host      localhost:8000
// @BasePath  /v1

// Start Routes @externalDocs.description  OpenAPI
func (s *Server) Start() {
	e := echo.New()
	e.Use(ValidateHeader)
	e.Validator = middleware.ConfigValidator()

	e.GET("/v1/tasks", s.manager.TaskController.FindTasks, AuthRole(TechRole, AdminRole))
	e.GET("/v1/tasks/:id", s.manager.TaskController.FindTaskByID, AuthRole(TechRole, AdminRole))
	e.POST("/v1/tasks", s.manager.TaskController.CreateTask, AuthRole(TechRole))
	e.PATCH("/v1/tasks/:id", s.manager.TaskController.UpdateTask, AuthRole(TechRole))
	e.DELETE("/v1/tasks/:id", s.manager.TaskController.DeleteTask, AuthRole(AdminRole))
	e.PATCH("/v1/tasks/:id/close", s.manager.TaskController.CloseTask, AuthRole(TechRole))
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	go func() {
		host := ":" + env.GetString("SERVER_PORT", "8000")
		if err := e.Start(host); err != nil && !errors.Is(err, http.ErrServerClosed) {
			e.Logger.Fatal(err)
		}
	}()

	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
