package secondary

import (
	"eletronic_point/src/core/domain/account"
	"eletronic_point/src/core/domain/errors"
	"eletronic_point/src/core/domain/person"
	updatepassword "eletronic_point/src/core/domain/updatePassword"
	"eletronic_point/src/core/services/filters"

	"github.com/google/uuid"
)

type AccountPort interface {
	List(filters.AccountFilters) ([]account.Account, errors.Error)
	FindByID(uID *uuid.UUID) (account.Account, errors.Error)
	Create(account.Account) (*uuid.UUID, errors.Error)
	Update(account.Account) errors.Error
	Delete(id uuid.UUID) errors.Error
	UpdateAccountProfile(person person.Person) errors.Error
	UpdateAccountPassword(accountID *uuid.UUID, data updatepassword.UpdatePassword) errors.Error
}
