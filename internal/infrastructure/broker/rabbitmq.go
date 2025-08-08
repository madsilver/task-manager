package broker

import (
	"context"
	"fmt"
	"github.com/madsilver/task-manager/internal/adapter/notify"
	"github.com/madsilver/task-manager/internal/infrastructure/env"
	"github.com/rabbitmq/amqp091-go"
	"log"
)

const (
	ContentType   = "text/plain"
	PrefetchCount = 4
)

type rabbitMQ struct {
	channel   *amqp091.Channel
	queueName string
}

func NewRabbitMQ(queueName string) notify.Broker {
	conn, err := amqp091.Dial(getDSN())
	if err != nil {
		log.Fatalf("rabbitmq connection error: %v", err)
	}

	channel, err := conn.Channel()
	if err != nil {
		log.Fatalf("rabbitmq channel error: %v", err)
	}

	_, err = channel.QueueDeclare(queueName, true, false, false, false, nil)
	if err != nil {
		log.Fatalf("rabbitmq queue declare error: %v", err)
	}

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
				log.Printf("failed to acknowledge message: %v", err)
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
