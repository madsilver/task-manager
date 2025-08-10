package domain

import (
	"context"
	"time"
)

type Task struct {
	ID        int64
	Title     string
	Summary   string
	OwnerID   int64     `json:"owner_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type TaskRepository interface {
	Create(ctx context.Context, task *Task) error
	GetAll(ctx context.Context) ([]*Task, error)
	GetByID(ctx context.Context, id int64) (*Task, error)
	Update(ctx context.Context, task *Task) error
	Delete(ctx context.Context, id int64) error
}

type TaskNotify interface {
	Send(ctx context.Context, task *Task)
	Subscribe(ctx context.Context)
}
