package adapters

import (
	"backend_template/src/core/domain/account"
	"backend_template/src/core/domain/errors"
	"backend_template/src/core/domain/person"
	updatepassword "backend_template/src/core/domain/updatePassword"

	"github.com/google/uuid"
)

type AccountAdapter interface {
	List() ([]account.Account, errors.Error)
	FindByID(uID *uuid.UUID) (account.Account, errors.Error)
	Create(account.Account) (*uuid.UUID, errors.Error)
	UpdateAccountProfile(person person.Person) errors.Error
	UpdateAccountPassword(accountID *uuid.UUID, data updatepassword.UpdatePassword) errors.Error
}
