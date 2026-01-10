package response

import (
	"eletronic_point/src/core/domain/account"

	"github.com/google/uuid"
)

type Account struct {
	ID           *uuid.UUID    `json:"id"`
	Email        string        `json:"email,omitempty"`
	Role         Role          `json:"role"`
	Person       Person        `json:"profile"`
	Professional *Professional `json:"professional,omitempty"`
	Student      *Student      `json:"student,omitempty"`
}

type accountBuilder struct{}

func AccountBuilder() *accountBuilder {
	return &accountBuilder{}
}

func (*accountBuilder) BuildFromDomain(data account.Account) Account {
	var professional *Professional
	if data.Professional() != nil {
		professional = ProfessionalBuilder().BuildFromDomain(data.Professional())
	}
	var _student *Student
	if data.Student() != nil {
		aux := StudentBuilder().BuildFromDomain(data.Student())
		_student = &aux
	}
	return Account{
		data.ID(),
		data.Email(),
		AccountRoleBuilder().BuildFromDomain(data.Role()),
		PersonBuilder().BuildFromDomain(data.Person()),
		professional,
		_student,
	}
}

func (*accountBuilder) BuildFromDomainList(data []account.Account) []Account {
	accounts := make([]Account, 0)
	for _, acc := range data {
		accounts = append(accounts, AccountBuilder().BuildFromDomain(acc))
	}
	return accounts
}
