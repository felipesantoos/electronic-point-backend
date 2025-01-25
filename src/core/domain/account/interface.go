package account

import (
	"eletronic_point/src/core/domain"
	"eletronic_point/src/core/domain/person"
	"eletronic_point/src/core/domain/professional"
	"eletronic_point/src/core/domain/role"
	"eletronic_point/src/core/domain/student"

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
	Student() student.Student

	SetID(uuid.UUID)
	SetEmail(string)
	SetPassword(string)
	SetRole(role.Role)
	SetPerson(person.Person)
	SetProfessional(professional.Professional)
	SetStudent(student.Student)
}
