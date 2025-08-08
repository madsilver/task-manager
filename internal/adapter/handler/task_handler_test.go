package handler

import (
	"errors"
	mock_handler "github.com/madsilver/task-manager/internal/adapter/handler/mock"
	"github.com/madsilver/task-manager/internal/adapter/presenter"
	"github.com/madsilver/task-manager/internal/domain"
	mock_usecase "github.com/madsilver/task-manager/internal/usecase/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"net/http"
	"testing"
)

func Test_taskHandler_Create_BindError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	bindError := errors.New("bind error")
	// Mocking ContextHTTP
	mockContext := mock_handler.NewMockContextHTTP(ctrl)
	mockContext.EXPECT().Bind(gomock.Any()).Return(bindError)
	mockContext.EXPECT().
		JSON(http.StatusBadRequest, presenter.ErrorResponse{Error: "invalid request"}).
		Return(nil)
	// Mocking TaskUsecase
	mockTaskUsecase := mock_usecase.NewMockTaskUseCase(ctrl)
	mockTaskUsecase.EXPECT().
		CreateTask(gomock.Any(), gomock.Any()).
		Times(0)
	// Mocking Logger
	mockLogger := mock_handler.NewMockLogger(ctrl)
	mockLogger.EXPECT().Error(mockContext, "failed to bind task", bindError)

	err := NewTaskHandler(mockTaskUsecase, mockLogger).Create(mockContext)

	assert.Nil(t, err)
}

func Test_taskHandler_Create_InternalError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	errDb := errors.New("connect: connection refused")
	// Mocking ContextHTTP
	mockContext := mock_handler.NewMockContextHTTP(ctrl)
	mockContext.EXPECT().Value("user_id").Return(float64(1))
	mockContext.EXPECT().Bind(gomock.Any()).Return(nil)
	mockContext.EXPECT().
		JSON(http.StatusInternalServerError, presenter.ErrorResponse{Error: "failed to create task"}).
		Return(nil)
	// Mocking TaskUsecase
	mockTaskUsecase := mock_usecase.NewMockTaskUseCase(ctrl)
	mockTaskUsecase.EXPECT().
		CreateTask(gomock.Any(), gomock.Any()).
		Return(errDb)
	// Mocking Logger
	mockLogger := mock_handler.NewMockLogger(ctrl)
	mockLogger.EXPECT().Error(mockContext, "failed to create task", errDb)

	err := NewTaskHandler(mockTaskUsecase, mockLogger).Create(mockContext)

	assert.Nil(t, err)
}

func Test_taskHandler_Create_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	// Mocking ContextHTTP
	mockContext := mock_handler.NewMockContextHTTP(ctrl)
	mockContext.EXPECT().Value("user_id").Return(float64(1))
	mockContext.EXPECT().Bind(gomock.Any()).Return(nil)
	mockContext.EXPECT().
		JSON(http.StatusCreated, &presenter.TaskResponse{
			ID:      123,
			OwnerID: 1,
		}).
		Return(nil)
	// Mocking TaskUsecase
	mockTaskUsecase := mock_usecase.NewMockTaskUseCase(ctrl)
	mockTaskUsecase.EXPECT().
		CreateTask(gomock.Any(), gomock.Any()).
		DoAndReturn(func(c ContextHTTP, task *domain.Task) error {
			task.ID = 123
			return nil
		})
	// Mocking Logger
	mockLogger := mock_handler.NewMockLogger(ctrl)

	err := NewTaskHandler(mockTaskUsecase, mockLogger).Create(mockContext)

	assert.Nil(t, err)
}

func Test_taskHandler_IsAllowedToRead(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	managerCtx := mock_handler.NewMockContextHTTP(ctrl)
	managerCtx.EXPECT().Value("roles").Return([]interface{}{"manager"})
	technicianCtx := mock_handler.NewMockContextHTTP(ctrl)
	technicianCtx.EXPECT().Value("roles").Return([]interface{}{"technician"}).AnyTimes()
	technicianCtx.EXPECT().Value("user_id").Return(float64(1)).AnyTimes()
	randomRoleCtx := mock_handler.NewMockContextHTTP(ctrl)
	randomRoleCtx.EXPECT().Value("roles").Return([]interface{}{"random_role"})
	randomRoleCtx.EXPECT().Value("user_id").Return(nil)

	type args struct {
		c    ContextHTTP
		task *domain.Task
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "manager allowed to read task",
			args: args{
				c:    managerCtx,
				task: &domain.Task{OwnerID: 1},
			},
			want: true,
		},
		{
			name: "technician allowed to read own task",
			args: args{
				c:    technicianCtx,
				task: &domain.Task{OwnerID: 1},
			},
			want: true,
		},
		{
			name: "technician not allowed to read other task",
			args: args{
				c:    technicianCtx,
				task: &domain.Task{OwnerID: 2},
			},
			want: false,
		},
		{
			name: "random role not allowed to read task",
			args: args{
				c:    randomRoleCtx,
				task: &domain.Task{OwnerID: 1},
			},
			want: false,
		},
	}
	h := &taskHandler{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, h.IsAllowedToRead(tt.args.c, tt.args.task), "IsAllowedToRead(%v, %v)", tt.args.c, tt.args.task)
		})
	}
}
