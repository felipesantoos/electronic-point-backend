package role

import (
	"eletronic_point/src/core/domain"

	"github.com/google/uuid"
)

type Role interface {
	domain.Model

	ID() *uuid.UUID
	Name() string
	Code() string
	IsProfessional() bool
	IsAdmin() bool
	IsTeacher() bool
	IsStudent() bool
}
