package person

import (
	"eletronic_point/src/core/domain"
	"eletronic_point/src/core/domain/errors"

	"github.com/google/uuid"
)

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
	SetName(string) errors.Error
	SetEmail(string) errors.Error
	SetBirthDate(string) errors.Error
	SetCPF(string) errors.Error
	SetPhone(string) errors.Error
}
