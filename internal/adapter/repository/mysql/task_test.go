package mysql

import (
	"errors"
	"github.com/golang/mock/gomock"
	mockcore "github.com/madsilver/task-manager/internal/adapter/core/mock"
	"github.com/madsilver/task-manager/internal/entity"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTaskRepository_FindAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockFn := mockQueryFunction(1, 1, "test summary", "2024-03-29 10:00:00")
	mockDB := mockcore.NewMockDB(ctrl)
	mockDB.EXPECT().
		Query(gomock.Any(), gomock.Any(), gomock.Any()).
		DoAndReturn(mockFn)
	repos := NewTaskRepository(mockDB)

	tasks, err := repos.FindAll(1)

	assert.Nil(t, err)
	assert.NotNil(t, tasks)
}

func TestTaskRepository_FindByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockFn := mockQueryRowFunction(1, 1, "test summary", "2024-03-29 10:00:00")
	mockDB := mockcore.NewMockDB(ctrl)
	mockDB.EXPECT().
		QueryRow(gomock.Any(), gomock.Any(), gomock.Any()).
		DoAndReturn(mockFn)
	repos := NewTaskRepository(mockDB)

	task, err := repos.FindByID(nil)

	assert.Nil(t, err)
	assert.NotNil(t, task)
}

func TestTaskRepository_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockDB := mockcore.NewMockDB(ctrl)
	mockDB.EXPECT().
		Save(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(int64(1), nil)
	repos := NewTaskRepository(mockDB)

	task := &entity.Task{}
	err := repos.Create(task)

	assert.Nil(t, err)
	assert.Equal(t, uint64(1), task.ID)
}

func TestTaskRepository_CreateError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockDB := mockcore.NewMockDB(ctrl)
	mockDB.EXPECT().
		Save(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(nil, errors.New("error"))
	repos := NewTaskRepository(mockDB)

	task := &entity.Task{}
	err := repos.Create(task)

	assert.NotNil(t, err)
	assert.Equal(t, uint64(0), task.ID)
}

func TestTaskRepository_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockDB := mockcore.NewMockDB(ctrl)
	mockDB.EXPECT().
		Update(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
		Return(nil)
	repos := NewTaskRepository(mockDB)

	err := repos.Update(&entity.Task{})

	assert.Nil(t, err)
}

func TestTaskRepository_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockDB := mockcore.NewMockDB(ctrl)
	mockDB.EXPECT().
		Delete(gomock.Any(), gomock.Any()).
		Return(nil)
	repos := NewTaskRepository(mockDB)

	err := repos.Delete(1)

	assert.Nil(t, err)
}

func mockQueryFunction(id uint64, userId uint64, summary string, date string) any {
	return func(query string, args any, fn func(scan func(dest ...any) error) error) error {
		return fn(func(dest ...any) error {
			*dest[0].(*uint64) = id
			*dest[1].(*uint64) = userId
			*dest[2].(*string) = summary
			*dest[3].(**string) = &date
			return nil
		})
	}
}

func mockQueryRowFunction(id uint64, userId uint64, summary string, date string) any {
	return func(query string, args any, fn func(scan func(dest ...any) error) error) error {
		return fn(func(dest ...any) error {
			*dest[0].(*uint64) = id
			*dest[1].(*uint64) = userId
			*dest[2].(*string) = summary
			*dest[3].(**string) = &date
			return nil
		})
	}
}
