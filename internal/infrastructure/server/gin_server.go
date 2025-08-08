package server

import (
	"errors"
	"github.com/gin-gonic/gin"
	docs "github.com/madsilver/task-manager/docs/swagger"
	"github.com/madsilver/task-manager/internal/adapter"
	"github.com/madsilver/task-manager/internal/adapter/handler"
	"github.com/madsilver/task-manager/internal/domain"
	"github.com/madsilver/task-manager/internal/infrastructure/env"
	"github.com/madsilver/task-manager/internal/infrastructure/server/middleware"
	swaggerfiles "github.com/swaggo/files"
	swaggerGin "github.com/swaggo/gin-swagger"
	"io"
	"net/http"
	"os"
	"os/signal"
)

type GinServer struct {
	manager *adapter.Manager
}

func NewGinServer(manager *adapter.Manager) *GinServer {
	return &GinServer{
		manager: manager,
	}
}

// Start starts the Gin server.
// Returns a channel to listen for errors during server operation.
func (s *GinServer) Start() chan error {
	gin.DefaultWriter = io.Discard
	e := gin.Default()

	s.Routes(e)

	srv := &http.Server{
		Addr:              ":" + env.App.ServerPort,
		Handler:           e,
		ReadHeaderTimeout: env.App.ServerReadHeaderTimeout,
		WriteTimeout:      env.App.ServerWriteTimeout,
		IdleTimeout:       env.App.ServerIdleTimeout,
	}

	err := make(chan error, 1)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	go func() {
		<-quit
		err <- srv.Close()
	}()

	go func() {
		if errSrv := srv.ListenAndServe(); errSrv != nil && !errors.Is(errSrv, http.ErrServerClosed) {
			err <- errSrv
		}
	}()

	return err
}

// @title           Task Manager API
// @version         1.0
// @description     Task Manager API is a Restful API for managing tasks.

// @contact.name   Rodrigo Prata
// @contact.email  rbpsilver@gmail.com

// @host      localhost:8000
// @BasePath  /api/v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

// Routes @externalDocs.description  OpenAPI
// @externalDocs.url          https://www.linkedin.com/in/silverdev
func (s *GinServer) Routes(e *gin.Engine) {
	docs.SwaggerInfo.BasePath = "/api/v1"
	e.GET("/swagger/*any", swaggerGin.WrapHandler(swaggerfiles.Handler))
	e.POST("/api/v1/login", route(s.manager.UserHandler.Login))

	api := e.Group("/api/v1").Use(middleware.JWTAuth())

	// User routes
	api.POST("/users", middleware.HasRole(domain.RoleManager), route(s.manager.UserHandler.Create))

	// Task routes
	api.POST("/tasks", middleware.HasRole(domain.RoleTechnician), route(s.manager.TaskHandler.Create))
	api.GET("/tasks", middleware.HasRole(domain.RoleManager), route(s.manager.TaskHandler.GetAll))
	api.GET("/tasks/:id", middleware.HasRole(domain.RoleTechnician, domain.RoleManager), route(s.manager.TaskHandler.GetByID))
	api.PUT("/tasks/:id", middleware.HasRole(domain.RoleManager), route(s.manager.TaskHandler.Update))
	api.DELETE("/tasks/:id", middleware.HasRole(domain.RoleManager), route(s.manager.TaskHandler.Delete))
}

type HandlerHTTP func(ctx handler.ContextHTTP) error

// route is a middleware function that wraps the HandlerHTTP function.
// handler: The HandlerHTTP function to be executed when the route is hit.
// Returns a gin.HandlerFunc that can be used in Gin routes.
func route(handler HandlerHTTP) gin.HandlerFunc {
	return func(c *gin.Context) {
		_ = handler(NewContextHTTP(c))
	}
}
