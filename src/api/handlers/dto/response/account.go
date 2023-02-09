package response

import (
	"dit_backend/src/core/domain/account"

	"github.com/google/uuid"
)

type Account struct {
	ID           *uuid.UUID    `json:"id"`
	Email        string        `json:"email,omitempty"`
	Role         Role          `json:"role"`
	Person       Person        `json:"profile"`
	Professional *Professional `json:"professional,omitempty"`
}

type accountBuilder struct{}

func AccountBuilder() *accountBuilder {
	return &accountBuilder{}
}

func (*accountBuilder) FromDomain(data account.Account) *Account {
	account := &Account{
		ID:     data.ID(),
		Role:   AccountRoleBuilder().FromDomain(data.Role()),
		Person: PersonBuilder().FromDomain(data.Person()),
	}
	if data.Professional() != nil {
		account.Professional = ProfessionalBuilder().FromDomain(data.Professional())
	}
	return account
}
