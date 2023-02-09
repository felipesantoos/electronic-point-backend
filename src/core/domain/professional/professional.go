package professional

import (
	"github.com/google/uuid"
)

type Professional interface {
	ID() *uuid.UUID
	PersonID() *uuid.UUID

	SetID(*uuid.UUID)
	SetPersonID(*uuid.UUID)
}

type professional struct {
	id       *uuid.UUID
	personID *uuid.UUID
}

func New(id *uuid.UUID, personID *uuid.UUID) Professional {
	return &professional{id, personID}
}

func NewFromMap(data map[string]interface{}) (Professional, error) {
	professional := &professional{}
	if id, err := uuid.Parse(string(data["id"].([]uint8))); err != nil {
		return nil, err
	} else {
		professional.id = &id
	}
	if id, err := uuid.Parse(string(data["person_id"].([]uint8))); err != nil {
		return nil, err
	} else {
		professional.personID = &id
	}
	return professional, nil
}

func (instance *professional) ID() *uuid.UUID {
	return instance.id
}

func (instance *professional) PersonID() *uuid.UUID {
	return instance.personID
}

func (instance *professional) SetID(id *uuid.UUID) {
	instance.id = id
}

func (instance *professional) SetPersonID(personID *uuid.UUID) {
	instance.personID = personID
}
