package repository

import (
	"context"
	"github.com/madsilver/task-manager/internal/adapter/repository/utils"
	"github.com/madsilver/task-manager/internal/domain"
	"gorm.io/gorm"
	"time"
)

type taskRepository struct {
	db *gorm.DB
}

type tasks struct {
	ID        int64 `gorm:"primaryKey"`
	Title     string
	Summary   string
	OwnerID   int64
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewTaskRepository(db *gorm.DB) domain.TaskRepository {
	return &taskRepository{db}
}

// Create inserts a new task into the database.
// ctx: the context for the operation.
// task: the task to be created
// Returns an error if the creation fails.
func (r *taskRepository) Create(ctx context.Context, task *domain.Task) error {
	model := tasks{
		Title:   task.Title,
		Summary: task.Summary,
		OwnerID: task.OwnerID,
	}
	if err := r.db.WithContext(ctx).Create(&model).Error; err != nil {
		return err
	}

	task.ID = model.ID
	task.CreatedAt = model.CreatedAt
	task.UpdatedAt = model.UpdatedAt

	return nil
}

// GetAll retrieves all tasks from the database.
// ctx: the context for the operation.
// Returns a slice of tasks or an error if the retrieval fails.
func (r *taskRepository) GetAll(ctx context.Context) ([]*domain.Task, error) {
	var models []*tasks
	if err := r.db.WithContext(ctx).Find(&models).Error; err != nil {
		return nil, err
	}

	var tasksList []*domain.Task
	for _, model := range models {
		tasksList = append(tasksList, &domain.Task{
			ID:        model.ID,
			Title:     model.Title,
			Summary:   model.Summary,
			OwnerID:   model.OwnerID,
			CreatedAt: model.CreatedAt,
			UpdatedAt: model.UpdatedAt,
		})
	}
	return tasksList, nil
}

// GetByID retrieves a task by its ID from the database.
// ctx: the context for the operation.
// id: the ID of the task to retrieve
// Returns the task if found, or an error if not found or any other issue occurs.
func (r *taskRepository) GetByID(ctx context.Context, id int64) (*domain.Task, error) {
	var model tasks
	if err := r.db.WithContext(ctx).First(&model, id).Error; err != nil {
		return nil, err
	}

	return &domain.Task{
		ID:        model.ID,
		Title:     model.Title,
		Summary:   model.Summary,
		OwnerID:   model.OwnerID,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	}, nil
}

// Update modifies an existing task in the database.
// ctx: the context for the operation.
// task: the task with updated information
// Returns an error if the update fails.
func (r *taskRepository) Update(ctx context.Context, task *domain.Task) error {
	model := tasks{
		ID:      task.ID,
		Title:   task.Title,
		Summary: task.Summary,
		OwnerID: task.OwnerID,
	}

	return r.db.WithContext(ctx).Save(&model).Error
}

// Delete removes a task from the database by its ID.
// ctx: the context for the operation.
// id: the ID of the task to delete
// Returns an error if the deletion fails.
func (r *taskRepository) Delete(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Delete(&tasks{}, id).Error
}

// BeforeSave hooks for encrypting the Summary field.
func (t *tasks) BeforeSave(tx *gorm.DB) (err error) {
	t.Summary, err = utils.Encrypt(t.Summary)
	return
}

// AfterFind hooks for decrypting the Summary field after retrieval
func (t *tasks) AfterFind(tx *gorm.DB) (err error) {
	t.Summary, err = utils.Decrypt(t.Summary)
	return
}
