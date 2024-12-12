package account

import (
	"eletronic_point/src/core/domain"
	"eletronic_point/src/core/domain/person"
	"eletronic_point/src/core/domain/professional"
	"eletronic_point/src/core/domain/role"

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
