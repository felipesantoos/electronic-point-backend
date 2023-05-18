package account

import (
	"backend_template/src/core/domain"
	"backend_template/src/core/domain/errors"
	"backend_template/src/core/domain/person"
	"backend_template/src/core/domain/professional"
	"backend_template/src/core/domain/role"
	"net/mail"

	"github.com/google/uuid"
)

type Account interface {
	domain.Model

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

func New(id *uuid.UUID, email, password string, role role.Role, person person.Person, professional professional.Professional) (Account, errors.Error) {
	data := &account{id, email, password, role, person, professional}
	return data, data.IsValid()
}

func (acc *account) ID() *uuid.UUID {
	return acc.id
}

func (acc *account) Email() string {
	return acc.email
}

func (acc *account) Password() string {
	return acc.password
}

func (acc *account) Role() role.Role {
	return acc.role
}

func (acc *account) Person() person.Person {
	return acc.person
}

func (acc *account) Professional() professional.Professional {
	return acc.professional
}

func (acc *account) SetID(id uuid.UUID) {
	acc.id = &id
}

func (acc *account) SetEmail(email string) {
	acc.email = email
}

func (acc *account) SetPassword(password string) {
	acc.password = password
}

func (acc *account) SetRole(role role.Role) {
	acc.role = role
}

func (acc *account) SetPerson(person person.Person) {
	acc.person = person
}

func (acc *account) SetProfessional(professional professional.Professional) {
	acc.professional = professional
}

func (acc *account) IsValid() errors.Error {
	var errorMessages = []string{}
	var metadata = map[string]interface{}{"fields": []string{}}
	if addr, _ := mail.ParseAddress(acc.email); addr == nil {
		errorMessages = append(errorMessages, "you must provide a valid email!")
		metadata["fields"] = append(metadata["fields"].([]string), "email")
	}
	if err := acc.person.IsValid(); err != nil {
		return err
	}
	if acc.professional != nil && acc.professional.IsValid() != nil {
		return acc.professional.IsValid()
	}
	if len(errorMessages) != 0 {
		return errors.NewValidationWithMetadata(errorMessages, metadata)
	}
	return nil
}
