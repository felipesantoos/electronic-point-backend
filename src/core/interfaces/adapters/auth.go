package adapters

import (
	"backend_template/src/core/domain/account"
	"backend_template/src/core/domain/credentials"
	"backend_template/src/core/domain/errors"

	"github.com/google/uuid"
)

type AuthAdapter interface {
	Login(credentials credentials.Credentials) (account.Account, errors.Error)
	ResetAccountPassword(accountID *uuid.UUID, newPassword string) errors.Error
}
