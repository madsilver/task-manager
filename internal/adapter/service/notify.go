package service

import "github.com/madsilver/task-manager/internal/adapter/core"

type notifyService struct {
	broker core.Broker
}

func NewNotifyService(broker core.Broker) core.NotifyService {
	return &notifyService{
		broker,
	}
}

func (c *notifyService) Listen() {
	c.broker.Consume()
}
