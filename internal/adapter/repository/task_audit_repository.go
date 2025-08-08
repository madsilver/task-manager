package repository

import (
	"context"
	"github.com/madsilver/task-manager/internal/domain"
	"gorm.io/gorm"
	"time"
)

type taskAuditRepository struct {
	db *gorm.DB
}

type taskAudit struct {
	ID          int64 `gorm:"primaryKey"`
	TaskID      int64
	Operation   string
	PerformedBy *int64
	PerformedAt time.Time
}

func NewTaskAuditRepository(db *gorm.DB) domain.TaskAuditRepository {
	return &taskAuditRepository{db}
}

// Create inserts a new audit log for a task operation.
// ctx: the context for the operation.
// taskID: the ID of the task being audited.
// op: the operation performed on the task (e.g., "CREATE", "UPDATE", "DELETE", "READ").
func (t *taskAuditRepository) Create(ctx context.Context, audit *domain.TaskAudit) error {
	return t.db.Create(&taskAudit{
		TaskID:      audit.TaskID,
		Operation:   audit.Operation,
		PerformedBy: audit.PerformedBy,
		PerformedAt: audit.PerformedAt,
	}).Error
}
