package person

import (
	"backend_template/src/core/domain"
	"backend_template/src/core/domain/errors"
	"net/mail"
	"regexp"
	"strings"

	"github.com/google/uuid"
	"github.com/paemuri/brdoc"
)

const birthDatePattern = `^[0-9]{4}-?[0-9]{2}-?[0-9]{2}$`

type Person interface {
	domain.Model

	ID() *uuid.UUID
	Name() string
	Email() string
	BirthDate() string
	CPF() string
	Phone() string
	CreatedAt() string
	UpdatedAt() string

	SetID(*uuid.UUID)
	SetStringID(string) error
}

type person struct {
	id        *uuid.UUID
	name      string
	birthDate string
	email     string
	cpf       string
	phone     string
	createdAt string
	updatedAt string
}

func New(id *uuid.UUID, name, birthDate, email, cpf, phone, createdAt, updatedAt string) (Person, errors.Error) {
	data := &person{id, name, birthDate, email, cpf, phone, createdAt, updatedAt}
	return data, data.IsValid()
}

func (instance *person) ID() *uuid.UUID {
	return instance.id
}

func (instance *person) Name() string {
	return instance.name
}

func (instance *person) Email() string {
	return instance.email
}

func (instance *person) BirthDate() string {
	return instance.birthDate
}

func (instance *person) CPF() string {
	return instance.cpf
}

func (instance *person) Phone() string {
	return instance.phone
}

func (instance *person) CreatedAt() string {
	return instance.createdAt
}

func (instance *person) UpdatedAt() string {
	return instance.updatedAt
}

func (instance *person) SetID(id *uuid.UUID) {
	instance.id = id
}

func (instance *person) SetStringID(id string) error {
	if id, err := uuid.Parse(id); err != nil {
		return err
	} else {
		instance.id = &id
	}
	return nil
}

func (instance *person) IsValid() errors.Error {
	var errorMessages = []string{}
	var fields = []string{}
	if len(strings.Split(instance.name, " ")) == 1 {
		errorMessages = append(errorMessages, "you need to provide a name with two words or more.")
		fields = append(fields, "name")
	}
	if len(instance.cpf) != 11 {
		errorMessages = append(errorMessages, "the CPF number must have 11 characters")
		fields = append(fields, "cpf")
	}
	if !brdoc.IsCPF(instance.cpf) {
		errorMessages = append(errorMessages, "you need to provide a valid CPF")
		fields = append(fields, "cpf")
	}
	if addr, _ := mail.ParseAddress(instance.email); addr == nil {
		errorMessages = append(errorMessages, "you must provide a valid email!")
		fields = append(fields, "email")
	}
	if ok, _ := regexp.Match(birthDatePattern, []byte(instance.birthDate)); !ok {
		errorMessages = append(errorMessages, "you must provide a date according to the following syntax: yyyy-MM-dd")
		fields = append(fields, "birth_date")
	}
	if len(errorMessages) != 0 {
		return errors.NewValidationWithMetadata(errorMessages, map[string]interface{}{"fields": fields})
	}
	return nil
}
