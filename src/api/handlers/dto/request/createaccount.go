package request

import (
	"dit_backend/src/core/domain/account"
	"dit_backend/src/core/domain/errors"
	"dit_backend/src/core/domain/person"
	"dit_backend/src/core/domain/role"
	"fmt"
	"net/mail"
	"regexp"
	"strings"

	"github.com/paemuri/brdoc"
)

var birthDatePattern = `^[0-9]{4}-?[0-9]{2}-?[0-9]{2}$`

type CreateAccount struct {
	Password  string `json:"password" novalidate:"true"`
	RoleCode  string `json:"role_code"`
	Name      string `json:"name"`
	BirthDate string `json:"birth_date"`
	Email     string `json:"email"`
	CPF       string `json:"cpf"`
	Phone     string `json:"phone"`
}

type createAccountBuilder struct{}

func CreateAccountBuilder() *createAccountBuilder {
	return &createAccountBuilder{}
}

func (*createAccountBuilder) FromBody(data map[string]interface{}) (*CreateAccount, errors.Error) {
	dto := &CreateAccount{}
	providedRoleCode := strings.ToLower(fmt.Sprint(data["role_code"]))
	possibleRoleCodes := role.PossibleRoleCodes()
	for _, roleCode := range possibleRoleCodes {
		if strings.ToLower(roleCode) == providedRoleCode {
			dto.RoleCode = providedRoleCode
		}
	}
	if dto.RoleCode == "" {
		return nil, errors.NewFromString("you must enter a valid role code. Valid Options: " + strings.Join(possibleRoleCodes, ", "))
	}
	dto.Name = fmt.Sprint(data["name"])
	birthDate := fmt.Sprint(data["birth_date"])
	if ok, _ := regexp.Match(birthDatePattern, []byte(birthDate)); !ok {
		return nil, errors.NewFromString("you must provide a date according to the following syntax: yyyy-MM-dd")
	}
	dto.BirthDate = birthDate
	email := fmt.Sprint(data["email"])
	if addr, _ := mail.ParseAddress(email); addr == nil {
		return nil, errors.NewFromString("you must provide a valid email!")
	}
	dto.Email = email
	if !brdoc.IsCPF(fmt.Sprint(data["cpf"])) {
		return nil, errors.NewFromString("you must provide a valid CPF!")
	}
	dto.CPF = fmt.Sprint(data["cpf"])
	dto.Phone = fmt.Sprint(data["phone"])
	return dto, nil
}

func (instance *CreateAccount) ToDomain() account.Account {
	return account.New(
		nil,
		instance.Email,
		instance.Password,
		role.New(nil, "", instance.RoleCode),
		person.New(nil, instance.Name, instance.BirthDate, instance.Email, instance.CPF, instance.Phone, "", ""),
		nil,
	)
}
