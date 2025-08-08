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
			task.OwnerID = 1
			return nil
		})
	// Mocking Logger
	mockLogger := mock_handler.NewMockLogger(ctrl)

	err := NewTaskHandler(mockTaskUsecase, mockLogger).Create(mockContext)

	assert.Nil(t, err)
}
