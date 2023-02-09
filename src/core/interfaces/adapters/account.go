package adapters

import (
	"dit_backend/src/core/domain/account"
	updatepassword "dit_backend/src/core/domain/updatePassword"
	"dit_backend/src/infra"

	"github.com/google/uuid"
)

type AccountAdapter interface {
	List() ([]account.Account, infra.Error)
	FindByID(uID uuid.UUID) (account.Account, infra.Error)
	Create(account.Account) (*uuid.UUID, infra.Error)
	UpdateAccountPassword(accountID uuid.UUID, data updatepassword.UpdatePassword) infra.Error
	UpdateAccountProfile(account account.Account) infra.Error
}
