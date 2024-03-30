package broker

import (
	"fmt"
	"github.com/labstack/gommon/log"
	"github.com/madsilver/task-manager/internal/infra/env"
	"github.com/rabbitmq/amqp091-go"
)

const QUEUE = "notify_manager"

type RabbitMQ struct {
	Channel *amqp091.Channel
}

func NewRabbitMQ() *RabbitMQ {
	conn, err := amqp091.Dial(getDSN())
	if err != nil {
		log.Fatal(err)
	}

	channel, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}

	_, err = channel.QueueDeclare(QUEUE, true, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Info("successfully connected to RabbitMQ")

	return &RabbitMQ{
		Channel: channel,
	}
}

func (r *RabbitMQ) Publish(data []byte) error {
	message := amqp091.Publishing{ContentType: "text/plain", Body: data}
	err := r.Channel.Publish("", QUEUE, false, false, message)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func (r *RabbitMQ) Consume() {
	err := r.Channel.Qos(4, 0, false)
	if err != nil {
		log.Fatal(err)
	}

	msgs, err := r.Channel.Consume(QUEUE, "", false, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	var forever chan struct{}

	go func() {
		for msg := range msgs {
			log.Info(string(msg.Body))
			_ = msg.Ack(true)
		}
	}()

	<-forever
}

func getDSN() string {
	return fmt.Sprintf("amqp://%s:%s@%s:%s",
		env.GetString("RABBITMQ_USER", "silver"),
		env.GetString("RABBITMQ_PASSWORD", "silver"),
		env.GetString("RABBITMQ_HOST", "127.0.0.1"),
		env.GetString("RABBITMQ_PORT", "5672"))
}
