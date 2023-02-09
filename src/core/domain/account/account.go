package account

import (
	"dit_backend/src/core/domain"
	"dit_backend/src/core/domain/person"
	"dit_backend/src/core/domain/professional"
	"dit_backend/src/core/domain/role"
	"errors"
	"fmt"
	"strings"

	"github.com/google/uuid"
)

type Account interface {
	ID() *uuid.UUID
	Email() string
	Password() string
	Role() role.Role
	Person() person.Person
	Professional() professional.Professional

	SetID(uuid.UUID)
	SetEmail(string)
	SetPassword(string)
	SetRole(role.Role)
	SetPerson(person.Person)
	SetProfessional(professional.Professional)
}

type account struct {
	id           *uuid.UUID
	email        string
	password     string
	role         role.Role
	person       person.Person
	professional professional.Professional
}

func New(id *uuid.UUID, email, password string, role role.Role, person person.Person, professional professional.Professional) Account {
	return &account{id, email, password, role, person, professional}
}

func NewFromMap(data map[string]interface{}) (Account, error) {
	account := &account{
		email:    fmt.Sprint(data["email"]),
		password: fmt.Sprint(data["password"]),
	}
	if id, err := uuid.Parse(string(data["id"].([]uint8))); err != nil {
		return nil, err
	} else {
		account.id = &id
	}
	if role, err := role.NewFromDerivedMap(data); err != nil {
		return nil, err
	} else {
		account.SetRole(role)
	}
	if person, err := person.NewFromDerivedMap(data); err != nil {
		return nil, err
	} else {
		account.person = person
	}
	if err := account.fillAccountRoleInfo(data); err != nil {
		return nil, err
	}
	return account, nil
}

func (instance *account) ID() *uuid.UUID {
	return instance.id
}

func (instance *account) Email() string {
	return instance.email
}

func (instance *account) Password() string {
	return instance.password
}

func (instance *account) Role() role.Role {
	return instance.role
}

func (instance *account) Person() person.Person {
	return instance.person
}

func (instance *account) Professional() professional.Professional {
	return instance.professional
}

func (instance *account) SetID(id uuid.UUID) {
	instance.id = &id
}

func (instance *account) SetEmail(email string) {
	instance.email = email
}

func (instance *account) SetPassword(password string) {
	instance.password = password
}

func (instance *account) SetRole(role role.Role) {
	instance.role = role
}

func (instance *account) SetPerson(person person.Person) {
	instance.person = person
}

func (instance *account) SetProfessional(professional professional.Professional) {
	instance.professional = professional
}

func (instance *account) fillAccountRoleInfo(data map[string]interface{}) error {
	if strings.ToLower(instance.role.Code()) == role.ProfessionalRoleCode() {
		if professionalData := domain.BuildMapWithParentName(data, role.ProfessionalRoleCode()); len(professionalData) == 0 {
			return errors.New("you must provide the professional instance properties")
		} else {
			professionalData["person_id"] = data["person_id"]
			professional, err := professional.NewFromMap(professionalData)
			if err != nil {
				return err
			}
			instance.professional = professional
		}
	}
	return nil
}
