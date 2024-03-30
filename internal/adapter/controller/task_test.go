package controller

import (
	"bytes"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	mockController "github.com/madsilver/task-manager/internal/adapter/controller/mock"
	"github.com/madsilver/task-manager/internal/entity"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTaskController_FindTasks(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := mockController.NewMockRepository(ctrl)
	e := echo.New()
	ctxManager := e.NewContext(httptest.NewRequest(http.MethodGet, "/v1/tasks", nil), httptest.NewRecorder())
	ctxManager.Set("role", "manager")
	ctxTech := e.NewContext(httptest.NewRequest(http.MethodGet, "/v1/tasks", nil), httptest.NewRecorder())
	ctxTech.Set("role", "technician")
	ctxTech.Set("user", uint64(1))
	tests := []struct {
		name       string
		statusCode int
		ctx        echo.Context
		err        error
	}{
		{
			name:       "should return status ok for role manager",
			ctx:        ctxManager,
			statusCode: http.StatusOK,
		},
		{
			name:       "should return status ok for role technician",
			ctx:        ctxTech,
			statusCode: http.StatusOK,
		},
		{
			name:       "should return internal server error",
			ctx:        ctxManager,
			statusCode: http.StatusInternalServerError,
			err:        errors.New("error"),
		},
	}
	c := NewTaskController(mockRepo)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo.EXPECT().FindAll(gomock.Any()).Return([]*entity.Task{}, tt.err)
			_ = c.FindTasks(tt.ctx)
			assert.Equal(t, tt.statusCode, tt.ctx.Response().Status)
		})
	}
}

func TestTaskController_FindTasksByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := mockController.NewMockRepository(ctrl)
	e := echo.New()
	ctx := e.NewContext(httptest.NewRequest(http.MethodGet, "/v1/tasks", nil), httptest.NewRecorder())
	ctx.SetParamNames("id")
	ctx.SetParamValues("1")

	tests := []struct {
		name       string
		statusCode int
		err        error
	}{
		{
			name:       "should return status ok",
			statusCode: http.StatusOK,
		},
		{
			name:       "should return internal server error",
			statusCode: http.StatusInternalServerError,
			err:        errors.New("error"),
		},
	}
	c := NewTaskController(mockRepo)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo.EXPECT().FindByID(gomock.Any()).Return(&entity.Task{}, tt.err)
			_ = c.FindTaskByID(ctx)
			assert.Equal(t, tt.statusCode, ctx.Response().Status)
		})
	}
}

func TestTaskController_CreateTask(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := mockController.NewMockRepository(ctrl)
	gomock.InOrder(
		mockRepo.EXPECT().Create(gomock.Any()).Return(nil),
		mockRepo.EXPECT().Create(gomock.Any()).Return(errors.New("error")),
	)
	e := echo.New()
	ctx := e.NewContext(httptest.NewRequest(http.MethodPost, "/v1/tasks", nil), httptest.NewRecorder())
	ctx.Request().Body = io.NopCloser(bytes.NewReader([]byte("{\"summary\": \"summary test 123\"}")))
	ctx.Request().Header.Set("Content-Type", "application/json")
	ctx.Request().ContentLength = 31
	ctx.Set("user", uint64(1))

	ctx1 := e.NewContext(httptest.NewRequest(http.MethodPost, "/v1/tasks", nil), httptest.NewRecorder())
	ctx1.Request().Body = io.NopCloser(bytes.NewReader([]byte("{\"summary\": \"summary test 123\"}")))
	ctx1.Request().Header.Set("Content-Type", "application/json")
	ctx1.Request().ContentLength = 31
	ctx1.Set("user", uint64(1))

	ctx2 := e.NewContext(httptest.NewRequest(http.MethodPost, "/v1/tasks", nil), httptest.NewRecorder())
	ctx2.Request().ContentLength = 31

	tests := []struct {
		name       string
		statusCode int
		ctx        echo.Context
		err        error
	}{
		{
			name:       "should return status ok",
			statusCode: http.StatusCreated,
			ctx:        ctx,
			err:        nil,
		},
		{
			name:       "should return internal server error",
			statusCode: http.StatusInternalServerError,
			ctx:        ctx1,
			err:        errors.New("error"),
		},
		{
			name:       "should return bad request",
			statusCode: http.StatusBadRequest,
			ctx:        ctx2,
			err:        nil,
		},
	}
	c := NewTaskController(mockRepo)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_ = c.CreateTask(tt.ctx)
			assert.Equal(t, tt.statusCode, tt.ctx.Response().Status)
		})
	}
}

