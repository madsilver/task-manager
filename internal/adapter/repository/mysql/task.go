package mysql

import (
	"github.com/madsilver/task-manager/internal/entity"
)

type DB interface {
	Query(query string, args any, fn func(scan func(dest ...any) error) error) error
	QueryRow(query string, args any, fn func(scan func(dest ...any) error) error) error
	Save(query string, args ...any) (any, error)
	Update(query string, args ...any) error
	Delete(query string, args any) error
}

type TaskRepository struct {
	db DB
}

func NewTaskRepository(db DB) *TaskRepository {
	return &TaskRepository{
		db,
	}
}

func (r *TaskRepository) FindAll(args any) ([]*entity.Task, error) {
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

func (r *TaskRepository) FindByID(args any) (*entity.Task, error) {
	query := "SELECT * FROM Tasks WHERE ID = ?"
	task := &entity.Task{}
	err := r.db.QueryRow(query, args, func(scan func(dest ...any) error) error {
		return scan(&task.ID, &task.UserID, &task.Summary, &task.Date)
	})
	return task, err
}

func (r *TaskRepository) Create(task *entity.Task) error {
	query := "INSERT INTO Tasks (UserID, Summary) VALUES (?,?)"
	res, err := r.db.Save(query, &task.UserID, &task.Summary)
	if err != nil {
		return err
	}
	task.ID = uint64(res.(int64))
	return nil
}

func (r *TaskRepository) Update(task *entity.Task) error {
	query := "UPDATE Tasks SET Summary = ? WHERE ID = ?"
	return r.db.Update(query, &task.Summary, &task.ID)
}

func (r *TaskRepository) Delete(id any) error {
	query := "DELETE FROM Tasks WHERE ID = ?"
	return r.db.Delete(query, id)
}
