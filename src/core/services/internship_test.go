package services

import (
	"eletronic_point/src/core/domain/errors"
	"eletronic_point/src/core/domain/internship"
	"eletronic_point/src/core/services/filters"
	"eletronic_point/src/core/services/mocks"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestInternshipService_List(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockInternshipPort(ctrl)
	service := NewInternshipService(mockRepo)

	f := filters.InternshipFilters{}
	expected := []internship.Internship{}

	mockRepo.EXPECT().List(f).Return(expected, nil)

	res, err := service.List(f)

	assert.Nil(t, err)
	assert.Equal(t, expected, res)
}

func TestInternshipService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockInternshipPort(ctrl)
	service := NewInternshipService(mockRepo)

	id := uuid.New()
	mockRepo.EXPECT().Create(gomock.Any()).Return(&id, nil)

	resID, err := service.Create(nil)
	assert.Nil(t, err)
	assert.Equal(t, &id, resID)
}

func TestInternshipService_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockInternshipPort(ctrl)
	service := NewInternshipService(mockRepo)

	id := uuid.New()
	expectedInternship, _ := internship.NewBuilder().WithID(id).WithStartedIn(time.Now()).WithLocation(nil).Build()
	mockRepo.EXPECT().Get(id).Return(expectedInternship, nil)

	intern, err := service.Get(id)
	assert.Nil(t, err)
	assert.Equal(t, expectedInternship, intern)
}

func TestInternshipService_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockInternshipPort(ctrl)
	service := NewInternshipService(mockRepo)

	intern, _ := internship.NewBuilder().WithID(uuid.New()).WithStartedIn(time.Now()).WithLocation(nil).Build()
	mockRepo.EXPECT().Update(intern).Return(nil)

	err := service.Update(intern)
	assert.Nil(t, err)
}

func TestInternshipService_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockInternshipPort(ctrl)
	service := NewInternshipService(mockRepo)

	id := uuid.New()
	mockRepo.EXPECT().Delete(id).Return(nil)

	err := service.Delete(id)
	assert.Nil(t, err)
}

func TestInternshipService_ErrorHandling(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockInternshipPort(ctrl)
	service := NewInternshipService(mockRepo)

	id := uuid.New()
	expectedErr := errors.NewUnexpected()
	mockRepo.EXPECT().Get(id).Return(nil, expectedErr)

	_, err := service.Get(id)
	assert.NotNil(t, err)
	assert.Equal(t, expectedErr, err)
}
