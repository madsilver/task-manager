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

func Test_userHandler_Create_BindError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	bindError := errors.New("bind error")
	// Mocking ContextHTTP
	mockContext := mock_handler.NewMockContextHTTP(ctrl)
	mockContext.EXPECT().
		Bind(gomock.Any()).
		Return(bindError)
	mockContext.EXPECT().
		JSON(http.StatusBadRequest, presenter.ErrorResponse{Error: "invalid request"}).
		Return(nil)
	// Mocking UserUsecase
	mockUserUsecase := mock_usecase.NewMockUserUsecase(ctrl)
	mockUserUsecase.EXPECT().
		CreateUser(gomock.Any(), gomock.Any()).
		Times(0)
	// Mocking Logger
	mockLogger := mock_handler.NewMockLogger(ctrl)
	mockLogger.EXPECT().
		Error(mockContext, "failed to bind user", bindError)

	err := NewUserHandler(mockUserUsecase, mockLogger).Create(mockContext)

	assert.Nil(t, err)
}

func Test_userHandler_Create_EmailExists(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	errEmailExists := errors.New("ERROR: duplicate key value violates unique constraint \"users_email_key\" (SQLSTATE 23505)")
	// Mocking ContextHTTP
	mockContext := mock_handler.NewMockContextHTTP(ctrl)
	mockContext.EXPECT().
		Bind(gomock.Any()).
		DoAndReturn(func(user *domain.User) error {
			user.Email = "test@test.com"
			user.Name = "Test User"
			return nil
		})
	mockContext.EXPECT().
		JSON(http.StatusBadRequest, presenter.ErrorResponse{Error: "email already exists"}).
		Return(nil)
	// Mocking UserUsecase
	mockUserUsecase := mock_usecase.NewMockUserUsecase(ctrl)
	mockUserUsecase.EXPECT().
		CreateUser(gomock.Any(), gomock.Any()).
		Return(errEmailExists)
	// Mocking Logger
	mockLogger := mock_handler.NewMockLogger(ctrl)
	mockLogger.EXPECT().
		Error(mockContext, "email already exists", errEmailExists)

	err := NewUserHandler(mockUserUsecase, mockLogger).Create(mockContext)

	assert.Nil(t, err)
}

func Test_userHandler_Create_InternalError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	errDb := errors.New("connect: connection refused")
	// Mocking ContextHTTP
	mockContext := mock_handler.NewMockContextHTTP(ctrl)
	mockContext.EXPECT().
		Bind(gomock.Any()).
		DoAndReturn(func(user *domain.User) error {
			user.Email = "test@test.com"
			user.Name = "Test User"
			return nil
		})
	mockContext.EXPECT().
		JSON(http.StatusInternalServerError, presenter.ErrorResponse{Error: "failed to create user"}).
		Return(nil)
	// Mocking UserUsecase
	mockUserUsecase := mock_usecase.NewMockUserUsecase(ctrl)
	mockUserUsecase.EXPECT().
		CreateUser(gomock.Any(), gomock.Any()).
		Return(errDb)
	// Mocking Logger
	mockLogger := mock_handler.NewMockLogger(ctrl)
	mockLogger.EXPECT().
		Error(mockContext, "failed to create user", errDb)

	err := NewUserHandler(mockUserUsecase, mockLogger).Create(mockContext)

	assert.Nil(t, err)
}

func Test_userHandler_Create_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	// Mocking ContextHTTP
	mockContext := mock_handler.NewMockContextHTTP(ctrl)
	mockContext.EXPECT().
		Bind(gomock.Any()).
		DoAndReturn(func(user *domain.User) error {
			user.Email = "test@test.com"
			user.Name = "Test User"
			return nil
		})
	mockContext.EXPECT().
		JSON(http.StatusCreated, &presenter.UserResponse{
			ID:    123,
			Name:  "Test User",
			Email: "test@test.com",
		}).
		Return(nil)
	// Mocking UserUsecase
	mockUserUsecase := mock_usecase.NewMockUserUsecase(ctrl)
	mockUserUsecase.EXPECT().
		CreateUser(gomock.Any(), gomock.Any()).
		DoAndReturn(func(c ContextHTTP, user *domain.User) error {
			user.ID = 123 // Simulating that the user was created with ID 123
			return nil
		})
	// Mocking Logger
	mockLogger := mock_handler.NewMockLogger(ctrl)

	err := NewUserHandler(mockUserUsecase, mockLogger).Create(mockContext)

	assert.Nil(t, err)
}
