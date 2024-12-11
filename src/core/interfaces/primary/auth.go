package primary

import (
	"eletronic_point/src/core/domain/authorization"
	"eletronic_point/src/core/domain/credentials"
	"eletronic_point/src/core/domain/errors"

	"github.com/google/uuid"
)

type AuthPort interface {
	Login(credentials.Credentials) (authorization.Authorization, errors.Error)
	Logout(accountID *uuid.UUID) errors.Error
	SessionExists(accountId *uuid.UUID, token string) (bool, errors.Error)
	AskPasswordResetMail(email string) errors.Error
	FindPasswordResetByToken(token string) errors.Error
	UpdatePasswordByPasswordReset(token, newPassword string) errors.Error
}
