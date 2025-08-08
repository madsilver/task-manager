package domain

import (
	"context"
	"time"
)

type TaskAudit struct {
	ID          int64
	TaskID      int64
	Operation   string
	PerformedBy *int64
	PerformedAt time.Time
}

type TaskAuditRepository interface {
	Create(ctx context.Context, audit *TaskAudit) error
}
