package services

import (
	"eletronic_point/src/core/domain/account"
	"eletronic_point/src/core/domain/authorization"
	"eletronic_point/src/core/domain/credentials"
	"eletronic_point/src/core/domain/errors"
	"eletronic_point/src/core/domain/person"
	"eletronic_point/src/core/domain/role"
	"eletronic_point/src/core/services/mocks"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestAuthService_Login(t *testing.T) {
	os.Setenv("SERVER_SECRET", "test-secret")
	defer os.Unsetenv("SERVER_SECRET")

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAuthPort := mocks.NewMockAuthPort(ctrl)
	mockSessionPort := mocks.NewMockSessionPort(ctrl)
	mockPwdResetPort := mocks.NewMockPasswordResetPort(ctrl)
	mockAccountPort := mocks.NewMockAccountPort(ctrl)
	service := NewAuthService(mockAuthPort, mockSessionPort, mockPwdResetPort, mockAccountPort)

	id := uuid.New()
	personID := uuid.New()
	p, _ := person.NewBuilder().WithID(personID).WithName("John Doe").WithEmail("john@example.com").WithBirthDate("1990-01-01").WithCPF("11144477735").WithPhone("82999999999").Build()
	r, _ := role.NewBuilder().WithCode(role.ADMIN_ROLE_CODE).WithName("Admin").Build()
	acc, _ := account.NewBuilder().WithID(id).WithEmail("john@example.com").WithPassword("pass123").WithRole(r).WithPerson(p).Build()
	creds := credentials.New("john@example.com", "pass123")

	// Test successful login
	t.Run("Successful login", func(t *testing.T) {
		mockAuthPort.EXPECT().Login(creds).Return(acc, nil)
		mockSessionPort.EXPECT().RemoveSession(&id).Return(nil)
		mockSessionPort.EXPECT().Store(&id, gomock.Any()).Return(nil)
		mockSessionPort.EXPECT().StoreRefreshToken(&id, gomock.Any()).Return(nil)

		auth, refresh, err := service.Login(creds)
		assert.Nil(t, err)
		assert.NotNil(t, auth)
		assert.NotNil(t, refresh)
		assert.NotEmpty(t, auth.Token())
		assert.NotEmpty(t, refresh.Token())
	})

	// Test login with adapter error
	t.Run("Login with adapter error", func(t *testing.T) {
		expectedErr := errors.NewUnexpected()
		mockAuthPort.EXPECT().Login(creds).Return(nil, expectedErr)

		auth, refresh, err := service.Login(creds)
		assert.NotNil(t, err)
		assert.Nil(t, auth)
		assert.Nil(t, refresh)
		assert.Equal(t, expectedErr, err)
	})

	// Test login with error storing session
	t.Run("Login with error storing session", func(t *testing.T) {
		expectedErr := errors.NewUnexpected()
		mockAuthPort.EXPECT().Login(creds).Return(acc, nil)
		mockSessionPort.EXPECT().RemoveSession(&id).Return(nil)
		mockSessionPort.EXPECT().Store(&id, gomock.Any()).Return(expectedErr)

		auth, refresh, err := service.Login(creds)
		assert.NotNil(t, err)
		assert.Nil(t, auth)
		assert.Nil(t, refresh)
		assert.Equal(t, expectedErr, err)
	})

	// Test login with error storing refresh token
	t.Run("Login with error storing refresh token", func(t *testing.T) {
		expectedErr := errors.NewUnexpected()
		mockAuthPort.EXPECT().Login(creds).Return(acc, nil)
		mockSessionPort.EXPECT().RemoveSession(&id).Return(nil)
		mockSessionPort.EXPECT().Store(&id, gomock.Any()).Return(nil)
		mockSessionPort.EXPECT().StoreRefreshToken(&id, gomock.Any()).Return(expectedErr)

		auth, refresh, err := service.Login(creds)
		assert.NotNil(t, err)
		assert.Nil(t, auth)
		assert.Nil(t, refresh)
		assert.Equal(t, expectedErr, err)
	})
}

func TestAuthService_Refresh(t *testing.T) {
	os.Setenv("SERVER_SECRET", "test-secret")
	defer os.Unsetenv("SERVER_SECRET")

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAuthPort := mocks.NewMockAuthPort(ctrl)
	mockSessionPort := mocks.NewMockSessionPort(ctrl)
	mockPwdResetPort := mocks.NewMockPasswordResetPort(ctrl)
	mockAccountPort := mocks.NewMockAccountPort(ctrl)
	service := NewAuthService(mockAuthPort, mockSessionPort, mockPwdResetPort, mockAccountPort)

	id := uuid.New()
	personID := uuid.New()
	p, _ := person.NewBuilder().WithID(personID).WithName("John Doe").WithEmail("john@example.com").WithBirthDate("1990-01-01").WithCPF("11144477735").WithPhone("82999999999").Build()
	r, _ := role.NewBuilder().WithCode(role.ADMIN_ROLE_CODE).WithName("Admin").Build()
	acc, _ := account.NewBuilder().WithID(id).WithEmail("john@example.com").WithPassword("pass123").WithRole(r).WithPerson(p).Build()

	refreshToken, _ := authorization.NewRefreshToken(acc)

	t.Run("Successful refresh", func(t *testing.T) {
		mockSessionPort.EXPECT().ValidateRefreshToken(&id, refreshToken.Token()).Return(true, nil)
		mockSessionPort.EXPECT().RemoveRefreshToken(&id, refreshToken.Token()).Return(nil)
		mockAccountPort.EXPECT().FindByID(&id).Return(acc, nil)
		mockSessionPort.EXPECT().Store(&id, gomock.Any()).Return(nil)
		mockSessionPort.EXPECT().StoreRefreshToken(&id, gomock.Any()).Return(nil)

		newAuth, newRefresh, err := service.Refresh(refreshToken.Token())
		assert.Nil(t, err)
		assert.NotNil(t, newAuth)
		assert.NotNil(t, newRefresh)
	})
}

func TestAuthService_Logout(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAuthPort := mocks.NewMockAuthPort(ctrl)
	mockSessionPort := mocks.NewMockSessionPort(ctrl)
	mockPwdResetPort := mocks.NewMockPasswordResetPort(ctrl)
	mockAccountPort := mocks.NewMockAccountPort(ctrl)
	service := NewAuthService(mockAuthPort, mockSessionPort, mockPwdResetPort, mockAccountPort)

	id := uuid.New()
	mockSessionPort.EXPECT().RemoveSession(&id).Return(nil)
	mockSessionPort.EXPECT().RemoveAllRefreshTokens(&id).Return(nil)

	err := service.Logout(&id)
	assert.Nil(t, err)
}

func TestAuthService_SessionExists(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAuthPort := mocks.NewMockAuthPort(ctrl)
	mockSessionPort := mocks.NewMockSessionPort(ctrl)
	mockPwdResetPort := mocks.NewMockPasswordResetPort(ctrl)
	mockAccountPort := mocks.NewMockAccountPort(ctrl)
	service := NewAuthService(mockAuthPort, mockSessionPort, mockPwdResetPort, mockAccountPort)

	id := uuid.New()
	token := "token"
	mockSessionPort.EXPECT().Exists(&id, token).Return(true, nil)

	exists, err := service.SessionExists(&id, token)
	assert.Nil(t, err)
	assert.True(t, exists)
}

func TestAuthService_AskPasswordResetMail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAuthPort := mocks.NewMockAuthPort(ctrl)
	mockSessionPort := mocks.NewMockSessionPort(ctrl)
	mockPwdResetPort := mocks.NewMockPasswordResetPort(ctrl)
	mockAccountPort := mocks.NewMockAccountPort(ctrl)
	service := NewAuthService(mockAuthPort, mockSessionPort, mockPwdResetPort, mockAccountPort)

	email := "test@example.com"
	mockPwdResetPort.EXPECT().AskPasswordResetMail(email).Return(nil)

	err := service.AskPasswordResetMail(email)
	assert.Nil(t, err)
}

