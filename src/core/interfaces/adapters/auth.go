package adapters

import (
	"backend_template/src/core/domain/account"
	"backend_template/src/core/domain/credentials"
	"backend_template/src/core/domain/errors"
)

type AuthAdapter interface {
	Login(credentials credentials.Credentials) (account.Account, errors.Error)
}
