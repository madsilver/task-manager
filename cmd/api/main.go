package main

import (
	"github.com/labstack/gommon/log"
	"github.com/madsilver/task-manager/internal/adapter/controller"
	"github.com/madsilver/task-manager/internal/adapter/repository/mysql"
	"github.com/madsilver/task-manager/internal/infra/broker"
	"github.com/madsilver/task-manager/internal/infra/db"
	"github.com/madsilver/task-manager/internal/infra/server"
)

func main() {
	log.Info("api running")

	manager := &server.Manager{
		TaskController: controller.NewTaskController(
			mysql.NewTaskRepository(db.NewMysqlDB()),
			broker.NewRabbitMQ()),
	}

	server.NewServer(manager).Start()
}
