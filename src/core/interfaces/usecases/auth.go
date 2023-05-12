package usecases

import (
	"backend_template/src/core/domain/authorization"
	"backend_template/src/core/domain/credentials"
	"backend_template/src/core/domain/errors"

	"github.com/google/uuid"
)

type AuthUseCase interface {
	Login(credentials.Credentials) (authorization.Authorization, errors.Error)
	Logout(accountID uuid.UUID) errors.Error
	SessionExists(accountId uuid.UUID, token string) (bool, errors.Error)
	AskPasswordResetMail(email string) errors.Error
	FindPasswordResetByToken(token string) errors.Error
	UpdatePasswordByPasswordReset(token, newPassword string) errors.Error
}
