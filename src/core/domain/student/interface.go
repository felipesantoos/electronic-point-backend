package student

import (
	"eletronic_point/src/core/domain/errors"
	"eletronic_point/src/core/domain/internship"
	"eletronic_point/src/core/domain/person"
	"eletronic_point/src/core/domain/timeRecord"
)

type Student interface {
	person.Person
	Registration() string
	ProfilePicture() *string
	Institution() string
	Course() string
	TotalWorkload() int
	WorkloadCompleted() int
	PendingWorkload() int
	CurrentInternship() internship.Internship
	InternshipHistory() []internship.Internship
	FrequencyHistory() []timeRecord.TimeRecord

	SetRegistration(string) errors.Error
	SetProfilePicture(*string) errors.Error
	SetInstitution(string) errors.Error
	SetCourse(string) errors.Error
	SetTotalWorkload(int) errors.Error
	SetWorkloadCompleted(int) errors.Error
	SetPendingWorkload(int) errors.Error
	SetCurrentInternship(internship.Internship) errors.Error
	SetInternshipHistory([]internship.Internship) errors.Error
	SetFrequencyHistory([]timeRecord.TimeRecord) errors.Error
}
