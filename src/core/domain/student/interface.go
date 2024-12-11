package student

import (
	"eletronic_point/src/core/domain/errors"

	"github.com/google/uuid"
)

type Student interface {
	ID() uuid.UUID
	Name() string
	Registration() string
	ProfilePicture() *string
	Institution() string
	Course() string
	InternshipLocationName() string
	InternshipAddress() string
	InternshipLocation() string
	TotalWorkload() int
	WorkloadCompleted() int
	PendingWorkload() int
	FrequencyHistory() string // TODO: create domain

	SetID(uuid.UUID) errors.Error
	SetName(string) errors.Error
	SetRegistration(string) errors.Error
	SetProfilePicture(*string) errors.Error
	SetInstitution(string) errors.Error
	SetCourse(string) errors.Error
	SetInternshipLocationName(string) errors.Error
	SetInternshipAddress(string) errors.Error
	SetInternshipLocation(string) errors.Error
	SetTotalWorkload(int) errors.Error
	SetWorkloadCompleted(int) errors.Error
	SetPendingWorkload(int) errors.Error
	SetFrequencyHistory(string) errors.Error // TODO: create domain
}
