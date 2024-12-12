package person

import (
	"eletronic_point/src/core/domain/errors"
	"net/mail"
	"regexp"
	"strings"

	"github.com/google/uuid"
	"github.com/paemuri/brdoc"
)

var _ Person = &person{}

const birthDatePattern = `^[0-9]{4}-?[0-9]{2}-?[0-9]{2}$`

type person struct {
	id        *uuid.UUID
	name      string
	email     string
	birthDate string
	cpf       string
	phone     string
	createdAt string
	updatedAt string
}

func New(id *uuid.UUID, name, email, birthDate, cpf, phone, createdAt, updatedAt string) (Person, errors.Error) {
	data := &person{id, name, email, birthDate, cpf, phone, createdAt, updatedAt}
	return data, data.IsValid()
}

func (p *person) ID() *uuid.UUID {
	return p.id
}

func (p *person) Name() string {
	return p.name
}

func (p *person) Email() string {
	return p.email
}

func (p *person) BirthDate() string {
	return p.birthDate
}

func (p *person) CPF() string {
	return p.cpf
}

func (p *person) Phone() string {
	return p.phone
}

func (p *person) CreatedAt() string {
	return p.createdAt
}

func (p *person) UpdatedAt() string {
	return p.updatedAt
}

func (p *person) SetID(id *uuid.UUID) {
	p.id = id
}

func (p *person) SetStringID(id string) error {
	parsedID, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	p.id = &parsedID
	return nil
}

func (p *person) SetName(name string) errors.Error {
	if len(strings.Split(name, " ")) < 2 {
		return errors.NewFromString("Name must contain at least two words.")
	}
	p.name = name
	return nil
}

func (p *person) SetEmail(email string) errors.Error {
	if _, err := mail.ParseAddress(email); err != nil {
		return errors.NewFromString("Invalid email address.")
	}
	p.email = email
	return nil
}

func (p *person) SetBirthDate(birthDate string) errors.Error {
	if ok, _ := regexp.MatchString(birthDatePattern, birthDate); !ok {
		return errors.NewFromString("Birth date must follow the format yyyy-MM-dd.")
	}
	p.birthDate = birthDate
	return nil
}

func (p *person) SetCPF(cpf string) errors.Error {
	if len(cpf) != 11 || !brdoc.IsCPF(cpf) {
		return errors.NewFromString("Invalid CPF number.")
	}
	p.cpf = cpf
	return nil
}

func (p *person) SetPhone(phone string) errors.Error {
	if len(phone) < 10 {
		return errors.NewFromString("Phone number must contain at least 10 digits.")
	}
	p.phone = phone
	return nil
}

func (p *person) IsValid() errors.Error {
	var errorMessages = []string{}
	var fields = []string{}
	if len(strings.Split(p.name, " ")) == 1 {
		errorMessages = append(errorMessages, "you need to provide a name with two words or more.")
		fields = append(fields, "name")
	}
	if len(p.cpf) != 11 {
		errorMessages = append(errorMessages, "the CPF number must have 11 characters")
		fields = append(fields, "cpf")
	}
	if !brdoc.IsCPF(p.cpf) {
		errorMessages = append(errorMessages, "you need to provide a valid CPF")
		fields = append(fields, "cpf")
	}
	if addr, _ := mail.ParseAddress(p.email); addr == nil {
		errorMessages = append(errorMessages, "you must provide a valid email!")
		fields = append(fields, "email")
	}
	if ok, _ := regexp.Match(birthDatePattern, []byte(p.birthDate)); !ok {
		errorMessages = append(errorMessages, "you must provide a date according to the following syntax: yyyy-MM-dd")
		fields = append(fields, "birth_date")
	}
	if len(errorMessages) != 0 {
		return errors.NewValidationWithMetadata(errorMessages, map[string]interface{}{"fields": fields})
	}
	return nil
}
