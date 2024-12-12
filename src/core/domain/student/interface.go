package student

import (
	"eletronic_point/src/core/domain/errors"
	"eletronic_point/src/core/domain/person"
	"eletronic_point/src/core/domain/timeRecord"

	"github.com/google/uuid"
)

type Student interface {
	person.Person
	StudentID() uuid.UUID
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
	FrequencyHistory() []timeRecord.TimeRecord

	SetStudentID(uuid.UUID) errors.Error
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
	SetFrequencyHistory([]timeRecord.TimeRecord) errors.Error
}
