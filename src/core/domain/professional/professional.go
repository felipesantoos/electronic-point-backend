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

func (p *professional) ID() *uuid.UUID {
	return p.id
}

func (p *professional) PersonID() *uuid.UUID {
	return p.personID
}

func (p *professional) SetPersonID(personID *uuid.UUID) {
	p.personID = personID
}

func (p *professional) IsValid() errors.Error {
	return nil
}
