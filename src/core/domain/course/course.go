package course

import (
	"eletronic_point/src/core/domain/errors"
	"eletronic_point/src/core/messages"
	"eletronic_point/src/utils/validator"

	"github.com/google/uuid"
)

var _ Course = &course{}

type course struct {
	id   uuid.UUID
	name string
}

func (i *course) ID() uuid.UUID {
	return i.id
}

func (i *course) Name() string {
	return i.name
}

func (i *course) SetID(id uuid.UUID) errors.Error {
	if !validator.IsUUIDValid(id) {
		return errors.NewValidationFromString(messages.CourseIDErrorMessage)
	}
	i.id = id
	return nil
}

func (i *course) SetName(name string) errors.Error {
	if name == "" {
		return errors.NewFromString(messages.CourseNameErrorMessage)
	}
	i.name = name
	return nil
}
