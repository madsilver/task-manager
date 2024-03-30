package main

import (
	"github.com/madsilver/task-manager/internal/adapter/controller"
	"github.com/madsilver/task-manager/internal/adapter/repository/mysql"
	"github.com/madsilver/task-manager/internal/infra/broker"
	mysqlDB "github.com/madsilver/task-manager/internal/infra/db/mysql"
	"github.com/madsilver/task-manager/internal/infra/server"
	"log"
)

func main() {
	log.Println("task manager running")

	manager := &server.Manager{
		TaskController: controller.NewTaskController(
			mysql.NewTaskRepository(mysqlDB.NewMysqlDB()),
			broker.NewRabbitMQ()),
	}

	server.NewServer(manager).Start()
}
