package request

import (
	"eletronic_point/src/core/domain/account"
	"eletronic_point/src/core/domain/errors"
	"eletronic_point/src/core/domain/person"
	"eletronic_point/src/core/domain/role"
)

const birthDatePattern = `^[0-9]{4}-?[0-9]{2}-?[0-9]{2}$`

type CreateAccount struct {
	Name      string `json:"name"`
	BirthDate string `json:"birth_date"`
	Email     string `json:"email"`
	CPF       string `json:"cpf"`
	Phone     string `json:"phone"`
	RoleCode  string `json:"role_code"`
}

func (c *CreateAccount) ToDomain() (account.Account, errors.Error) {
	role, err := role.New(nil, "", c.RoleCode)
	if err != nil {
		return nil, err
	}
	person, err := person.New(nil, c.Name, c.Email, c.BirthDate, c.CPF, c.Phone, "", "")
	if err != nil {
		return nil, err
	}
	return account.New(
		nil,
		c.Email,
		"",
		role,
		person,
		nil,
		nil,
	)
}