func TestTaskController_UpdateTask(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := mockController.NewMockRepository(ctrl)
	mockRepo.EXPECT().FindByID(gomock.Any()).Return(&entity.Task{UserID: 1}, nil)
	mockRepo.EXPECT().Update(gomock.Any()).Return(nil)
	e := echo.New()
	ctx := e.NewContext(httptest.NewRequest(http.MethodPatch, "/v1/tasks/1", nil), httptest.NewRecorder())
	ctx.Request().Body = io.NopCloser(bytes.NewReader([]byte("{\"summary\": \"summary test 123\"}")))
	ctx.Request().Header.Set("Content-Type", "application/json")
	ctx.Request().ContentLength = 31
	ctx.Set("user", uint64(1))
	controller := NewTaskController(mockRepo)

	_ = controller.UpdateTask(ctx)

	assert.Equal(t, http.StatusOK, ctx.Response().Status)
}

func TestTaskController_UpdateTask_BindError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	e := echo.New()
	ctx := e.NewContext(httptest.NewRequest(http.MethodPatch, "/v1/tasks/1", nil), httptest.NewRecorder())
	ctx.Request().ContentLength = 31
	controller := NewTaskController(mockController.NewMockRepository(ctrl))

	_ = controller.UpdateTask(ctx)

	assert.Equal(t, http.StatusBadRequest, ctx.Response().Status)
}

func TestTaskController_UpdateTask_NotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := mockController.NewMockRepository(ctrl)
	mockRepo.EXPECT().FindByID(gomock.Any()).Return(nil, errors.New("error"))
	e := echo.New()
	ctx := e.NewContext(httptest.NewRequest(http.MethodPatch, "/v1/tasks/1", nil), httptest.NewRecorder())
	ctx.Request().Body = io.NopCloser(bytes.NewReader([]byte("{\"summary\": \"summary test 123\"}")))
	ctx.Request().Header.Set("Content-Type", "application/json")
	ctx.Request().ContentLength = 31
	ctx.Set("user", uint64(1))
	controller := NewTaskController(mockRepo)

	_ = controller.UpdateTask(ctx)

	assert.Equal(t, http.StatusNotFound, ctx.Response().Status)
}

func TestTaskController_UpdateTask_Forbidden(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := mockController.NewMockRepository(ctrl)
	mockRepo.EXPECT().FindByID(gomock.Any()).Return(&entity.Task{UserID: 1}, nil)
	e := echo.New()
	ctx := e.NewContext(httptest.NewRequest(http.MethodPatch, "/v1/tasks/1", nil), httptest.NewRecorder())
	ctx.Request().Body = io.NopCloser(bytes.NewReader([]byte("{\"summary\": \"summary test 123\"}")))
	ctx.Request().Header.Set("Content-Type", "application/json")
	ctx.Request().ContentLength = 31
	ctx.Set("user", uint64(99))
	controller := NewTaskController(mockRepo)

	_ = controller.UpdateTask(ctx)

	assert.Equal(t, http.StatusForbidden, ctx.Response().Status)
}

func TestTaskController_UpdateTask_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := mockController.NewMockRepository(ctrl)
	mockRepo.EXPECT().FindByID(gomock.Any()).Return(&entity.Task{UserID: 1}, nil)
	mockRepo.EXPECT().Update(gomock.Any()).Return(errors.New("error"))
	e := echo.New()
	ctx := e.NewContext(httptest.NewRequest(http.MethodPatch, "/v1/tasks/1", nil), httptest.NewRecorder())
	ctx.Request().Body = io.NopCloser(bytes.NewReader([]byte("{\"summary\": \"summary test 123\"}")))
	ctx.Request().Header.Set("Content-Type", "application/json")
	ctx.Request().ContentLength = 31
	ctx.Set("user", uint64(1))
	controller := NewTaskController(mockRepo)

	_ = controller.UpdateTask(ctx)

	assert.Equal(t, http.StatusInternalServerError, ctx.Response().Status)
}

func TestTaskController_DeleteTask(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := mockController.NewMockRepository(ctrl)
	mockRepo.EXPECT().FindByID(gomock.Any()).Return(nil, nil)
	mockRepo.EXPECT().Delete(gomock.Any()).Return(nil)
	e := echo.New()
	ctx := e.NewContext(httptest.NewRequest(http.MethodDelete, "/v1/tasks/1", nil), httptest.NewRecorder())
	ctx.SetParamNames("id")
	ctx.SetParamValues("1")
	controller := NewTaskController(mockRepo)

	_ = controller.DeleteTask(ctx)

	assert.Equal(t, http.StatusOK, ctx.Response().Status)
}

