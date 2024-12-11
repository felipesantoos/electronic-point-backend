package services

import (
	"eletronic_point/src/core/domain/account"
	"eletronic_point/src/core/domain/errors"
	"eletronic_point/src/core/domain/person"
	updatepassword "eletronic_point/src/core/domain/updatePassword"
	"eletronic_point/src/core/interfaces/primary"
	"eletronic_point/src/core/interfaces/secondary"

	"github.com/google/uuid"
)

type accountService struct {
	adapter secondary.AccountPort
}

func NewAccountService(repository secondary.AccountPort) primary.AccountPort {
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
