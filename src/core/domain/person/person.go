package person

import (
	"fmt"

	"github.com/google/uuid"
)

type Person interface {
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

func New(id *uuid.UUID, name, birthDate, email, cpf, phone, createdAt, updatedAt string) Person {
	return &person{id, name, birthDate, email, cpf, phone, createdAt, updatedAt}
}

func NewEmpty() Person {
	return &person{}
}

func NewFromDerivedMap(data map[string]interface{}) (Person, error) {
	person := &person{}
	person.name = fmt.Sprint(data["person_name"])
	person.birthDate = fmt.Sprint(data["person_birth_date"])
	person.cpf = fmt.Sprint(data["person_cpf"])
	person.email = fmt.Sprint(data["account_email"])
	person.phone = fmt.Sprint(data["person_phone"])
	person.createdAt = fmt.Sprint(data["person_created_at"])
	person.updatedAt = fmt.Sprint(data["person_updated_at"])
	if id, err := uuid.Parse(string(data["person_id"].([]uint8))); err != nil {
		return nil, err
	} else {
		person.id = &id
	}
	return person, nil
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
