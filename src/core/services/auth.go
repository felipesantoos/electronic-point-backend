package services

import (
	"dit_backend/src/core/domain/authorization"
	"dit_backend/src/core/domain/credentials"
	"dit_backend/src/core/domain/errors"
	"dit_backend/src/core/interfaces/adapters"
	"dit_backend/src/core/interfaces/usecases"

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
		return nil, errors.NewFromInfra(err)
	}
	token, err := instance.sessionAdapter.GetSessionByAccountID(*account.ID())
	var auth authorization.Authorization
	var authErr errors.Error
	if err != nil {
		return nil, errors.NewFromInfra(err)
	} else if token != "" {
		auth = authorization.NewFromToken(token)
	} else {
		auth, authErr = authorization.NewFromAccount(account)
		if authErr != nil {
			return nil, authErr
		}
		if err := instance.sessionAdapter.Store(*account.ID(), auth.Token()); err != nil {
			return nil, errors.NewFromInfra(err)
		}
	}
	return auth, nil
}

func (instance *authService) Logout(accountID uuid.UUID) errors.Error {
	if err := instance.sessionAdapter.RemoveSession(accountID); err != nil {
		return errors.NewFromInfra(err)
	}
	return nil
}

func (instance *authService) SessionExists(accountID uuid.UUID, token string) (bool, errors.Error) {
	exists, err := instance.sessionAdapter.Exists(accountID, token)
	if err != nil {
		return false, errors.NewFromInfra(err)
	}
	return exists, nil
}

func (instance *authService) AskPasswordResetMail(email string) errors.Error {
	err := instance.passwordResetAdapter.AskPasswordResetMail(email)
	if err != nil {
		return errors.NewFromInfra(err)
	}
	return nil
}

func (instance *authService) FindPasswordResetByToken(token string) errors.Error {
	if err := instance.passwordResetAdapter.FindPasswordResetByToken(token); err != nil {
		return errors.NewFromInfra(err)
	}
	return nil
}

func (instance *authService) UpdatePasswordByPasswordReset(token, newPassword string) errors.Error {
	if err := instance.passwordResetAdapter.UpdatePasswordByPasswordReset(token, newPassword); err != nil {
		return errors.NewFromInfra(err)
	}
	return nil
}
