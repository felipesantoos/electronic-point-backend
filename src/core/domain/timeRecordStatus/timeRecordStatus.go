package timeRecordStatus

import (
	"eletronic_point/src/core/domain/errors"
	"eletronic_point/src/core/messages"
	"eletronic_point/src/utils/validator"

	"github.com/google/uuid"
)

var _ TimeRecordStatus = &timeRecordStatus{}

type timeRecordStatus struct {
	id   uuid.UUID
	name string
}

func (trs *timeRecordStatus) ID() uuid.UUID {
	return trs.id
}

func (trs *timeRecordStatus) Name() string {
	return trs.name
}

func (trs *timeRecordStatus) SetID(id uuid.UUID) errors.Error {
	if !validator.IsUUIDValid(id) {
		return errors.NewValidationFromString(messages.TimeRecordStatusIDErrorMessage)
	}
	trs.id = id
	return nil
}

func (trs *timeRecordStatus) SetName(name string) errors.Error {
	if name == "" {
		return errors.NewFromString(messages.TimeRecordStatusNameErrorMessage)
	}
	trs.name = name
	return nil
}
