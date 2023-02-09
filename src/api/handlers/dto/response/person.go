package response

import (
	"dit_backend/src/core/domain/person"

	"github.com/google/uuid"
)

type Person struct {
	ID        *uuid.UUID `json:"id"`
	Name      string     `json:"name"`
	BirthDate string     `json:"birth_date"`
	Email     string     `json:"email"`
	CPF       string     `json:"cpf"`
	Phone     string     `json:"phone"`
	CreatedAt string     `json:"created_at"`
	UpdatedAt string     `json:"updated_at"`
}

type personBuilder struct{}

func PersonBuilder() *personBuilder {
	return &personBuilder{}
}

func (*personBuilder) FromDomain(data person.Person) Person {
	return Person{
		ID:        data.ID(),
		Name:      data.Name(),
		BirthDate: data.BirthDate(),
		Email:     data.Email(),
		CPF:       data.CPF(),
		Phone:     data.Phone(),
		CreatedAt: data.CreatedAt(),
		UpdatedAt: data.UpdatedAt(),
	}
}

func (instance *Person) ToDomain() person.Person {
	return person.New(
		instance.ID,
		instance.Name,
		instance.BirthDate,
		instance.Email,
		instance.CPF,
		instance.Phone,
		"",
		"",
	)
}
