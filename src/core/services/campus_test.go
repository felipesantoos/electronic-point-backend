package services

import (
	"eletronic_point/src/core/domain/campus"
	"eletronic_point/src/core/domain/errors"
	"eletronic_point/src/core/services/filters"
	"eletronic_point/src/core/services/mocks"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestCampusService_List(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockCampusPort(ctrl)
	service := NewCampusService(mockRepo)

	f := filters.CampusFilters{}
	expected := []campus.Campus{}

	mockRepo.EXPECT().List(f).Return(expected, nil)

	res, err := service.List(f)

	assert.Nil(t, err)
	assert.Equal(t, expected, res)
}

func TestCampusService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockCampusPort(ctrl)
	service := NewCampusService(mockRepo)

	id := uuid.New()
	mockRepo.EXPECT().Create(gomock.Any()).Return(&id, nil)

	resID, err := service.Create(nil)
	assert.Nil(t, err)
	assert.Equal(t, &id, resID)
}

func TestCampusService_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockCampusPort(ctrl)
	service := NewCampusService(mockRepo)

	id := uuid.New()
	expectedCampus, _ := campus.NewBuilder().WithID(id).WithName("Campus A").WithInstitutionID(uuid.New()).Build()
	mockRepo.EXPECT().Get(id).Return(expectedCampus, nil)

	c, err := service.Get(id)
	assert.Nil(t, err)
	assert.Equal(t, expectedCampus, c)
}

func TestCampusService_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockCampusPort(ctrl)
	service := NewCampusService(mockRepo)

	c, _ := campus.NewBuilder().WithID(uuid.New()).WithName("Campus A").WithInstitutionID(uuid.New()).Build()
	mockRepo.EXPECT().Update(c).Return(nil)

	err := service.Update(c)
	assert.Nil(t, err)
}

func TestCampusService_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockCampusPort(ctrl)
	service := NewCampusService(mockRepo)

	id := uuid.New()
	mockRepo.EXPECT().Delete(id).Return(nil)

	err := service.Delete(id)
	assert.Nil(t, err)
}

func TestCampusService_ErrorHandling(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockCampusPort(ctrl)
	service := NewCampusService(mockRepo)

	id := uuid.New()
	expectedErr := errors.NewUnexpected()
	mockRepo.EXPECT().Get(id).Return(nil, expectedErr)

	_, err := service.Get(id)
	assert.NotNil(t, err)
	assert.Equal(t, expectedErr, err)
}
