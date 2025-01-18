package campus

import (
	"eletronic_point/src/core/domain/errors"
	"eletronic_point/src/core/messages"
	"eletronic_point/src/utils/validator"

	"github.com/google/uuid"
)

var _ Campus = &campus{}

type campus struct {
	id            uuid.UUID
	name          string
	institutionID uuid.UUID
}

func (i *campus) ID() uuid.UUID {
	return i.id
}

func (i *campus) Name() string {
	return i.name
}

func (i *campus) InstitutionID() uuid.UUID {
	return i.institutionID
}

func (i *campus) SetID(id uuid.UUID) errors.Error {
	if !validator.IsUUIDValid(id) {
		return errors.NewValidationFromString(messages.CampusIDErrorMessage)
	}
	i.id = id
	return nil
}

func (i *campus) SetName(name string) errors.Error {
	if name == "" {
		return errors.NewFromString(messages.CampusNameErrorMessage)
	}
	i.name = name
	return nil
}

func (i *campus) SetInstitutionID(institutionID uuid.UUID) errors.Error {
	if !validator.IsUUIDValid(institutionID) {
		return errors.NewValidationFromString(messages.CampusInstitutionIDErrorMessage)
	}
	i.institutionID = institutionID
	return nil
}
