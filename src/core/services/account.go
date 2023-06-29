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

func (s *accountService) List() ([]account.Account, errors.Error) {
	return s.adapter.List()
}

func (s *accountService) FindByID(uID *uuid.UUID) (account.Account, errors.Error) {
	return s.adapter.FindByID(uID)
}

func (s *accountService) Create(account account.Account) (*uuid.UUID, errors.Error) {
	return s.adapter.Create(account)
}

func (s *accountService) UpdateAccountProfile(person person.Person) errors.Error {
	return s.adapter.UpdateAccountProfile(person)
}

func (s *accountService) UpdateAccountPassword(accountID *uuid.UUID, data updatepassword.UpdatePassword) errors.Error {
	return s.adapter.UpdateAccountPassword(accountID, data)
}
