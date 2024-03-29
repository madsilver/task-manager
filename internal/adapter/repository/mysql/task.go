package mysql

import (
	"github.com/madsilver/task-manager/internal/entity"
	"github.com/madsilver/task-manager/internal/infra/db"
)

type TaskRepository struct {
	db db.DB
}

func NewTaskRepository(db db.DB) *TaskRepository {
	return &TaskRepository{
		db,
	}
}

func (r *TaskRepository) FindAllTasks(args any) ([]*entity.Task, error) {
	query := "SELECT * FROM Tasks"
	if args != "" {
		query = query + " WHERE ID = ?"
	}
	return r.FindAll(query, args)
}

func (r *TaskRepository) FindAllTasksByUser(args any) ([]*entity.Task, error) {
	query := "SELECT * FROM Tasks WHERE UserID = ?"
	return r.FindAll(query, args)
}

func (r *TaskRepository) FindAll(query string, args any) ([]*entity.Task, error) {
	var tasks []*entity.Task
	err := r.db.Query(query, args, func(scan func(dest ...any) error) error {
		task := &entity.Task{}
		err := scan(&task.ID, &task.UserID, &task.Summary, &task.Date)
		if err == nil {
			tasks = append(tasks, task)
		}
		return err
	})
	return tasks, err
}
