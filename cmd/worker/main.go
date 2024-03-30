package main

import (
	"github.com/labstack/gommon/log"
	"github.com/madsilver/task-manager/internal/adapter/controller"
	"github.com/madsilver/task-manager/internal/infra/broker"
)

func main() {
	log.Info("worker running")

	controller.NewNotifyController(broker.NewRabbitMQ()).Listen()
}
