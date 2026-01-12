package student

import (
	"eletronic_point/src/core/domain/campus"
	"eletronic_point/src/core/domain/course"
	"eletronic_point/src/core/domain/errors"
	"eletronic_point/src/core/domain/institution"
	"eletronic_point/src/core/domain/internship"
	"eletronic_point/src/core/domain/person"
	"eletronic_point/src/core/domain/timeRecord"
	"eletronic_point/src/core/messages"
	"eletronic_point/src/utils/validator"

	"github.com/google/uuid"
)

var _ Student = &student{}

type student struct {
	person.Person
	name                 string
	registration         string
	profilePicture       *string
	_institution         institution.Institution
	_campus              campus.Campus
	_course              course.Course
	totalWorkload        int
	workloadCompleted    int
	pendingWorkload      int
	responsibleTeacherID uuid.UUID
	currentInternships   []internship.Internship
	internshipHistory    []internship.Internship
	frequencyHistory     []timeRecord.TimeRecord
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

func (s *student) Course() course.Course {
	return s._course
}

func (s *student) CurrentInternships() []internship.Internship {
	return s.currentInternships
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

func (s *student) ResponsibleTeacherID() uuid.UUID {
	return s.responsibleTeacherID
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

func (s *student) SetCourse(_course course.Course) errors.Error {
	if _course == nil {
		return errors.NewFromString(messages.StudentCourseErrorMessage)
	}
	s._course = _course
	return nil
}

func (s *student) SetCurrentInternships(currentInternships []internship.Internship) errors.Error {
	s.currentInternships = currentInternships
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

func (s *student) SetResponsibleTeacherID(responsibleTeacherID uuid.UUID) errors.Error {
	if !validator.IsUUIDValid(responsibleTeacherID) {
		return errors.NewFromString(messages.StudentResponsibleTeacherIDErrorMessage)
	}
	s.responsibleTeacherID = responsibleTeacherID
	return nil
}

func (s *student) SetFrequencyHistory(frequencyHistory []timeRecord.TimeRecord) errors.Error {
	s.frequencyHistory = frequencyHistory
	return nil
}

func (s *student) IsValid() errors.Error {
	if s.Person == nil {
		return errors.NewValidationFromString(messages.PersonErrorMessage)
	}
	return s.Person.IsValid()
}
