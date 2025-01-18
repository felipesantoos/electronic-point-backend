package student

import (
	"eletronic_point/src/core/domain/campus"
	"eletronic_point/src/core/domain/errors"
	"eletronic_point/src/core/domain/institution"
	"eletronic_point/src/core/domain/internship"
	"eletronic_point/src/core/domain/person"
	"eletronic_point/src/core/domain/timeRecord"
)

type Student interface {
	person.Person
	Registration() string
	ProfilePicture() *string
	Institution() institution.Institution
	Campus() campus.Campus
	Course() string
	TotalWorkload() int
	WorkloadCompleted() int
	PendingWorkload() int
	CurrentInternship() internship.Internship
	InternshipHistory() []internship.Internship
	FrequencyHistory() []timeRecord.TimeRecord

	SetRegistration(string) errors.Error
	SetProfilePicture(*string) errors.Error
	SetInstitution(institution.Institution) errors.Error
	SetCampus(campus.Campus) errors.Error
	SetCourse(string) errors.Error
	SetTotalWorkload(int) errors.Error
	SetWorkloadCompleted(int) errors.Error
	SetPendingWorkload(int) errors.Error
	SetCurrentInternship(internship.Internship) errors.Error
	SetInternshipHistory([]internship.Internship) errors.Error
	SetFrequencyHistory([]timeRecord.TimeRecord) errors.Error
}
