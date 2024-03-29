package controller

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	mockcontroller "github.com/madsilver/task-manager/internal/adapter/controller/mock"
	"github.com/madsilver/task-manager/internal/entity"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTaskController_FindTasks(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := mockcontroller.NewMockRepository(ctrl)
	gomock.InOrder(
		repo.EXPECT().FindAllTasks(gomock.Any()).Return([]*entity.Task{}, nil),
		repo.EXPECT().FindAllTasks(gomock.Any()).Return(nil, errors.New("error")),
	)
	e := echo.New()
	ctx := e.NewContext(httptest.NewRequest(http.MethodGet, "/", nil), httptest.NewRecorder())

	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "Status ok",
			wantErr: false,
		},
		{
			name:    "Internal server error",
			wantErr: false,
		},
	}
	c := NewTaskController(repo)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := c.FindTasks(ctx); (err != nil) != tt.wantErr {
				t.Errorf("FindTasks() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