func TestAuthService_FindPasswordResetByToken(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAuthPort := mocks.NewMockAuthPort(ctrl)
	mockSessionPort := mocks.NewMockSessionPort(ctrl)
	mockPwdResetPort := mocks.NewMockPasswordResetPort(ctrl)
	mockAccountPort := mocks.NewMockAccountPort(ctrl)
	service := NewAuthService(mockAuthPort, mockSessionPort, mockPwdResetPort, mockAccountPort)

	token := "reset-token"
	mockPwdResetPort.EXPECT().FindPasswordResetByToken(token).Return(nil)

	err := service.FindPasswordResetByToken(token)
	assert.Nil(t, err)
}

func TestAuthService_UpdatePasswordByPasswordReset(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAuthPort := mocks.NewMockAuthPort(ctrl)
	mockSessionPort := mocks.NewMockSessionPort(ctrl)
	mockPwdResetPort := mocks.NewMockPasswordResetPort(ctrl)
	mockAccountPort := mocks.NewMockAccountPort(ctrl)
	service := NewAuthService(mockAuthPort, mockSessionPort, mockPwdResetPort, mockAccountPort)

	token := "reset-token"
	newPassword := "newpass123"
	accountID := uuid.New()

	// Test successful password reset
	t.Run("Successful password reset", func(t *testing.T) {
		mockPwdResetPort.EXPECT().GetAccountIDByResetPasswordToken(token).Return(&accountID, nil)
		mockAuthPort.EXPECT().ResetAccountPassword(&accountID, newPassword).Return(nil)
		mockPwdResetPort.EXPECT().DeleteResetPasswordEntry(token).Return(nil)

		err := service.UpdatePasswordByPasswordReset(token, newPassword)
		assert.Nil(t, err)
	})
}
