package adapters

import (
	"dit_backend/src/core/domain/account"
	"dit_backend/src/core/domain/credentials"
	"dit_backend/src/infra"
)

type AuthAdapter interface {
	Login(credentials credentials.Credentials) (account.Account, infra.Error)
}
