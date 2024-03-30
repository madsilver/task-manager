package controller

import (
	"github.com/golang/mock/gomock"
	mockController "github.com/madsilver/task-manager/internal/adapter/controller/mock"
	"testing"
)

func TestNotifyController_Listen(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockBroker := mockController.NewMockBroker(ctrl)
	mockBroker.EXPECT().Consume().Times(1)
	controller := NewNotifyController(mockBroker)
	controller.Listen()
}
