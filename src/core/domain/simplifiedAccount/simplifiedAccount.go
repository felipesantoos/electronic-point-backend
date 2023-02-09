package simplifiedAccount

import (
	"fmt"

	"github.com/google/uuid"
)

type SimplifiedAccount interface {
	ID() *uuid.UUID
	Name() string
	BirthDate() string
	Email() string
	CPF() string

	SetID(*uuid.UUID)
	SetName(string)
	SetBirthDate(string)
	SetEmail(string)
	SetCPF(string)
}

type simplifiedAccount struct {
	id        *uuid.UUID
	name      string
	birthDate string
	email     string
	cpf       string
}

func New(id *uuid.UUID, name, birthDate, email, cpf string) SimplifiedAccount {
	return &simplifiedAccount{id, name, birthDate, email, cpf}
}

func NewFromMap(data map[string]interface{}) (SimplifiedAccount, error) {
	account := &simplifiedAccount{}
	if id, err := uuid.Parse(string(data["account_id"].([]uint8))); err != nil {
		return nil, err
	} else {
		account.id = &id
	}
	account.SetName(fmt.Sprint(data["person_name"]))
	account.SetBirthDate(fmt.Sprint(data["person_birth_date"]))
	account.SetEmail(fmt.Sprint(data["account_email"]))
	account.SetCPF(fmt.Sprint(data["person_cpf"]))
	return account, nil
}

func (instance *simplifiedAccount) ID() *uuid.UUID {
	return instance.id
}

func (instance *simplifiedAccount) Name() string {
	return instance.name
}

func (instance *simplifiedAccount) BirthDate() string {
	return instance.birthDate
}

func (instance *simplifiedAccount) Email() string {
	return instance.email
}

func (instance *simplifiedAccount) CPF() string {
	return instance.cpf
}

func (instance *simplifiedAccount) SetID(id *uuid.UUID) {
	instance.id = id
}

func (instance *simplifiedAccount) SetName(name string) {
	instance.name = name
}

func (instance *simplifiedAccount) SetBirthDate(birthDate string) {
	instance.birthDate = birthDate
}

func (instance *simplifiedAccount) SetEmail(email string) {
	instance.email = email
}

func (instance *simplifiedAccount) SetCPF(cpf string) {
	instance.cpf = cpf
}
