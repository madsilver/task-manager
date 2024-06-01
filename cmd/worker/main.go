package main

import (
	"github.com/labstack/gommon/log"
	"github.com/madsilver/task-manager/internal/adapter/controller"
	"github.com/madsilver/task-manager/internal/adapter/service"
	"github.com/madsilver/task-manager/internal/infra/broker"
)

func main() {
	log.Info("worker running")

	svc := service.NewNotifyService(broker.NewRabbitMQ())

	controller.NewNotifyController(svc).Listen()
}
