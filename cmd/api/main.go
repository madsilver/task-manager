package main

import (
	"context"
	"github.com/madsilver/task-manager/internal/adapter"
	"github.com/madsilver/task-manager/internal/infrastructure/broker"
	"github.com/madsilver/task-manager/internal/infrastructure/db"
	"github.com/madsilver/task-manager/internal/infrastructure/logger"
	"github.com/madsilver/task-manager/internal/infrastructure/server"
)

const NotificationQueue = "task-manager.task.created"

func main() {
	ctx := context.Background()
	log := logger.NewLogger()
	log.Info(ctx, "starting api...")

	srv := server.NewGinServer(
		adapter.NewManager(
			db.NewMySQLDB(),
			broker.NewRabbitMQ(ctx, NotificationQueue, log),
			log,
		),
	)

	err := <-srv.Start()
	if err != nil {
		log.Error(ctx, "api down", err)
	} else {
		log.Info(ctx, "api shutdown")
	}
}
