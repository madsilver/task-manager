package notify

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/madsilver/task-manager/internal/adapter/handler"
	"github.com/madsilver/task-manager/internal/domain"
)

type Broker interface {
	Publish(ctx context.Context, data []byte) error
	Consume(ctx context.Context) (<-chan []byte, error)
}

type taskNotify struct {
	broker Broker
	log    handler.Logger
}

func NewTaskNotify(broker Broker, log handler.Logger) domain.TaskNotify {
	return &taskNotify{broker, log}
}

// Send sends a task notification to the message broker.
// ctx: the context for the operation
// task: the task to send
func (t taskNotify) Send(ctx context.Context, task *domain.Task) {
	task.Summary = "[protected]"
	data, err := json.Marshal(task)
	if err != nil {
		t.log.Error(ctx, "failed to marshal task", err)
		return
	}
	if err = t.broker.Publish(ctx, data); err != nil {
		t.log.Error(ctx, "failed to publish task to broker", err)
	}
}

// Subscribe listens for task notifications from the message broker.
// ctx: the context for the operation
func (t taskNotify) Subscribe(ctx context.Context) {
	msgCh, err := t.broker.Consume(ctx)
	if err != nil {
		t.log.Error(ctx, "failed to start broker consumption", err)
		return
	}

	go func() {
		for {
			select {
			case <-ctx.Done():
				t.log.Info(context.Background(), "subscription context cancelled")
				return
			case msg, ok := <-msgCh:
				if !ok {
					t.log.Info(context.Background(), "message channel closed")
					return
				}

				var task domain.Task
				if err = json.Unmarshal(msg, &task); err != nil {
					t.log.Error(context.Background(), "failed to unmarshal task message", err)
					continue
				}

				t.log.Info(context.Background(), fmt.Sprintf("notification received, task ID: %d", task.ID))

				// TODO: send this through a notification channel but out-of-door from tech challenge.
			}
		}
	}()
}
