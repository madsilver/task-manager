package main

import (
	"github.com/madsilver/task-manager/internal/adapter/controller"
	"github.com/madsilver/task-manager/internal/adapter/repository/mysql"
	mysqlDB "github.com/madsilver/task-manager/internal/infra/db/mysql"
	"github.com/madsilver/task-manager/internal/infra/server"
	"log"
)

func main() {
	log.Println("task manager running")

	db := mysqlDB.NewMysqlDB()

	repository := mysql.NewTaskRepository(db)

	manager := &server.Manager{
		TaskController: controller.NewTaskController(repository),
	}

	server.NewServer(manager).Start()
}
