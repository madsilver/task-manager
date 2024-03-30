package controller

type NotifyController struct {
	broker Broker
}

func NewNotifyController(broker Broker) *NotifyController {
	return &NotifyController{
		broker,
	}
}

func (c *NotifyController) Listen() {
	c.broker.Consume()
}
