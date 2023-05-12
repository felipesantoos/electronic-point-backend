package response

import (
	"backend_template/src/core/domain/account"

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

func (instance *accountBuilder) BuildFromDomain(data account.Account) Account {
	var professional *Professional
	if data.Professional() != nil {
		professional = ProfessionalBuilder().BuildFromDomain(data.Professional())
	}
	return Account{
		data.ID(),
		data.Email(),
		AccountRoleBuilder().BuildFromDomain(data.Role()),
		PersonBuilder().BuildFromDomain(data.Person()),
		professional,
	}
}
