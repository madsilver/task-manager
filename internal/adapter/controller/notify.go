package controller

import "github.com/madsilver/task-manager/internal/adapter/core"

type NotifyController struct {
	service core.NotifyService
}

func NewNotifyController(service core.NotifyService) *NotifyController {
	return &NotifyController{
		service,
	}
}

func (c *NotifyController) Listen() {
	c.service.Listen()
}
