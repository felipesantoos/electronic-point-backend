package professional

import (
	"backend_template/src/core/domain"
	"backend_template/src/core/domain/errors"

	"github.com/google/uuid"
)

type Professional interface {
	domain.Model

	ID() *uuid.UUID
	PersonID() *uuid.UUID

	SetPersonID(*uuid.UUID)
}

type professional struct {
	id       *uuid.UUID
	personID *uuid.UUID
}

func New(id *uuid.UUID, personID *uuid.UUID) (Professional, errors.Error) {
	data := &professional{id, personID}
	return data, data.IsValid()
}

func (instance *professional) ID() *uuid.UUID {
	return instance.id
}

func (instance *professional) PersonID() *uuid.UUID {
	return instance.personID
}

func (instance *professional) SetPersonID(personID *uuid.UUID) {
	instance.personID = personID
}

func (instance *professional) IsValid() errors.Error {
	return nil
}
