package services

import (
	"eletronic_point/src/core/domain/errors"
	"eletronic_point/src/core/domain/student"
	"eletronic_point/src/core/services/filters"
	"eletronic_point/src/core/services/mocks"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestStudentService_List(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockStudentPort(ctrl)
	service := NewStudentService(mockRepo)

	f := filters.StudentFilters{}
	expected := []student.Student{}

	mockRepo.EXPECT().List(f).Return(expected, nil)

	res, err := service.List(f)

	assert.Nil(t, err)
	assert.Equal(t, expected, res)
}

func TestStudentService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockStudentPort(ctrl)
	service := NewStudentService(mockRepo)

	id := uuid.New()
	mockRepo.EXPECT().Create(gomock.Any()).Return(&id, nil)

	resID, err := service.Create(nil)
	assert.Nil(t, err)
	assert.Equal(t, &id, resID)
}

func TestStudentService_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockStudentPort(ctrl)
	service := NewStudentService(mockRepo)

	id := uuid.New()
	f := filters.StudentFilters{}
	expectedStudent, _ := student.NewBuilder().WithRegistration("123").WithPerson(nil).Build()
	mockRepo.EXPECT().Get(id, f).Return(expectedStudent, nil)

	s, err := service.Get(id, f)
	assert.Nil(t, err)
	assert.Equal(t, expectedStudent, s)
}

func TestStudentService_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockStudentPort(ctrl)
	service := NewStudentService(mockRepo)

	s, _ := student.NewBuilder().WithRegistration("123").WithPerson(nil).Build()
	mockRepo.EXPECT().Update(s).Return(nil)

	err := service.Update(s)
	assert.Nil(t, err)
}

func TestStudentService_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockStudentPort(ctrl)
	service := NewStudentService(mockRepo)

	id := uuid.New()
	mockRepo.EXPECT().Delete(id).Return(nil)

	err := service.Delete(id)
	assert.Nil(t, err)
}

func TestStudentService_ErrorHandling(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockStudentPort(ctrl)
	service := NewStudentService(mockRepo)

	id := uuid.New()
	f := filters.StudentFilters{}
	expectedErr := errors.NewUnexpected()
	mockRepo.EXPECT().Get(id, f).Return(nil, expectedErr)

	_, err := service.Get(id, f)
	assert.NotNil(t, err)
	assert.Equal(t, expectedErr, err)
}
