package usecases

import (
	"dit_backend/src/core/domain/account"
	"dit_backend/src/core/domain/errors"
	updatepassword "dit_backend/src/core/domain/updatePassword"

	"github.com/google/uuid"
)

type AccountUseCase interface {
	List() ([]account.Account, errors.Error)
	FindByID(uID uuid.UUID) (account.Account, errors.Error)
	Create(account.Account) (*uuid.UUID, errors.Error)
	UpdateAccountPassword(accountID uuid.UUID, data updatepassword.UpdatePassword) errors.Error
	UpdateAccountProfile(account account.Account) errors.Error
}
