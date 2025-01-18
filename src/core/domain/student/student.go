package student

import (
	"eletronic_point/src/core/domain/campus"
	"eletronic_point/src/core/domain/errors"
	"eletronic_point/src/core/domain/institution"
	"eletronic_point/src/core/domain/internship"
	"eletronic_point/src/core/domain/person"
	"eletronic_point/src/core/domain/timeRecord"
	"eletronic_point/src/core/messages"
)

var _ Student = &student{}

type student struct {
	person.Person
	name              string
	registration      string
	profilePicture    *string
	_institution      institution.Institution
	_campus           campus.Campus
	course            string
	totalWorkload     int
	workloadCompleted int
	pendingWorkload   int
	currentInternship internship.Internship
	internshipHistory []internship.Internship
	frequencyHistory  []timeRecord.TimeRecord
}

func (s *student) Registration() string {
	return s.registration
}

func (s *student) ProfilePicture() *string {
	return s.profilePicture
}

func (s *student) Institution() institution.Institution {
	return s._institution
}

func (s *student) Campus() campus.Campus {
	return s._campus
}

func (s *student) Course() string {
	return s.course
}

func (s *student) CurrentInternship() internship.Internship {
	return s.currentInternship
}

func (s *student) InternshipHistory() []internship.Internship {
	return s.internshipHistory
}

func (s *student) TotalWorkload() int {
	return s.totalWorkload
}

func (s *student) WorkloadCompleted() int {
	return s.workloadCompleted
}

func (s *student) PendingWorkload() int {
	return s.pendingWorkload
}

func (s *student) FrequencyHistory() []timeRecord.TimeRecord {
	return s.frequencyHistory
}

func (s *student) SetName(name string) errors.Error {
	if name == "" {
		return errors.NewFromString(messages.StudentNameErrorMessage)
	}
	s.name = name
	return nil
}

func (s *student) SetRegistration(registration string) errors.Error {
	if registration == "" {
		return errors.NewFromString(messages.StudentRegistrationErrorMessage)
	}
	s.registration = registration
	return nil
}

func (s *student) SetProfilePicture(profilePicture *string) errors.Error {
	s.profilePicture = profilePicture
	return nil
}

func (s *student) SetInstitution(_institution institution.Institution) errors.Error {
	if _institution == nil {
		return errors.NewFromString(messages.StudentInstitutionErrorMessage)
	}
	s._institution = _institution
	return nil
}

func (s *student) SetCampus(_campus campus.Campus) errors.Error {
	if _campus == nil {
		return errors.NewFromString(messages.StudentCampusErrorMessage)
	}
	s._campus = _campus
	return nil
}

func (s *student) SetCourse(course string) errors.Error {
	if course == "" {
		return errors.NewFromString(messages.StudentCourseErrorMessage)
	}
	s.course = course
	return nil
}

func (s *student) SetCurrentInternship(currentInternship internship.Internship) errors.Error {
	if currentInternship == nil {
		return errors.NewFromString(messages.InternshipErrorMessage)
	}
	s.currentInternship = currentInternship
	return nil
}

func (s *student) SetInternshipHistory(internshipHistory []internship.Internship) errors.Error {
	s.internshipHistory = internshipHistory
	return nil
}

func (s *student) SetTotalWorkload(totalWorkload int) errors.Error {
	if totalWorkload < 0 {
		return errors.NewFromString(messages.StudentTotalWorkloadErrorMessage)
	}
	s.totalWorkload = totalWorkload
	return nil
}

func (s *student) SetWorkloadCompleted(workloadCompleted int) errors.Error {
	if workloadCompleted < 0 {
		return errors.NewFromString(messages.StudentWorkloadCompletedErrorMessage)
	}
	s.workloadCompleted = workloadCompleted
	return nil
}

func (s *student) SetPendingWorkload(pendingWorkload int) errors.Error {
	if pendingWorkload < 0 {
		return errors.NewFromString(messages.StudentPendingWorkloadErrorMessage)
	}
	s.pendingWorkload = pendingWorkload
	return nil
}

func (s *student) SetFrequencyHistory(frequencyHistory []timeRecord.TimeRecord) errors.Error {
	s.frequencyHistory = frequencyHistory
	return nil
}
