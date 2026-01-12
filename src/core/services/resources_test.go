package services

import (
	"eletronic_point/src/core/domain/role"
	"eletronic_point/src/core/services/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestResourcesService_ListAccountRoles(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockResourcesPort(ctrl)
	service := NewResourcesService(mockRepo)

	expected := []role.Role{}
	mockRepo.EXPECT().ListAccountRoles().Return(expected, nil)

	res, err := service.ListAccountRoles()
	assert.Nil(t, err)
	assert.Equal(t, expected, res)
}
