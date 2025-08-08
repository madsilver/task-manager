package broker

import (
	"context"
	"fmt"
	"github.com/madsilver/task-manager/internal/adapter/handler"
	"github.com/madsilver/task-manager/internal/adapter/notify"
	"github.com/madsilver/task-manager/internal/infrastructure/env"
	"github.com/rabbitmq/amqp091-go"
)

const (
	ContentType   = "text/plain"
	PrefetchCount = 4
)

type rabbitMQ struct {
	channel   *amqp091.Channel
	queueName string
	log       handler.Logger
}

func NewRabbitMQ(ctx context.Context, queueName string, log handler.Logger) notify.Broker {
	conn, err := amqp091.Dial(getDSN())
	if err != nil {
		log.Panic(ctx, "rabbitmq connection error:"+err.Error())
	}

	channel, err := conn.Channel()
	if err != nil {
		log.Panic(ctx, "rabbitmq channel error:"+err.Error())
	}

	_, err = channel.QueueDeclare(queueName, true, false, false, false, nil)
	if err != nil {
		log.Panic(ctx, "rabbitmq queue declare  error:"+err.Error())
	}

	log.Info(ctx, "rabbitmq connection established")

	return &rabbitMQ{
		channel:   channel,
		queueName: queueName,
	}
}

// Publish sends a message to the RabbitMQ queue.
// ctx: the context for the operation
// data: the message data to be sent
// returns error if the operation fails
func (r *rabbitMQ) Publish(ctx context.Context, data []byte) error {
	message := amqp091.Publishing{ContentType: ContentType, Body: data}
	return r.channel.PublishWithContext(ctx, "", r.queueName, false, false, message)
}

// Consume starts consuming messages from the RabbitMQ queue.
// ctx: the context for the operation
func (r *rabbitMQ) Consume(ctx context.Context) (<-chan []byte, error) {
	if err := r.channel.Qos(PrefetchCount, 0, false); err != nil {
		return nil, fmt.Errorf("qos: %w", err)
	}

	deliveryCh, err := r.channel.Consume(
		r.queueName,
		"tasks-consumer",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("consume: %w", err)
	}

	out := make(chan []byte, 1)
	go func() {
		defer close(out)
		for d := range deliveryCh {
			out <- d.Body
			if err := d.Ack(false); err != nil {
				r.log.Error(ctx, "failed to acknowledge message", err)
			}
		}
	}()
	return out, nil
}

func getDSN() string {
	return fmt.Sprintf("amqp://%s:%s@%s:%s",
		env.App.RabbitMQUser,
		env.App.RabbitMQPassword,
		env.App.RabbitMQHost,
		env.App.RabbitMQPort)
}
