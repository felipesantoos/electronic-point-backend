package institution

import (
	"eletronic_point/src/core/domain/errors"
	"eletronic_point/src/core/messages"
	"eletronic_point/src/utils/validator"

	"github.com/google/uuid"
)

var _ Institution = &institution{}

type institution struct {
	id   uuid.UUID
	name string
}

func (i *institution) ID() uuid.UUID {
	return i.id
}

func (i *institution) Name() string {
	return i.name
}

func (i *institution) SetID(id uuid.UUID) errors.Error {
	if !validator.IsUUIDValid(id) {
		return errors.NewValidationFromString(messages.InstitutionIDErrorMessage)
	}
	i.id = id
	return nil
}

func (i *institution) SetName(name string) errors.Error {
	if name == "" {
		return errors.NewFromString(messages.InstitutionNameErrorMessage)
	}
	i.name = name
	return nil
}
