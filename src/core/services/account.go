package services

import (
	"dit_backend/src/core/domain/account"
	"dit_backend/src/core/domain/errors"
	updatepassword "dit_backend/src/core/domain/updatePassword"
	"dit_backend/src/core/interfaces/adapters"
	"dit_backend/src/core/interfaces/usecases"

	"github.com/google/uuid"
)

type accountService struct {
	adapter adapters.AccountAdapter
}

func NewAccountService(repository adapters.AccountAdapter) usecases.AccountUseCase {
	return &accountService{repository}
}

func (instance *accountService) List() ([]account.Account, errors.Error) {
	accounts, err := instance.adapter.List()
	if err != nil {
		return nil, errors.NewFromInfra(err)
	}
	return accounts, nil
}

func (instance *accountService) FindByID(uID uuid.UUID) (account.Account, errors.Error) {
	account, err := instance.adapter.FindByID(uID)
	if err != nil {
		return nil, errors.NewFromInfra(err)
	}
	return account, nil
}

func (instance *accountService) Create(account account.Account) (*uuid.UUID, errors.Error) {
	id, err := instance.adapter.Create(account)
	if err != nil {
		return nil, errors.NewFromInfra(err)
	}
	return id, nil
}

func (instance *accountService) UpdateAccountProfile(account account.Account) errors.Error {
	if err := instance.adapter.UpdateAccountProfile(account); err != nil {
		return errors.NewFromInfra(err)
	}
	return nil
}

func (instance *accountService) UpdateAccountPassword(accountID uuid.UUID, data updatepassword.UpdatePassword) errors.Error {
	err := instance.adapter.UpdateAccountPassword(accountID, data)
	if err != nil {
		return errors.NewFromInfra(err)
	}
	return nil
}
