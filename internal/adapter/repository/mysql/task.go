package mysql

import (
	"github.com/madsilver/task-manager/internal/adapter/core"
	"github.com/madsilver/task-manager/internal/entity"
)

type taskRepository struct {
	db core.DB
}

func NewTaskRepository(db core.DB) core.TaskRepository {
	return &taskRepository{
		db,
	}
}

func (r *taskRepository) FindAll(args any) ([]*entity.Task, error) {
	query := "SELECT * FROM Tasks"
	if args != nil {
		query = query + " WHERE UserID = ?"
	}
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

func (r *taskRepository) FindByID(args any) (*entity.Task, error) {
	query := "SELECT * FROM Tasks WHERE ID = ?"
	task := &entity.Task{}
	err := r.db.QueryRow(query, args, func(scan func(dest ...any) error) error {
		return scan(&task.ID, &task.UserID, &task.Summary, &task.Date)
	})
	return task, err
}

func (r *taskRepository) Create(task *entity.Task) error {
	query := "INSERT INTO Tasks (UserID, Summary) VALUES (?,?)"
	res, err := r.db.Save(query, &task.UserID, &task.Summary)
	if err != nil {
		return err
	}
	task.ID = uint64(res.(int64))
	return nil
}

func (r *taskRepository) Update(task *entity.Task) error {
	query := "UPDATE Tasks SET Summary = ?, Date = ? WHERE ID = ?"
	return r.db.Update(query, &task.Summary, &task.Date, &task.ID)
}

func (r *taskRepository) Delete(id any) error {
	query := "DELETE FROM Tasks WHERE ID = ?"
	return r.db.Delete(query, id)
}
