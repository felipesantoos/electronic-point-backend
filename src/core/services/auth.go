package services

import (
	"backend_template/src/core/domain/authorization"
	"backend_template/src/core/domain/credentials"
	"backend_template/src/core/domain/errors"
	"backend_template/src/core/interfaces/adapters"
	"backend_template/src/core/interfaces/usecases"

	"github.com/google/uuid"
)

type authService struct {
	adapter              adapters.AuthAdapter
	sessionAdapter       adapters.SessionAdapter
	passwordResetAdapter adapters.PasswordResetAdapter
}

func NewAuthService(
	adapter adapters.AuthAdapter,
	sessionAdapter adapters.SessionAdapter,
	passwordResetAdapter adapters.PasswordResetAdapter,
) usecases.AuthUseCase {
	return &authService{adapter, sessionAdapter, passwordResetAdapter}
}

func (s *authService) Login(credentials credentials.Credentials) (authorization.Authorization, errors.Error) {
	account, err := s.adapter.Login(credentials)
	if err != nil {
		return nil, err
	}
	auth, err := s.sessionAdapter.GetSessionByAccountID(account.ID())
	var authErr errors.Error
	if err != nil {
		return nil, err
	} else if auth != nil {
		return auth, nil
	} else {
		auth, authErr = authorization.NewFromAccount(account)
		if authErr != nil {
			return nil, authErr
		}
		if err := s.sessionAdapter.Store(account.ID(), auth.Token()); err != nil {
			return nil, err
		}
	}
	return auth, nil
}

func (s *authService) Logout(accountID *uuid.UUID) errors.Error {
	return s.sessionAdapter.RemoveSession(accountID)
}

func (s *authService) SessionExists(accountID *uuid.UUID, token string) (bool, errors.Error) {
	return s.sessionAdapter.Exists(accountID, token)
}

func (s *authService) AskPasswordResetMail(email string) errors.Error {
	return s.passwordResetAdapter.AskPasswordResetMail(email)
}

func (s *authService) FindPasswordResetByToken(token string) errors.Error {
	return s.passwordResetAdapter.FindPasswordResetByToken(token)
}

func (s *authService) UpdatePasswordByPasswordReset(token, newPassword string) errors.Error {
	accountID, err := s.passwordResetAdapter.GetAccountIDByResetPasswordToken(token)
	if err != nil {
		return err
	}
	err = s.adapter.ResetAccountPassword(accountID, newPassword)
	if err != nil {
		return err
	}
	err = s.passwordResetAdapter.DeleteResetPasswordEntry(token)
	if err != nil {
		return err
	}
	return nil
}
