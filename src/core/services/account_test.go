package services

import (
	"eletronic_point/src/core/domain/account"
	"eletronic_point/src/core/domain/errors"
	"eletronic_point/src/core/domain/person"
	updatepassword "eletronic_point/src/core/domain/updatePassword"
	"eletronic_point/src/core/services/filters"
	"eletronic_point/src/core/services/mocks"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestAccountService_List(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockAccountPort(ctrl)
	service := NewAccountService(mockRepo)

	f := filters.AccountFilters{}
	expectedAccounts := []account.Account{}

	mockRepo.EXPECT().List(f).Return(expectedAccounts, nil)

	accounts, err := service.List(f)

	assert.Nil(t, err)
	assert.Equal(t, expectedAccounts, accounts)
}

func TestAccountService_FindByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockAccountPort(ctrl)
	service := NewAccountService(mockRepo)

	id := uuid.New()
	mockRepo.EXPECT().FindByID(&id).Return(nil, nil)

	_, err := service.FindByID(&id)
	assert.Nil(t, err)
}

func TestAccountService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockAccountPort(ctrl)
	service := NewAccountService(mockRepo)

	id := uuid.New()
	mockRepo.EXPECT().Create(gomock.Any()).Return(&id, nil)

	resID, err := service.Create(nil)
	assert.Nil(t, err)
	assert.Equal(t, &id, resID)
}

func TestAccountService_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockAccountPort(ctrl)
	service := NewAccountService(mockRepo)

	id := uuid.New()
	mockRepo.EXPECT().Delete(id).Return(nil)

	err := service.Delete(id)
	assert.Nil(t, err)
}

func TestAccountService_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockAccountPort(ctrl)
	service := NewAccountService(mockRepo)

	p, _ := person.NewBuilder().WithName("John Doe").WithEmail("john@example.com").WithBirthDate("1990-01-01").WithCPF("11144477735").WithPhone("82999999999").Build()
	acc, _ := account.NewBuilder().WithEmail("test@example.com").WithPassword("password123").WithPerson(p).Build()
	mockRepo.EXPECT().Update(acc).Return(nil)

	err := service.Update(acc)
	assert.Nil(t, err)
}

func TestAccountService_UpdateAccountProfile(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockAccountPort(ctrl)
	service := NewAccountService(mockRepo)

	p, _ := person.NewBuilder().WithName("John Doe").WithEmail("john@example.com").WithBirthDate("1990-01-01").WithCPF("11144477735").WithPhone("82999999999").Build()
	mockRepo.EXPECT().UpdateAccountProfile(p).Return(nil)

	err := service.UpdateAccountProfile(p)
	assert.Nil(t, err)
}

func TestAccountService_UpdateAccountPassword(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockAccountPort(ctrl)
	service := NewAccountService(mockRepo)

	accountID := uuid.New()
	updatePassword := updatepassword.New("oldpass123", "newpass123")
	mockRepo.EXPECT().UpdateAccountPassword(&accountID, updatePassword).Return(nil)

	err := service.UpdateAccountPassword(&accountID, updatePassword)
	assert.Nil(t, err)
}

func TestAccountService_ErrorHandling(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockAccountPort(ctrl)
	service := NewAccountService(mockRepo)

	id := uuid.New()
	expectedErr := errors.NewUnexpected()
	mockRepo.EXPECT().FindByID(&id).Return(nil, expectedErr)

	_, err := service.FindByID(&id)
	assert.NotNil(t, err)
	assert.Equal(t, expectedErr, err)
}

func TestAccountService_Update_ErrorHandling(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockAccountPort(ctrl)
	service := NewAccountService(mockRepo)

	p, _ := person.NewBuilder().WithName("John Doe").WithEmail("john@example.com").WithBirthDate("1990-01-01").WithCPF("11144477735").WithPhone("82999999999").Build()
	acc, _ := account.NewBuilder().WithEmail("test@example.com").WithPassword("password123").WithPerson(p).Build()
	expectedErr := errors.NewUnexpected()
	mockRepo.EXPECT().Update(acc).Return(expectedErr)

	err := service.Update(acc)
	assert.NotNil(t, err)
	assert.Equal(t, expectedErr, err)
}

func TestAccountService_UpdateAccountProfile_ErrorHandling(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockAccountPort(ctrl)
	service := NewAccountService(mockRepo)

	p, _ := person.NewBuilder().WithName("John Doe").WithEmail("john@example.com").WithBirthDate("1990-01-01").WithCPF("11144477735").WithPhone("82999999999").Build()
	expectedErr := errors.NewUnexpected()
	mockRepo.EXPECT().UpdateAccountProfile(p).Return(expectedErr)

	err := service.UpdateAccountProfile(p)
	assert.NotNil(t, err)
	assert.Equal(t, expectedErr, err)
}

func TestAccountService_UpdateAccountPassword_ErrorHandling(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockAccountPort(ctrl)
	service := NewAccountService(mockRepo)

	accountID := uuid.New()
	updatePassword := updatepassword.New("oldpass123", "newpass123")
	expectedErr := errors.NewUnexpected()
	mockRepo.EXPECT().UpdateAccountPassword(&accountID, updatePassword).Return(expectedErr)

	err := service.UpdateAccountPassword(&accountID, updatePassword)
	assert.NotNil(t, err)
	assert.Equal(t, expectedErr, err)
}
