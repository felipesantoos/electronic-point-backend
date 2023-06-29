package simplifiedAccount

import (
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

func (sacc *simplifiedAccount) ID() *uuid.UUID {
	return sacc.id
}

func (sacc *simplifiedAccount) Name() string {
	return sacc.name
}

func (sacc *simplifiedAccount) BirthDate() string {
	return sacc.birthDate
}

func (sacc *simplifiedAccount) Email() string {
	return sacc.email
}

func (sacc *simplifiedAccount) CPF() string {
	return sacc.cpf
}

func (sacc *simplifiedAccount) SetID(id *uuid.UUID) {
	sacc.id = id
}

func (sacc *simplifiedAccount) SetName(name string) {
	sacc.name = name
}

func (sacc *simplifiedAccount) SetBirthDate(birthDate string) {
	sacc.birthDate = birthDate
}

func (sacc *simplifiedAccount) SetEmail(email string) {
	sacc.email = email
}

func (sacc *simplifiedAccount) SetCPF(cpf string) {
	sacc.cpf = cpf
}
