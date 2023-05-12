package services

import (
	"backend_template/src/core/domain/account"
	"backend_template/src/core/domain/errors"
	"backend_template/src/core/domain/person"
	updatepassword "backend_template/src/core/domain/updatePassword"
	"backend_template/src/core/interfaces/adapters"
	"backend_template/src/core/interfaces/usecases"

	"github.com/google/uuid"
)

type accountService struct {
	adapter adapters.AccountAdapter
}

func NewAccountService(repository adapters.AccountAdapter) usecases.AccountUseCase {
	return &accountService{repository}
}

func (instance *accountService) List() ([]account.Account, errors.Error) {
	return instance.adapter.List()
}

func (instance *accountService) FindByID(uID uuid.UUID) (account.Account, errors.Error) {
	return instance.adapter.FindByID(uID)
}

func (instance *accountService) Create(account account.Account) (*uuid.UUID, errors.Error) {
	return instance.adapter.Create(account)
}

func (instance *accountService) UpdateAccountProfile(person person.Person) errors.Error {
	return instance.adapter.UpdateAccountProfile(person)
}

func (instance *accountService) UpdateAccountPassword(accountID uuid.UUID, data updatepassword.UpdatePassword) errors.Error {
	return instance.adapter.UpdateAccountPassword(accountID, data)
}
