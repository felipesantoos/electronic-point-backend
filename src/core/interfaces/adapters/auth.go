package secondary

import (
	"eletronic_point/src/core/domain/account"
	"eletronic_point/src/core/domain/credentials"
	"eletronic_point/src/core/domain/errors"

	"github.com/google/uuid"
)

type AuthPort interface {
	Login(credentials credentials.Credentials) (account.Account, errors.Error)
	ResetAccountPassword(accountID *uuid.UUID, newPassword string) errors.Error
}
