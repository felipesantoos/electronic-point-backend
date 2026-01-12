package services

import (
	"eletronic_point/src/core/domain/errors"
	"eletronic_point/src/core/domain/internshipLocation"
	"eletronic_point/src/core/services/filters"
	"eletronic_point/src/core/services/mocks"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestInternshipLocationService_List(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockInternshipLocationPort(ctrl)
	service := NewInternshipLocationService(mockRepo)

	f := filters.InternshipLocationFilters{}
	expected := []internshipLocation.InternshipLocation{}

	mockRepo.EXPECT().List(f).Return(expected, nil)

	res, err := service.List(f)

	assert.Nil(t, err)
	assert.Equal(t, expected, res)
}

func TestInternshipLocationService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockInternshipLocationPort(ctrl)
	service := NewInternshipLocationService(mockRepo)

	id := uuid.New()
	mockRepo.EXPECT().Create(gomock.Any()).Return(&id, nil)

	resID, err := service.Create(nil)
	assert.Nil(t, err)
	assert.Equal(t, &id, resID)
}

func TestInternshipLocationService_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockInternshipLocationPort(ctrl)
	service := NewInternshipLocationService(mockRepo)

	id := uuid.New()
	expectedLocation, _ := internshipLocation.NewBuilder().WithID(id).WithName("Location").WithNumber("123").WithStreet("Street").WithNeighborhood("Neighborhood").WithCity("City").WithZipCode("12345").Build()
	mockRepo.EXPECT().Get(id).Return(expectedLocation, nil)

	location, err := service.Get(id)
	assert.Nil(t, err)
	assert.Equal(t, expectedLocation, location)
}

func TestInternshipLocationService_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockInternshipLocationPort(ctrl)
	service := NewInternshipLocationService(mockRepo)

	location, _ := internshipLocation.NewBuilder().WithID(uuid.New()).WithName("Location").WithNumber("123").WithStreet("Street").WithNeighborhood("Neighborhood").WithCity("City").WithZipCode("12345").Build()
	mockRepo.EXPECT().Update(location).Return(nil)

	err := service.Update(location)
	assert.Nil(t, err)
}

func TestInternshipLocationService_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockInternshipLocationPort(ctrl)
	service := NewInternshipLocationService(mockRepo)

	id := uuid.New()
	mockRepo.EXPECT().Delete(id).Return(nil)

	err := service.Delete(id)
	assert.Nil(t, err)
}

func TestInternshipLocationService_ErrorHandling(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockInternshipLocationPort(ctrl)
	service := NewInternshipLocationService(mockRepo)

	id := uuid.New()
	expectedErr := errors.NewUnexpected()
	mockRepo.EXPECT().Get(id).Return(nil, expectedErr)

	_, err := service.Get(id)
	assert.NotNil(t, err)
	assert.Equal(t, expectedErr, err)
}
