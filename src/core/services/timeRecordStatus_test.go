package services

import (
	"eletronic_point/src/core/domain/errors"
	"eletronic_point/src/core/domain/timeRecordStatus"
	"eletronic_point/src/core/services/mocks"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestTimeRecordStatusService_List(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockTimeRecordStatusPort(ctrl)
	service := NewTimeRecordStatusService(mockRepo)

	expected := []timeRecordStatus.TimeRecordStatus{}

	mockRepo.EXPECT().List().Return(expected, nil)

	res, err := service.List()

	assert.Nil(t, err)
	assert.Equal(t, expected, res)
}

func TestTimeRecordStatusService_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockTimeRecordStatusPort(ctrl)
	service := NewTimeRecordStatusService(mockRepo)

	id := uuid.New()
	expectedStatus, _ := timeRecordStatus.NewBuilder().WithID(id).WithName("Pending").Build()
	mockRepo.EXPECT().Get(id).Return(expectedStatus, nil)

	status, err := service.Get(id)
	assert.Nil(t, err)
	assert.Equal(t, expectedStatus, status)
}

func TestTimeRecordStatusService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockTimeRecordStatusPort(ctrl)
	service := NewTimeRecordStatusService(mockRepo)

	id := uuid.New()
	mockRepo.EXPECT().Create(gomock.Any()).Return(&id, nil)

	resID, err := service.Create(nil)
	assert.Nil(t, err)
	assert.Equal(t, &id, resID)
}

func TestTimeRecordStatusService_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockTimeRecordStatusPort(ctrl)
	service := NewTimeRecordStatusService(mockRepo)

	status, _ := timeRecordStatus.NewBuilder().WithID(uuid.New()).WithName("Pending").Build()
	mockRepo.EXPECT().Update(status).Return(nil)

	err := service.Update(status)
	assert.Nil(t, err)
}

func TestTimeRecordStatusService_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockTimeRecordStatusPort(ctrl)
	service := NewTimeRecordStatusService(mockRepo)

	id := uuid.New()
	mockRepo.EXPECT().Delete(id).Return(nil)

	err := service.Delete(id)
	assert.Nil(t, err)
}

func TestTimeRecordStatusService_ErrorHandling(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockTimeRecordStatusPort(ctrl)
	service := NewTimeRecordStatusService(mockRepo)

	id := uuid.New()
	expectedErr := errors.NewUnexpected()
	mockRepo.EXPECT().Get(id).Return(nil, expectedErr)

	_, err := service.Get(id)
	assert.NotNil(t, err)
	assert.Equal(t, expectedErr, err)
}
