package services

import (
	"eletronic_point/src/core/services/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestFileService_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockFilePort(ctrl)
	service := NewFileService(mockRepo)

	name := "test.txt"
	mockRepo.EXPECT().Get(name).Return(nil, nil)

	_, err := service.Get(name)
	assert.Nil(t, err)
}
