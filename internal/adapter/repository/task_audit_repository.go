package repository

import (
	"context"
	"time"
)

const (
	CreateOperation = "CREATE"
	UpdateOperation = "UPDATE"
	DeleteOperation = "DELETE"
	ReadOperation   = "READ"
)

type taskAuditLogs struct {
	ID          int64 `gorm:"primaryKey"`
	TaskID      int64
	Operation   string
	PerformedBy *int64
	PerformedAt time.Time
}

func (r *taskRepository) logTaskAudit(ctx context.Context, taskID int64, op string) {
	var performedBy int64

	userID := ctx.Value("user_id")
	if userID != nil {
		performedBy = int64(userID.(float64))
	}

	log := &taskAuditLogs{
		TaskID:      taskID,
		Operation:   op,
		PerformedBy: &performedBy,
		PerformedAt: time.Now(),
	}

	r.db.Create(log)
}
