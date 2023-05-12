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

func (instance *account) IsValid() errors.Error {
	var errorMessages = []string{}
	var metadata = map[string]interface{}{"fields": []string{}}
	if addr, _ := mail.ParseAddress(instance.email); addr == nil {
		errorMessages = append(errorMessages, "you must provide a valid email!")
		metadata["fields"] = append(metadata["fields"].([]string), "email")
	}
	if err := instance.person.IsValid(); err != nil {
		return err
	}
	if instance.professional != nil && instance.professional.IsValid() != nil {
		return instance.professional.IsValid()
	}
	if len(errorMessages) != 0 {
		return errors.NewValidationWithMetadata(errorMessages, metadata)
	}
	return nil
}