func TestTaskController_DeleteTask_NotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := mockController.NewMockRepository(ctrl)
	mockRepo.EXPECT().FindByID(gomock.Any()).Return(nil, errors.New("error"))
	e := echo.New()
	ctx := e.NewContext(httptest.NewRequest(http.MethodDelete, "/v1/tasks/1", nil), httptest.NewRecorder())
	ctx.SetParamNames("id")
	ctx.SetParamValues("1")
	controller := NewTaskController(mockRepo)

	_ = controller.DeleteTask(ctx)

	assert.Equal(t, http.StatusNotFound, ctx.Response().Status)
}

func TestTaskController_DeleteTask_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := mockController.NewMockRepository(ctrl)
	mockRepo.EXPECT().FindByID(gomock.Any()).Return(nil, nil)
	mockRepo.EXPECT().Delete(gomock.Any()).Return(errors.New("error"))
	e := echo.New()
	ctx := e.NewContext(httptest.NewRequest(http.MethodDelete, "/v1/tasks/1", nil), httptest.NewRecorder())
	ctx.SetParamNames("id")
	ctx.SetParamValues("1")
	controller := NewTaskController(mockRepo)

	_ = controller.DeleteTask(ctx)

	assert.Equal(t, http.StatusInternalServerError, ctx.Response().Status)
}

func TestTaskController_CloseTask(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := mockController.NewMockRepository(ctrl)
	mockRepo.EXPECT().FindByID(gomock.Any()).Return(&entity.Task{UserID: 1}, nil)
	mockRepo.EXPECT().Update(gomock.Any()).Return(nil)
	e := echo.New()
	ctx := e.NewContext(httptest.NewRequest(http.MethodPatch, "/v1/tasks/1/close", nil), httptest.NewRecorder())
	ctx.SetParamNames("id")
	ctx.SetParamValues("1")
	ctx.Set("user", uint64(1))
	controller := NewTaskController(mockRepo)

	_ = controller.CloseTask(ctx)

	assert.Equal(t, http.StatusOK, ctx.Response().Status)
}

func TestTaskController_CloseTask_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := mockController.NewMockRepository(ctrl)
	mockRepo.EXPECT().FindByID(gomock.Any()).Return(&entity.Task{UserID: 1}, nil)
	mockRepo.EXPECT().Update(gomock.Any()).Return(errors.New("error"))
	e := echo.New()
	ctx := e.NewContext(httptest.NewRequest(http.MethodPatch, "/v1/tasks/1", nil), httptest.NewRecorder())
	ctx.SetParamNames("id")
	ctx.SetParamValues("1")
	ctx.Set("user", uint64(1))
	controller := NewTaskController(mockRepo)

	_ = controller.CloseTask(ctx)

	assert.Equal(t, http.StatusInternalServerError, ctx.Response().Status)
}

func TestTaskController_CloseTask_NotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := mockController.NewMockRepository(ctrl)
	mockRepo.EXPECT().FindByID(gomock.Any()).Return(nil, errors.New("error"))
	e := echo.New()
	ctx := e.NewContext(httptest.NewRequest(http.MethodPatch, "/v1/tasks/1", nil), httptest.NewRecorder())
	ctx.SetParamNames("id")
	ctx.SetParamValues("1")
	ctx.Set("user", uint64(1))
	controller := NewTaskController(mockRepo)

	_ = controller.CloseTask(ctx)

	assert.Equal(t, http.StatusNotFound, ctx.Response().Status)
}

func TestTaskController_CloseTask_Forbidden(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := mockController.NewMockRepository(ctrl)
	mockRepo.EXPECT().FindByID(gomock.Any()).Return(&entity.Task{UserID: 1}, nil)
	e := echo.New()
	ctx := e.NewContext(httptest.NewRequest(http.MethodPatch, "/v1/tasks/1", nil), httptest.NewRecorder())
	ctx.SetParamNames("id")
	ctx.SetParamValues("1")
	ctx.Set("user", uint64(99))
	controller := NewTaskController(mockRepo)

	_ = controller.CloseTask(ctx)

	assert.Equal(t, http.StatusForbidden, ctx.Response().Status)
}
