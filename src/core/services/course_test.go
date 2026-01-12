package services

import (
	"eletronic_point/src/core/domain/course"
	"eletronic_point/src/core/domain/errors"
	"eletronic_point/src/core/services/filters"
	"eletronic_point/src/core/services/mocks"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestCourseService_List(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockCoursePort(ctrl)
	service := NewCourseService(mockRepo)

	f := filters.CourseFilters{}
	expected := []course.Course{}

	mockRepo.EXPECT().List(f).Return(expected, nil)

	res, err := service.List(f)

	assert.Nil(t, err)
	assert.Equal(t, expected, res)
}

func TestCourseService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockCoursePort(ctrl)
	service := NewCourseService(mockRepo)

	id := uuid.New()
	mockRepo.EXPECT().Create(gomock.Any()).Return(&id, nil)

	resID, err := service.Create(nil)
	assert.Nil(t, err)
	assert.Equal(t, &id, resID)
}

func TestCourseService_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockCoursePort(ctrl)
	service := NewCourseService(mockRepo)

	id := uuid.New()
	expectedCourse, _ := course.NewBuilder().WithID(id).WithName("Course A").Build()
	mockRepo.EXPECT().Get(id).Return(expectedCourse, nil)

	c, err := service.Get(id)
	assert.Nil(t, err)
	assert.Equal(t, expectedCourse, c)
}

func TestCourseService_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockCoursePort(ctrl)
	service := NewCourseService(mockRepo)

	c, _ := course.NewBuilder().WithID(uuid.New()).WithName("Course A").Build()
	mockRepo.EXPECT().Update(c).Return(nil)

	err := service.Update(c)
	assert.Nil(t, err)
}

func TestCourseService_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockCoursePort(ctrl)
	service := NewCourseService(mockRepo)

	id := uuid.New()
	mockRepo.EXPECT().Delete(id).Return(nil)

	err := service.Delete(id)
	assert.Nil(t, err)
}

func TestCourseService_ErrorHandling(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockCoursePort(ctrl)
	service := NewCourseService(mockRepo)

	id := uuid.New()
	expectedErr := errors.NewUnexpected()
	mockRepo.EXPECT().Get(id).Return(nil, expectedErr)

	_, err := service.Get(id)
	assert.NotNil(t, err)
	assert.Equal(t, expectedErr, err)
}
