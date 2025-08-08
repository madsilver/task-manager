package usecase

import (
	"context"
	mock_handler "github.com/madsilver/task-manager/internal/adapter/handler/mock"
	"github.com/madsilver/task-manager/internal/domain"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
)

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
		c    context.Context
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
	h := &taskUseCase{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, h.IsAllowedToRead(tt.args.c, tt.args.task), "IsAllowedToRead(%v, %v)", tt.args.c, tt.args.task)
		})
	}
}
