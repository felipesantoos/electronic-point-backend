package account

import (
	"eletronic_point/src/core/domain/errors"
	"eletronic_point/src/core/domain/person"
	"eletronic_point/src/core/domain/professional"
	"eletronic_point/src/core/domain/role"
	"eletronic_point/src/core/domain/student"
	"net/mail"

	"github.com/google/uuid"
)

type account struct {
	id           *uuid.UUID
	email        string
	password     string
	role         role.Role
	person       person.Person
	professional professional.Professional
	_student     student.Student
}

func New(id *uuid.UUID, email, password string, role role.Role, person person.Person, professional professional.Professional, _student student.Student) (Account, errors.Error) {
	data := &account{id, email, password, role, person, professional, _student}
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

func (acc *account) Student() student.Student {
	return acc._student
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

func (acc *account) SetStudent(_student student.Student) {
	acc._student = _student
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
	if acc._student != nil && acc._student.IsValid() != nil {
		return acc._student.IsValid()
	}
	if len(errorMessages) != 0 {
		return errors.NewValidationWithMetadata(errorMessages, metadata)
	}
	return nil
}
