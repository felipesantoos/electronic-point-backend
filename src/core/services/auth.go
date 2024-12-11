package services

import (
	"eletronic_point/src/core/domain/authorization"
	"eletronic_point/src/core/domain/credentials"
	"eletronic_point/src/core/domain/errors"
	"eletronic_point/src/core/interfaces/primary"
	"eletronic_point/src/core/interfaces/secondary"

	"github.com/google/uuid"
)

type authService struct {
	adapter           secondary.AuthPort
	sessionPort       secondary.SessionPort
	passwordResetPort secondary.PasswordResetPort
}

func NewAuthService(
	adapter secondary.AuthPort,
	sessionPort secondary.SessionPort,
	passwordResetPort secondary.PasswordResetPort,
) primary.AuthPort {
	return &authService{adapter, sessionPort, passwordResetPort}
}

func (s *authService) Login(credentials credentials.Credentials) (authorization.Authorization, errors.Error) {
	account, err := s.adapter.Login(credentials)
	if err != nil {
		return nil, err
	}
	auth, err := s.sessionPort.GetSessionByAccountID(account.ID())
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
		if err := s.sessionPort.Store(account.ID(), auth.Token()); err != nil {
			return nil, err
		}
	}
	return auth, nil
}

func (s *authService) Logout(accountID *uuid.UUID) errors.Error {
	return s.sessionPort.RemoveSession(accountID)
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
