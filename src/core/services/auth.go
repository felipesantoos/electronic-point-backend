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

func (instance *authService) Login(credentials credentials.Credentials) (authorization.Authorization, errors.Error) {
	account, err := instance.adapter.Login(credentials)
	if err != nil {
		return nil, err
	}
	token, err := instance.sessionAdapter.GetSessionByAccountID(*account.ID())
	var auth authorization.Authorization
	var authErr errors.Error
	if err != nil {
		return nil, err
	} else if token != "" {
		auth = authorization.NewFromToken(token)
	} else {
		auth, authErr = authorization.NewFromAccount(account)
		if authErr != nil {
			return nil, authErr
		}
		if err := instance.sessionAdapter.Store(*account.ID(), auth.Token()); err != nil {
			return nil, err
		}
	}
	return auth, nil
}

func (instance *authService) Logout(accountID uuid.UUID) errors.Error {
	return instance.sessionAdapter.RemoveSession(accountID)
}

func (instance *authService) SessionExists(accountID uuid.UUID, token string) (bool, errors.Error) {
	return instance.sessionAdapter.Exists(accountID, token)
}

func (instance *authService) AskPasswordResetMail(email string) errors.Error {
	return instance.passwordResetAdapter.AskPasswordResetMail(email)
}

func (instance *authService) FindPasswordResetByToken(token string) errors.Error {
	return instance.passwordResetAdapter.FindPasswordResetByToken(token)
}

func (instance *authService) UpdatePasswordByPasswordReset(token, newPassword string) errors.Error {
	return instance.passwordResetAdapter.UpdatePasswordByPasswordReset(token, newPassword)
}
