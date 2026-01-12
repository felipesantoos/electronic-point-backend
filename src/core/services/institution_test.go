package services

import (
	"eletronic_point/src/core/domain/errors"
	"eletronic_point/src/core/domain/institution"
	"eletronic_point/src/core/services/filters"
	"eletronic_point/src/core/services/mocks"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestInstitutionService_List(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockInstitutionPort(ctrl)
	service := NewInstitutionService(mockRepo)

	f := filters.InstitutionFilters{}
	expected := []institution.Institution{}

	mockRepo.EXPECT().List(f).Return(expected, nil)

	res, err := service.List(f)

	assert.Nil(t, err)
	assert.Equal(t, expected, res)
}

func TestInstitutionService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockInstitutionPort(ctrl)
	service := NewInstitutionService(mockRepo)

	id := uuid.New()
	mockRepo.EXPECT().Create(gomock.Any()).Return(&id, nil)

	resID, err := service.Create(nil)
	assert.Nil(t, err)
	assert.Equal(t, &id, resID)
}

func TestInstitutionService_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockInstitutionPort(ctrl)
	service := NewInstitutionService(mockRepo)

	id := uuid.New()
	expectedInstitution, _ := institution.NewBuilder().WithID(id).WithName("Institution A").Build()
	mockRepo.EXPECT().Get(id).Return(expectedInstitution, nil)

	inst, err := service.Get(id)
	assert.Nil(t, err)
	assert.Equal(t, expectedInstitution, inst)
}

func TestInstitutionService_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockInstitutionPort(ctrl)
	service := NewInstitutionService(mockRepo)

	inst, _ := institution.NewBuilder().WithID(uuid.New()).WithName("Institution A").Build()
	mockRepo.EXPECT().Update(inst).Return(nil)

	err := service.Update(inst)
	assert.Nil(t, err)
}

func TestInstitutionService_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockInstitutionPort(ctrl)
	service := NewInstitutionService(mockRepo)

	id := uuid.New()
	mockRepo.EXPECT().Delete(id).Return(nil)

	err := service.Delete(id)
	assert.Nil(t, err)
}

func TestInstitutionService_ErrorHandling(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockInstitutionPort(ctrl)
	service := NewInstitutionService(mockRepo)

	id := uuid.New()
	expectedErr := errors.NewUnexpected()
	mockRepo.EXPECT().Get(id).Return(nil, expectedErr)

	_, err := service.Get(id)
	assert.NotNil(t, err)
	assert.Equal(t, expectedErr, err)
}
