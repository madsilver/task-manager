package controller

import (
	"context"
	"github.com/madsilver/task-manager/internal/entity"
)

type Repository interface {
	FindAll(args any) ([]*entity.Task, error)
	FindByID(args any) (*entity.Task, error)
	Create(task *entity.Task) error
	Update(task *entity.Task) error
	Delete(id any) error
}

type Broker interface {
	Publish(ctx context.Context, data []byte) error
	Consume()
}
