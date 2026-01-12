package services

import (
	"eletronic_point/src/apps/api/utils"
	"eletronic_point/src/core/domain/authorization"
	"eletronic_point/src/core/domain/credentials"
	"eletronic_point/src/core/domain/errors"
	"eletronic_point/src/core/interfaces/primary"
	"eletronic_point/src/core/interfaces/secondary"
	"eletronic_point/src/core/messages"

	"github.com/google/uuid"
)

type authService struct {
	adapter           secondary.AuthPort
	sessionPort       secondary.SessionPort
	passwordResetPort secondary.PasswordResetPort
	accountPort       secondary.AccountPort
}

func NewAuthService(
	adapter secondary.AuthPort,
	sessionPort secondary.SessionPort,
	passwordResetPort secondary.PasswordResetPort,
	accountPort secondary.AccountPort,
) primary.AuthPort {
	return &authService{adapter, sessionPort, passwordResetPort, accountPort}
}

func (s *authService) Login(credentials credentials.Credentials) (authorization.Authorization, authorization.Authorization, errors.Error) {
	account, err := s.adapter.Login(credentials)
	if err != nil {
		return nil, nil, err
	}

	// Always generate new tokens
	accessToken, authErr := authorization.NewFromAccount(account)
	if authErr != nil {
		return nil, nil, authErr
	}

	refreshToken, authErr := authorization.NewRefreshToken(account)
	if authErr != nil {
		return nil, nil, authErr
	}

	// Remove old session if exists
	_ = s.sessionPort.RemoveSession(account.ID())

	// Store new session
	if err := s.sessionPort.Store(account.ID(), accessToken.Token()); err != nil {
		return nil, nil, err
	}

	// Store new refresh token
	if err := s.sessionPort.StoreRefreshToken(account.ID(), refreshToken.Token()); err != nil {
		return nil, nil, err
	}

	return accessToken, refreshToken, nil
}

func (s *authService) Refresh(refreshToken string) (authorization.Authorization, authorization.Authorization, errors.Error) {
	claims, extractErr := utils.ExtractTokenClaims(refreshToken)
	if extractErr != nil {
		return nil, nil, errors.NewFromString(messages.InvalidRefreshTokenErrorMessage)
	}

	if claims.Type != authorization.REFRESH_TOKEN_TYPE {
		return nil, nil, errors.NewFromString(messages.InvalidRefreshTokenErrorMessage)
	}

	accountID, uuidErr := uuid.Parse(claims.AccountID)
	if uuidErr != nil {
		return nil, nil, errors.NewFromString(messages.InvalidRefreshTokenErrorMessage)
	}

	// Validate refresh token in Redis
	valid, err := s.sessionPort.ValidateRefreshToken(&accountID, refreshToken)
	if err != nil {
		return nil, nil, err
	}
	if !valid {
		return nil, nil, errors.NewFromString(messages.InvalidRefreshTokenErrorMessage)
	}

	// Rotation: remove old refresh token
	if err := s.sessionPort.RemoveRefreshToken(&accountID, refreshToken); err != nil {
		return nil, nil, err
	}

	// Get full account
	account, err := s.accountPort.FindByID(&accountID)
	if err != nil {
		return nil, nil, err
	}

	// Generate new tokens
	newAccessToken, authErr := authorization.NewFromAccount(account)
	if authErr != nil {
		return nil, nil, authErr
	}

	newRefreshToken, authErr := authorization.NewRefreshToken(account)
	if authErr != nil {
		return nil, nil, authErr
	}

	// Update session
	if err := s.sessionPort.Store(&accountID, newAccessToken.Token()); err != nil {
		return nil, nil, err
	}

	// Store new refresh token
	if err := s.sessionPort.StoreRefreshToken(&accountID, newRefreshToken.Token()); err != nil {
		return nil, nil, err
	}

	return newAccessToken, newRefreshToken, nil
}

func (s *authService) Logout(accountID *uuid.UUID) errors.Error {
	if err := s.sessionPort.RemoveSession(accountID); err != nil {
		return err
	}
	return s.sessionPort.RemoveAllRefreshTokens(accountID)
}

func (s *authService) SessionExists(accountID *uuid.UUID, token string) (bool, errors.Error) {
	return s.sessionPort.Exists(accountID, token)
}

func (s *authService) AskPasswordResetMail(email string) errors.Error {
	return s.passwordResetPort.AskPasswordResetMail(email)
}

func (s *authService) FindPasswordResetByToken(token string) errors.Error {
	return s.passwordResetPort.FindPasswordResetByToken(token)
}

func (s *authService) UpdatePasswordByPasswordReset(token, newPassword string) errors.Error {
	accountID, err := s.passwordResetPort.GetAccountIDByResetPasswordToken(token)
	if err != nil {
		return err
	}
	err = s.adapter.ResetAccountPassword(accountID, newPassword)
	if err != nil {
		return err
	}
	err = s.passwordResetPort.DeleteResetPasswordEntry(token)
	if err != nil {
		return err
	}
	return nil
}
